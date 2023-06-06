import axios from "axios";
import { useRecoilState } from "recoil";
import { tokenState } from "../states";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";

const useAxios = () => {
	const [token, setToken] = useRecoilState(tokenState);
	const navigate = useNavigate();

	const instance = axios.create({
		baseURL: import.meta.env.VITE_REST_API,
	});

	// Add a request interceptor
	instance.interceptors.request.use(
		(config) => {
			if (token != "" && config.headers) {
				config.headers["Authorization"] = token;
			}
			return config;
		},
		(error) => Promise.reject(error)
	);

	// Add a response interceptor
	instance.interceptors.response.use(
		(response) => {
			const authorization = response.headers.authorization;
			if (authorization) {
				setToken(authorization);
			}
			return response;
		},
		(error) => {
			if (error.response.status == 401) {
				toast.error(error.response.data.message);
				navigate("/login");
			}
			return Promise.reject(error);
		}
	);

	return instance;
};
export default useAxios;
