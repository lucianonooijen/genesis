import React, { useContext, useEffect } from "react";
import { Text, View } from "react-native";
import { profileGet } from "@genesis/api";
import UserProfileStateContext from "data/UserProfileState/UserProfileState";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { requestPushPermissions } from "data/pushNotifications/pushNotifications";
import { generateLoadUserProfileStateEffect } from "data/api/profile";
import { HomeProps } from "./Home.types";

const Home: React.FC<HomeProps> = ({ getUserProfile = profileGet }) => {
    const userProfileState = useContext(UserProfileStateContext);
    const appState = useContext(AppStateContext);
    const apiConfig = getApiConfig(appState);

    // If the userProfileState has not loaded, load it
    // eslint-disable-next-line react-hooks/exhaustive-deps
    useEffect(
        generateLoadUserProfileStateEffect(
            userProfileState,
            apiConfig,
            getUserProfile,
        ),
        [appState, getUserProfile, userProfileState],
    );

    const onPressPushNotifications = () => {
        requestPushPermissions(); // Note: must be done at a moment where an accurate JWT is set in local storage.
    };

    return (
        <View
            style={{ flex: 1, alignItems: "center", justifyContent: "center" }}
        >
            <Text>Welcome to Home, {userProfileState.profile.firstName}</Text>
            <ButtonPrimary
                title="I want push notifications"
                onPress={onPressPushNotifications}
            />
        </View>
    );
};

export default Home;
