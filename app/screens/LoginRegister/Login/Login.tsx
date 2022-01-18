import React, { useContext, useState } from "react";
import { login } from "@genesis/api";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "components/Input/TextInput/TextInput";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import { LoginProps } from "./Login.types";

const Login: React.FC<LoginProps> = ({ navigation, apiCall = login }) => {
    const appState = useContext(AppStateContext);
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const canSubmitForm = email && password;

    const submit = async () => {
        const config = getApiConfig(appState);
        try {
            const res = await apiCall(config, { email, password });
            appState.setJwt(res.jwt);
        } catch (e) {
            console.warn(e); // eslint-disable-line no-console
        }
    };

    return (
        <PaddedEmptyLayout>
            <Title>Log in</Title>
            <SubTitle>Enter your account details to log in</SubTitle>
            <TextInput
                type={InputFieldType.Email}
                label="Email"
                onChange={setEmail}
            />
            <TextInput
                type={InputFieldType.Password}
                label="Password"
                onChange={setPassword}
            />
            <ButtonPrimary
                title="Log in"
                onPress={submit}
                disabled={!canSubmitForm}
            />
            <ButtonPrimary
                title="Forgot password"
                onPress={() =>
                    navigation.navigate(
                        LoginRegisterScreens.PasswordForgotStart,
                    )
                }
            />
            <ButtonPrimary
                title="Create an account"
                onPress={() =>
                    navigation.navigate(LoginRegisterScreens.Register)
                }
            />
        </PaddedEmptyLayout>
    );
};

export default Login;
