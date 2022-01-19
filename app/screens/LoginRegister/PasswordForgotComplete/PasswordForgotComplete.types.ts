import { StackNavigationProps } from "types/Navigation";
import { PostApiCall } from "@genesis/api/types";
import { PasswordResetCompleteRequest } from "@genesis/api/types/PasswordReset";
import { JwtResponse } from "@genesis/api/types/Jwt";

export interface PasswordForgotCompleteProps extends StackNavigationProps {
    passwordResetCompleteApiCall?: PostApiCall<
        PasswordResetCompleteRequest,
        JwtResponse
    >;
}
