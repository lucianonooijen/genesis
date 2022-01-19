import { StackNavigationProps } from "types/Navigation";
import { PostApiCall } from "@genesis/api/types";
import { RegisterRequest } from "@genesis/api/types/Register";
import { JwtResponse } from "@genesis/api/types/Jwt";

export interface RegisterProps extends StackNavigationProps {
    registerApiCall?: PostApiCall<RegisterRequest, JwtResponse>;
}
