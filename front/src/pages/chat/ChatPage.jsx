import axios from "axios";
import { useState } from "react";
import { connect } from "react-redux";
import { Input } from "semantic-ui-react";
import { env } from "../../global/variables";

import "../../static/scss/ChatPage.scss";
import "../../static/scss/ChatPageResponsive.scss";

const ChatPage = ({ props, dispatch }) => {
  const [messages, setMessages] = useState([]);
  const [inputValue, setInputValue] = useState("");
  const [inputIsFocused] = useState(false);
  const [inputIsEnabled, setInputIsEnabled] = useState(false);
  const senderList = ["BOT", "USER"];

  function capitalizeFirstLetter(string) {
    let res = "";

    for (let i = 0; i < string.length; i++) {
      if (i === 0) {
        res = `${res}${string[i].toUpperCase()}`;
      } else {
        res = `${res}${string[i].toLowerCase()}`;
      }
    }
    return res;
  }

  async function sendMessage(e) {
    if (inputValue) {
      setInputIsEnabled(true);
    } else {
      setInputIsEnabled(false);
    }

    if (e.key === "Enter") {
      if (!inputValue) {
        return;
      }
      const userMessage = {
        sender: senderList[1],
        message: inputValue,
      };
      setMessages((messages) => [...messages, userMessage]);

      await axios
        .get(`${env.conversationBase}/${inputValue}`)
        .then((res) => {
          const botMessage = {
            sender: senderList[0],
            message: res.data.message,
            answer: res.data.answer,
            category: res.data.category,
          };

          setMessages((messages) => [...messages, botMessage]);
          setInputValue("");
        })
        .catch((_) => {
          return;
        });
    }
  }

  return (
    <div className="chat_main">
      <div className="chat_content">
        <div className="messages">
          {messages.map((m, i) =>
            m.sender === senderList[0] ? (
              <div className="botMessageStyle" key={i}>
                <p>{capitalizeFirstLetter(m.answer)}</p>
              </div>
            ) : (
              <div className="userMessageStyle" key={i}>
                <p>{m.message}</p>
              </div>
            )
          )}
        </div>
        {inputIsEnabled ? (
          <Input
            icon={{
              name: "send",
              circular: true,
              link: true,
            }}
            placeholder="Search..."
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            focus={inputIsFocused}
            onKeyDown={sendMessage}
            maxLength={255}
            className={"message_field"}
          />
        ) : (
          <Input
            placeholder="Search..."
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            focus={inputIsFocused}
            onKeyDown={sendMessage}
            maxLength={255}
            className={"message_field"}
          />
        )}
      </div>
    </div>
  );
};

export default connect((state) => ({ props: state }))(ChatPage);
