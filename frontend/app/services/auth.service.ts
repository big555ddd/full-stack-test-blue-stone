import type {
  ForgotRequest,
  LoginRequest,
  RegisterRequest,
  VerifyOtpRequest,
} from "~/models/auth.model";
import { client } from "./httpClient";

export const Login = (body: LoginRequest) => {
  return client({
    url: `/auth/login`,
    method: "post",
    body: body,
  });
};

export const Register = (body: RegisterRequest) => {
  return client({
    url: `/auth/register`,
    method: "post",
    body: body,
  });
};

export const ForgotPassword = (body: ForgotRequest) => {
  return client({
    url: `/auth/forgot-password`,
    method: "post",
    body: body,
  });
};

export const VerifyOtp = (body: VerifyOtpRequest) => {
  return client({
    url: `/auth/verify-otp`,
    method: "post",
    body: body,
  });
};
