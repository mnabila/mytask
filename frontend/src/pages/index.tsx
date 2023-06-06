import { createBrowserRouter } from "react-router-dom";
import Login from "./Login";
import Home from "./Home";
import MainLayout from "../layouts";
import Todo from "./Todo";
import UserSetting from "./UserSetting";

const routers = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/dashboard",
    element: <MainLayout />,
    children: [
      {
        path: "",
        element: <Todo />,
      },
      {
        path: "setting",
        element: <UserSetting />,
      },
    ],
  },
]);
export default routers;
