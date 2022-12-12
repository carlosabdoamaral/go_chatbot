import { Header } from "semantic-ui-react";
import { connect } from "react-redux";

const HeaderWidget = ({props}) => {
  return (
    <Header as="h1" content={props.nav_list[props.activePathIndex].title}/>
  )
};

export default connect(state => ({props : state}))(HeaderWidget);
