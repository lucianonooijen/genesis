import generatePostApiCall from "./internal/postapicall";
import { RegisterRequest } from "./types/Register";
import { PostEndpoints } from "./internal/endpoints";
import {
    jwtResponseDecoder,
    loginRequestDecoder,
    passwordResetCompleteRequestDecoder,
    passwordResetStartRequestDecoder,
    registerRequestDecoder,
} from "./internal/decoders";
import { LoginRequest } from "./types/Login";
import { JwtResponse } from "./types/Jwt";
import {
    PasswordResetCompleteRequest,
    PasswordResetStartRequest,
} from "./types/PasswordReset";

export const register = generatePostApiCall<RegisterRequest, JwtResponse>(
    PostEndpoints.CreateAccount,
    false,
    registerRequestDecoder,
    jwtResponseDecoder,
);

export const login = generatePostApiCall<LoginRequest, JwtResponse>(
    PostEndpoints.Login,
    false,
    loginRequestDecoder,
    jwtResponseDecoder,
);

export const passwordResetStart = generatePostApiCall<
    PasswordResetStartRequest,
    null
>(PostEndpoints.PassResetStart, false, passwordResetStartRequestDecoder);

export const passwordResetComplete = generatePostApiCall<
    PasswordResetCompleteRequest,
    JwtResponse
>(
    PostEndpoints.PassResetComplete,
    false,
    passwordResetCompleteRequestDecoder,
    jwtResponseDecoder,
);
