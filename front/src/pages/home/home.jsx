import { MessageWidget } from "../../widgets/message";
import React from "react";
import { useState } from "react";
import axios from "axios";
import { Table, Message, Container } from "semantic-ui-react";
import { connect } from "react-redux";
import { env } from "../../global/variables";
import { useEffect } from "react";

const Home = ({ props, dispatch }) => {
  let [didLoad, setDidLoad] = useState(false);
  let [didFail, setDidFail] = useState(false);
  let [list, setList] = useState([]);

  async function getKnowledge() {
    await axios
      .get(env.allKnowledge)
      .then((res) => {
        setDidLoad(true);
        setDidFail(false);
        setList(res.data);
        return;
      })
      .catch((err) => {
        setDidLoad(false);
        setDidFail(true);
        setList([]);
        return;
      });
  }

  useEffect(() => {
    getKnowledge();
  }, []);

  return (
    <Container style={{paddingTop : 30}}>
      {!didLoad && <MessageWidget />}

      {didFail && (
        <Message negative>
          <Message.Header>Important!</Message.Header>
          <p>Something wrong happen</p>
        </Message>
      )}

      {didLoad && (
        <Table celled>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Message</Table.HeaderCell>
              <Table.HeaderCell>Answer</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {list.map((know, i) => (
              <Table.Row key={i}>
                <Table.Cell>{know.message}</Table.Cell>
                <Table.Cell>{know.answer}</Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      )}
    </Container>
  );
};

export default connect((state) => ({ props: state }))(Home);
