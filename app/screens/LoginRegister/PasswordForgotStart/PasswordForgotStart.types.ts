import { StackNavigationProps } from "types/Navigation";
import { PostApiCall } from "@genesis/api/types";
import { PasswordResetStartRequest } from "@genesis/api/types/PasswordReset";

export interface PasswordForgotStartProps extends StackNavigationProps {
    apiCall?: PostApiCall<PasswordResetStartRequest, null>;
}
