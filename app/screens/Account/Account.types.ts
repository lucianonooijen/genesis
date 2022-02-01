import { DeleteApiCall, GetApiCall, PutApiCall } from "@genesis/api/types";
import { DeleteAccountRequest, UserProfile } from "@genesis/api/types/Profile";

export interface AccountProps {
    getUserProfile?: GetApiCall<UserProfile>;
    updateUserProfile?: PutApiCall<UserProfile, UserProfile>;
    deleteAccount?: DeleteApiCall<DeleteAccountRequest, null>;
}
