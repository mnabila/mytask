import { Fragment } from "react";
import { Navbar } from "../components";
import { Outlet } from "react-router-dom";
import { Toaster } from "react-hot-toast";

export default function MainLayout() {
  return (
    <Fragment>
      <div className="flex flex-col items-center bg-base-300 h-screen space-y-5">
        <Navbar />
        <Outlet />
      </div>
      <Toaster position="top-center" reverseOrder={false} />
    </Fragment>
  );
}
