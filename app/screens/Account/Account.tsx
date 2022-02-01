import React, { useContext, useEffect, useState } from "react";
import { Text } from "react-native";
import RNRestart from "react-native-restart";
import { accountDelete, profileGet, profileUpdate } from "@genesis/api";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import UserProfileStateContext from "data/UserProfileState/UserProfileState";
import AppStateContext from "data/AppState/AppState";
import { generateLoadUserProfileStateEffect } from "data/api/profile";
import { getApiConfig } from "data/api/api";
import ErrorBanner from "components/ErrorBanner/ErrorBanner";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import TextInput from "components/Input/TextInput/TextInput";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { AccountProps } from "./Account.types";

const Account: React.FC<AccountProps> = ({
    getUserProfile = profileGet,
    updateUserProfile = profileUpdate,
    deleteAccount = accountDelete,
}) => {
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

    const [error, setError] = useState<Error | null>(null);
    const [firstName, setFirstName] = useState("");
    const [password, setPassword] = useState("");

    const updateProfile = async () => {
        try {
            const profile = await updateUserProfile(apiConfig, {
                firstName,
            });
            userProfileState.setProfile(profile);
        } catch (e) {
            setError(e as Error);
        }
    };

    const deleteProfile = async () => {
        try {
            await deleteAccount(apiConfig, { password });
            setError(new Error("Account has been deleted, resetting the app"));
            setTimeout(RNRestart.Restart, 1000);
        } catch (e) {
            setError(e as Error);
        }
    };

    if (!userProfileState.hasLoaded) {
        return null;
    }

    return (
        <PaddedEmptyLayout>
            <ErrorBanner error={error} />
            <Text>Update profile</Text>
            <TextInput
                label="First name"
                onChange={setFirstName}
                initialValue={userProfileState.profile.firstName}
            />
            <ButtonPrimary
                title="Save profile"
                onPress={updateProfile}
                disabled={firstName === ""}
            />
            <Text>Delete account</Text>
            <TextInput
                type={InputFieldType.Password}
                label="Password"
                onChange={setPassword}
            />
            <ButtonPrimary
                title="Delete account"
                onPress={deleteProfile}
                disabled={password === ""}
            />
        </PaddedEmptyLayout>
    );
};

export default Account;
