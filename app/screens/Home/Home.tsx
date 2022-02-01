import React, { useContext, useEffect } from "react";
import { Text, View } from "react-native";
import { profileGet } from "@genesis/api";
import UserProfileStateContext from "data/UserProfileState/UserProfileState";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import { HomeProps } from "./Home.types";
import { generateLoadUserProfileStateEffect } from "../../data/api/profile";

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

    return (
        <View
            style={{ flex: 1, alignItems: "center", justifyContent: "center" }}
        >
            <Text>Welcome to Home, {userProfileState.profile.firstName}</Text>
        </View>
    );
};

export default Home;
