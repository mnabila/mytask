import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { RecoilRoot } from "recoil";
import { RouterProvider } from "react-router-dom";
import routers from "./pages";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RecoilRoot>
      <RouterProvider router={routers} />
    </RecoilRoot>
  </React.StrictMode>
);
