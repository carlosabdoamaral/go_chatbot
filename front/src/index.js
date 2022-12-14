import React from "react";
import ReactDOM from "react-dom/client";
import reportWebVitals from "./reportWebVitals";
import "semantic-ui-css/semantic.min.css";
import { Route, Routes, BrowserRouter } from "react-router-dom";
import Controller from "./controller";
import { Provider } from "react-redux";
import store from "./store";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <Provider store={store}>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Controller/>}/>
      </Routes>
    </BrowserRouter>
  </Provider>
);

reportWebVitals();
