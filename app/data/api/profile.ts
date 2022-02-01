import { UserProfileState } from "data/UserProfileState/UserProfileState";
import { GetApiCall } from "@genesis/api/types";
import { UserProfile } from "@genesis/api/types/Profile";
import ApiConfig from "@genesis/api/types/Config";

export const generateLoadUserProfileStateEffect =
    (
        state: UserProfileState,
        apiConfig: ApiConfig,
        apiCall: GetApiCall<UserProfile>,
    ) =>
    () => {
        // note the extra `() =>` to make it a closure
        if (state.hasLoaded) {
            return;
        }

        const loader = async () => {
            try {
                const profile = await apiCall(apiConfig);
                state.setProfile(profile);
            } catch (e) {
                console.warn(e); // eslint-disable-line no-console
            }
        };

        loader();
    };
