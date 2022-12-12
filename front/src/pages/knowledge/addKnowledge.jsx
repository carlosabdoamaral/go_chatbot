import {
  Form,
  Step,
  Container,
  Button,
  Message,
  Menu,
  TextArea,
  Grid,
} from "semantic-ui-react";
import React from "react";
import { useState } from "react";
import { connect } from "react-redux";
import axios from "axios";
import { env } from "../../global/variables";

const _style = {
  form: {
    margin: "2em 0",
  },

  textareaForm: {
    minHeight: 200,
    margin: "1em 0",
  },
};

const AddKnowledge = ({ props, dispatch }) => {
  const [internalServerError, setInternalServerError] = useState(false);
  const [formIsLoading, setFormIsLoading] = useState(false);

  const [message, setMessage] = useState("");
  const [messageHasError, setMessageHasError] = useState(false);

  const [answer, setAnswer] = useState("");
  const [answerHasError, setAnswerHasError] = useState(false);

  const [textAreaValue, setTextAreaValue] = useState(
    "Your message\nYour answer\n\nYour message\nYour answer\n"
  );
  const [textAreaPlaceholder] = useState(
    "Your message\nYour answer\n\nYour message\nYour answer\n"
  );

  const [currentStep, setCurrentStep] = useState(-1);

  const [activeViewCategory, setActiveViewCategory] = useState("SINGLE");

  const steps = [
    {
      title: "Preparando objeto",
      description: "Choose your shipping options",
    },
    {
      title: "Adicionar na fila",
      description: "Choose your shipping options",
    },
    {
      title: "Confirmado!",
      description: "",
    },
  ];

  async function handleSendKnowledgeRequest(isList, body) {
    if (isList) {
      await axios
        .post(env.sendListKnowledge, body)
        .then((res) => {
          setCurrentStep(2);
        })
        .catch((err) => {
          setInternalServerError(true);
          return;
        });
    } else {
      await axios
        .post(env.sendSingleKnowledge, body)
        .then((res) => {
          setCurrentStep(2);
        })
        .catch((err) => {
          setInternalServerError(true);
          return;
        });
    }
  }

  function sendSingleKnowledge() {
    const messageIsEmpty = !message;
    setMessageHasError(messageIsEmpty);

    const answerIsEmpty = !answer;
    setAnswerHasError(answerIsEmpty);

    if (messageIsEmpty || answerIsEmpty) {
      setFormIsLoading(false);
      setCurrentStep(-1);
      return;
    }

    const requestBody = {
      message: message,
      answer: answer,
    };

    handleSendKnowledgeRequest(false, requestBody);
    setCurrentStep(1);

    setTimeout(() => {
      cleanForm();
    }, 1500);
  }

  async function sendListKnowledge() {
    let bodyArr = [];
    let emptySpaceID = null;

    let lines = textAreaValue.split("\n");
    for (let i = 0; i < lines.length; i++) {
      if (lines[i] === "") {
        emptySpaceID = i;
      }

      if (i === emptySpaceID) {
        let body = {
          message: lines[i - 2],
          answer: lines[i - 1],
        };

        if (lines[i] !== "") {
          body.category = lines[i];
        }

        bodyArr.push(body);
      }
    }

    setCurrentStep(1);
    handleSendKnowledgeRequest(true, bodyArr);

    setTimeout(() => {
      cleanForm();
    }, 1500);
  }

  async function sendKnowledge() {
    setFormIsLoading(true);
    setCurrentStep(0); // Preparando
    setInternalServerError(false);

    if (activeViewCategory === "SINGLE") {
      sendSingleKnowledge();
    } else {
      sendListKnowledge();
    }
  }

  function cleanForm(isList) {
    setFormIsLoading(false);
    setInternalServerError(false);
    setCurrentStep(-1);

    if (isList) {
      setTextAreaValue("");
    } else {
      setMessage("");
      setAnswer("");
    }
  }

  function changeCategory(e, { name }) {
    setActiveViewCategory(name);
  }

  return (
    <Container style={{ paddingTop: 30 }}>
      {internalServerError ||
        messageHasError ||
        (answerHasError && (
          <Message negative>
            <Message.Header>Important!</Message.Header>
            <p>Something wrong happen</p>
          </Message>
        ))}

      <Menu attached="top" tabular>
        {["Single", "List"].map((c, i) => (
          <Menu.Item
            name={c}
            key={i}
            active={activeViewCategory.toUpperCase() === c.toUpperCase()}
            onClick={changeCategory}
          />
        ))}
      </Menu>
      <Form loading={formIsLoading} style={_style.form}>
        {activeViewCategory.toUpperCase() === "SINGLE" && (
          <div>
            <Form.Group widths={"equal"}>
              <Form.Input
                fluid
                label="Message"
                placeholder="type a message"
                value={message}
                error={messageHasError}
                onChange={(e) => {
                  setMessage(e.target.value);
                }}
              />
              <Form.Input
                fluid
                label="Answer"
                placeholder="type an answer"
                value={answer}
                error={answerHasError}
                onChange={(e) => {
                  setAnswer(e.target.value);
                }}
              />
            </Form.Group>
          </div>
        )}
        {activeViewCategory.toUpperCase() === "LIST" && (
          <Container>
            <TextArea
              style={_style.textareaForm}
              placeholder={textAreaPlaceholder}
              value={textAreaValue}
              onChange={(e) => {
                setTextAreaValue(e.target.value);
              }}
            />
            <Grid>
              <Grid.Column
                floated="right"
                style={{
                  color: "red",
                  width: "fit-content",
                  marginBottom: "1em",
                }}
              >
                <p>* The last line must be empty!</p>
              </Grid.Column>
            </Grid>
          </Container>
        )}

        <Form.Group widths={"equal"}>
          <Button
            fluid
            content="Send!"
            icon="right arrow"
            labelPosition="right"
            onClick={() => {
              sendKnowledge();
            }}
          />
        </Form.Group>
      </Form>

      <Step.Group ordered widths={8}>
        {steps.map((step, i) => (
          <Step
            key={i}
            completed={currentStep > i}
            active={currentStep === i}
            disabled={currentStep < i}
          >
            <Step.Content>
              <Step.Title>{step.title}</Step.Title>
              <Step.Description>{step.description}</Step.Description>
            </Step.Content>
          </Step>
        ))}
      </Step.Group>
    </Container>
  );
};

export default connect((state) => ({ props: state }))(AddKnowledge);
