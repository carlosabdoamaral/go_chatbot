import { configureStore } from "@reduxjs/toolkit";

const INITIAL_STATE = {
  chatbotName : "Raven",
  isShowingFirstAnimation: true,
  // isShowingFirstAnimation: false, //Remover
  activePathIndex: 0,
  nav_list: [
    {
      title: "Chat",
      path: "/",
    },

    {
      title: "Knowledge",
      path: "/",
    },

    {
      title: "Add Knowledge",
      path: "/",
    },
  ],
};

function reducer(state = INITIAL_STATE, action) {
  if (action.type === "SET_ROUTE_ACTIVE") {
    return {
      ...state,
      activePathIndex: action.route_index,
    };
  }

  if (action.type === "SET_INITIAL_ANIMATION_END") {
    return {
      ...state,
      isShowingFirstAnimation: false,
    };
  }

  return state;
}

const store = configureStore({ reducer: reducer });
export default store;
