import React, { useContext, useEffect } from "react";
import { Text, View } from "react-native";
import { profileGet } from "@genesis/api";
import UserProfileStateContext from "data/UserProfileState/UserProfileState";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import { HomeProps } from "./Home.types";

const Home: React.FC<HomeProps> = ({ getUserProfile = profileGet }) => {
    const userProfileState = useContext(UserProfileStateContext);
    const appState = useContext(AppStateContext);

    // If the userProfileState has not loaded, load it
    useEffect(() => {
        const load = async () => {
            if (!userProfileState.hasLoaded) {
                try {
                    const profile = await getUserProfile(
                        getApiConfig(appState),
                    );
                    userProfileState.setProfile(profile);
                } catch (e) {
                    console.error(e); // eslint-disable-line no-console
                }
            }
        };
        load();
    }, [appState, getUserProfile, userProfileState]);

    return (
        <View
            style={{ flex: 1, alignItems: "center", justifyContent: "center" }}
        >
            <Text>Welcome to Home, {userProfileState.profile.firstName}</Text>
        </View>
    );
};

export default Home;
