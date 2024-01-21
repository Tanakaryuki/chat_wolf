import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import styles from "./main.module.css";
import { BrowserRouter } from "react-router-dom";

// biome-ignore lint/style/noNonNullAssertion: <explanation>
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <BrowserRouter>
      <div className={styles.container}>
        <App />
      </div>
    </BrowserRouter>
  </React.StrictMode>
);
