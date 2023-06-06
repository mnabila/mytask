import { useForm } from "react-hook-form";
import { Button, FormInput } from "../../components";

export default function UserSetting() {
  const formUpdatePassword = useForm();
  const formUpdateProfile = useForm();
  return (
    <div className="flex space-x-5">
      <div className="card w-96 bg-base-100 shadow-xl">
        <div className="card-body">
          <h2 className="card-title">Update Profile</h2>
          <form>
            <FormInput
              label="name"
              name="name"
              control={formUpdateProfile.control}
              type="text"
              required={true}
            />
            <FormInput
              label="email"
              name="email"
              control={formUpdateProfile.control}
              type="email"
              required={true}
            />
            <FormInput
              label="Confirm"
              name="password"
              control={formUpdateProfile.control}
              type="password"
              required={true}
            />
            <div className="card-actions justify-end mt-3">
              <Button label="Update" type="submit" />
            </div>
          </form>
        </div>
      </div>
      <div className="card w-96 bg-red-300 shadow-xl">
        <div className="card-body">
          <h2 className="card-title">Update Password</h2>
          <form>
            <FormInput
              label="Old Password"
              name="oldPassword"
              control={formUpdatePassword.control}
              type="password"
              required={true}
            />
            <FormInput
              label="New Password"
              name="newPassword"
              control={formUpdatePassword.control}
              type="password"
              required={true}
            />
            <FormInput
              label="Confirm New Password"
              name="confirmPassword"
              control={formUpdatePassword.control}
              type="password"
              required={true}
            />
            <div className="card-actions justify-end mt-3">
              <Button label="Update" type="submit" />
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
