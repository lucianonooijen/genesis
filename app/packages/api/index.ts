import { DeleteAccountRequest, UserProfile } from "@genesis/api/types/Profile";
import ApiError from "@genesis/api/types/ApiError";
import generatePostApiCall from "./internal/apicall-post";
import generateGetApiCall from "./internal/apicall-get";
import generatePutApiCall from "./internal/apicall-put";
import generateDeleteApiCall from "./internal/apicall-delete";
import {
    DeleteEndpoints,
    GetEndpoints,
    PostEndpoints,
    PutEndpoints,
} from "./internal/endpoints";

import {
    deleteAccountRequestDecoder,
    jwtResponseDecoder,
    loginRequestDecoder,
    passwordResetCompleteRequestDecoder,
    passwordResetStartRequestDecoder,
    registerRequestDecoder,
    userProfileDecoder,
    registerTokenRequestDecoder,
} from "./internal/decoders";
import { RegisterRequest } from "./types/Register";
import { LoginRequest } from "./types/Login";
import { JwtResponse } from "./types/Jwt";
import {
    PasswordResetCompleteRequest,
    PasswordResetStartRequest,
} from "./types/PasswordReset";
import { RegisterTokenRequest } from "./types/Notification";

export const versionCheck = generateGetApiCall<null>(
    GetEndpoints.VersionCheck,
    false,
);

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

export const profileGet = generateGetApiCall<UserProfile>(
    GetEndpoints.UserProfile,
    true,
    userProfileDecoder,
);

export const profileUpdate = generatePutApiCall<UserProfile, UserProfile>(
    PutEndpoints.UserProfile,
    true,
    userProfileDecoder,
    userProfileDecoder,
);

export const pushNotificationRegisterToken = generatePostApiCall<
    RegisterTokenRequest,
    null
>(PostEndpoints.PushNotifications, true, registerTokenRequestDecoder);

export const accountDelete = generateDeleteApiCall<DeleteAccountRequest, null>(
    DeleteEndpoints.UserProfile,
    true,
    deleteAccountRequestDecoder,
);

export class GenesisApiError extends Error {
    public err: ApiError;

    constructor(err: ApiError) {
        super(`Error ${err.status}: ${err.title}`);

        this.err = err;
    }
}
