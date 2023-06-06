import { atom } from "recoil";
import { UserResponse } from "../dto";
import { recoilPersist } from "recoil-persist";

const { persistAtom } = recoilPersist();

export const tokenState = atom({
  key: "tokenState",
  default: "",
});

export const userState = atom<UserResponse>({
  key: "userState",
  default: undefined,
  effects: [persistAtom],
});
