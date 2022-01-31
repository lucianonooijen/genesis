import { GetApiCall } from "@genesis/api/types";
import { UserProfile } from "@genesis/api/types/Profile";

export interface HomeProps {
    getUserProfile?: GetApiCall<UserProfile>;
}
