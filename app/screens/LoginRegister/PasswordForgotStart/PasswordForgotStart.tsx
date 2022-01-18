import React, { useContext, useState } from "react";
import { passwordResetStart } from "@genesis/api";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import TextInput from "components/Input/TextInput/TextInput";
import { getApiConfig } from "data/api/api";
import AppStateContext from "data/AppState/AppState";
import { PasswordForgotStartProps } from "./PasswordForgotStart.types";

const PasswordForgotStart: React.FC<PasswordForgotStartProps> = ({
    navigation,
    apiCall = passwordResetStart,
}) => {
    const appState = useContext(AppStateContext);
    const [email, setEmail] = useState("");

    const submit = async () => {
        const config = getApiConfig(appState);
        try {
            await apiCall(config, { email });
            navigation.navigate(LoginRegisterScreens.PasswordForgotComplete);
        } catch (e) {
            console.warn(e); // eslint-disable-line no-console
        }
    };

    return (
        <PaddedEmptyLayout>
            <Title>Reset password</Title>
            <SubTitle>Enter your email to reset your password</SubTitle>
            <TextInput
                type={InputFieldType.Email}
                label="Email"
                onChange={setEmail}
            />
            <ButtonPrimary
                title="Reset my password"
                onPress={submit}
                disabled={!email}
            />
        </PaddedEmptyLayout>
    );
};

export default PasswordForgotStart;
