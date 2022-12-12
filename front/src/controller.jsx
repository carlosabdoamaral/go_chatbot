import React, { useEffect } from "react";
import { connect } from "react-redux";
import Home from "./pages/home/home";
import AddKnowledge from "./pages/knowledge/addKnowledge";
import NavbarWidget from "./widgets/navbar";
import ChatPage from "./pages/chat/ChatPage";
import Animation from "./pages/animation/AnimationPage";

// SCSS
import "./static/scss/_fonts.scss";
import "./static/scss/_master.scss";
function EndAnimation() {
  return {
    type: "SET_INITIAL_ANIMATION_END",
  };
}

const Controller = ({ props, dispatch }) => {
  useEffect(() => {
    setTimeout(() => {
      dispatch(EndAnimation());
    }, 6000);
  });

  return (
    <main>
      {props.isShowingFirstAnimation ? (
        <Animation />
      ) : (
        // <div style={MainStyle.main}>
        //   <Container>
        //   </Container>
        // </div>

        <div>
          <NavbarWidget />
          {props.activePathIndex === 0 && <ChatPage />}
          {props.activePathIndex === 1 && <Home />}
          {props.activePathIndex === 2 && <AddKnowledge />}
        </div>
      )}
    </main>
  );
};

export default connect((state) => ({ props: state }))(Controller);
