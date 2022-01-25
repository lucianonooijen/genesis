import { StackNavigationProps } from "types/Navigation";
import { PostApiCall } from "@genesis/api/types";
import { LoginRequest } from "@genesis/api/types/Login";
import { JwtResponse } from "@genesis/api/types/Jwt";

export interface LoginProps extends StackNavigationProps {
    loginApiCall?: PostApiCall<LoginRequest, JwtResponse>;
}
