import { SubmitHandler, useForm } from "react-hook-form";
import { AuthRequest } from "../../dto";
import { Alert, Button, FormInput } from "../../components";
import { IconLogin } from "@tabler/icons-react";
import useAxios from "../../hooks";
import { useNavigate } from "react-router-dom";
import { Fragment, useState } from "react";
import { useSetRecoilState } from "recoil";
import { userState } from "../../states";

export default function Login() {
  const navigate = useNavigate();
  const setUser = useSetRecoilState(userState);
  const [message, setMessage] = useState("");
  const axios = useAxios();

  const { handleSubmit, control } = useForm<AuthRequest>();

  const onSubmit: SubmitHandler<AuthRequest> = (data) => {
    axios
      .post("/user/login", data)
      .then((res) => {
        setUser(res.data.data);
        navigate("/dashboard");
      })
      .catch((err) => setMessage(err.response.data.message));
  };

  return (
    <Fragment>
      <div className="hero min-h-screen bg-base-200">
        <div className="hero-content flex-col lg:flex-row">
          <div className="text-center lg:text-left">
            <h1 className="text-5xl font-bold">Login now!</h1>
            <p className="py-6">
              Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
              excepturi exercitationem quasi. In deleniti eaque aut repudiandae
              et a id nisi.
            </p>
          </div>
          <div className="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
            <div className="card-body">
              <form onSubmit={handleSubmit(onSubmit)}>
                <FormInput
                  label="Email"
                  name="email"
                  type="email"
                  control={control}
                  required
                />
                <FormInput
                  label="Password"
                  name="password"
                  type="password"
                  control={control}
                  required
                />
                <div className="form-control mt-3">
                  <Alert message={message} />
                </div>
                <div className="form-control mt-6">
                  <Button
                    icon={<IconLogin />}
                    label="Login"
                    className="btn-primary"
                  />
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </Fragment>
  );
}
