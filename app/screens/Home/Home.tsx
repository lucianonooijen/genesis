import React, { useContext, useEffect } from "react";
import { Text, View } from "react-native";
import { profileGet } from "@genesis/api";
import UserProfileStateContext from "data/UserProfileState/UserProfileState";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { requestPushPermissions } from "data/pushNotifications/pushNotifications";
import { generateLoadUserProfileStateEffect } from "data/api/profile";
import { logPushNotificationsRequest } from "data/analytics/analytics";
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
        logPushNotificationsRequest();
        requestPushPermissions(); // Note: must be done at a moment where an accurate JWT is set in local storage.
    };

    const nuke = () => {
        appState.reset();

        appState.setFatalError(
            "De schijt heeft de ventilator geraakt",
            "OOPSIE WOOPSIE!! Uwu We make a fucky wucky!! A wittle fucko boingo! The code monkeys at our headquarters are working VEWY HAWD to fix this!",
        );
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
            <ButtonPrimary title="Nuke app" onPress={nuke} />
        </View>
    );
};

export default Home;
