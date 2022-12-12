import { Menu, Segment } from "semantic-ui-react";

import { connect } from "react-redux";

function ChangeRoute(index) {
  return {
    type: "SET_ROUTE_ACTIVE",
    route_index: index,
  };
}

const style = {
  segment : {
    margin : 0,
    borderRadius : 0,
    zIndex : 1,
  }
}
const NavbarWidget = ({ props, dispatch }) => {
  return (
    <Segment inverted style={style.segment}>
      <Menu inverted pointing secondary>
        {props.nav_list.map((section, i) => (
          <Menu.Item
            key={i}
            name={section.title}
            active={props.activePathIndex === i}
            onClick={() => dispatch(ChangeRoute(i))}
          />
        ))}
      </Menu>
    </Segment>
  );
};

export default connect((state) => ({ props: state }))(NavbarWidget);
