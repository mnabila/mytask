import { IconUser } from "@tabler/icons-react";
import { useRecoilValue } from "recoil";
import { userState } from "../states";
import { Link } from "react-router-dom";

export default function Navbar() {
  const user = useRecoilValue(userState);
  return (
    <div className="navbar bg-base-100">
      <div className="flex-1">
        <a className="btn btn-ghost normal-case text-xl">My Task</a>
      </div>
      <div className="flex-none gap-2">
        <div className="dropdown dropdown-end">
          <label tabIndex={0} className="btn btn-ghost space-x-2">
            <span className="uppercase">{user.name}</span>
            <IconUser />
          </label>
          <ul
            tabIndex={0}
            className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52"
          >
            <li>
              <a className="justify-between">
                Profile
                <span className="badge">New</span>
              </a>
            </li>
            <li>
              <Link to="/dashboard/setting">Settings</Link>
            </li>
            <li>
              <a>Logout</a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
