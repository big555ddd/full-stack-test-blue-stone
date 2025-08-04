export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface ForgotRequest {
  email: string;
}

export interface VerifyOtpRequest {
  id: string;
  otp: string;
  newPassword: string;
  confirmPassword: string;
}

export interface ForgotResponse {
  code: number;
  message: string;
  data: Data;
}

export interface Data {
  id: string;
  userId: string;
  otp: string;
  expiresAt: number;
  used: boolean;
  createdAt: number;
}
