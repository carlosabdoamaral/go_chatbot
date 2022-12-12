import { connect } from "react-redux";
import "../../static/scss/AnimationPage.scss";
import "../../static/scss/Stars.scss";

const AnimationPage = ({ props, dispatch }) => {
  return (
    <div className="initial-animation-view">
      <div className="bg-animation">
        <div id="stars"></div>
        <div id="stars2"></div>
        <div id="stars3"></div>
      </div>

      <h3> Say hello to </h3>
      <h1>{props.chatbotName}</h1>
    </div>
  );
};

export default connect((state) => ({ props: state }))(AnimationPage);