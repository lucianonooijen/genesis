import React, { useContext, useState } from "react";
import { login } from "@genesis/api";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "components/Input/TextInput/TextInput";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import ErrorBanner from "components/ErrorBanner/ErrorBanner";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import { LoginProps } from "./Login.types";
import {
    useEmailValidation,
    usePasswordValidation,
} from "../../../components/Input/TextInput/TextInput.validation";

const Login: React.FC<LoginProps> = ({ navigation, loginApiCall = login }) => {
    const appState = useContext(AppStateContext);
    const [error, setError] = useState<Error | null>(null);

    const [email, setEmail] = useState("");
    const [emailError, validateEmail] = useEmailValidation();

    const [password, setPassword] = useState("");
    const [passwordError, validatePassword] = usePasswordValidation();

    const canSubmitForm = email && password && !emailError && !passwordError;

    const submit = async () => {
        setError(null);
        const config = getApiConfig(appState);
        try {
            const res = await loginApiCall(config, { email, password });
            appState.setJwt(res.jwt);
        } catch (e) {
            setError(e as Error);
        }
    };

    return (
        <PaddedEmptyLayout>
            <ErrorBanner error={error} />
            <Title>Log in</Title>
            <SubTitle>Enter your account details to log in</SubTitle>
            <TextInput
                type={InputFieldType.Email}
                label="Email"
                onChange={setEmail}
                validatorFunc={validateEmail}
                validationError={emailError}
            />
            <TextInput
                type={InputFieldType.Password}
                label="Password"
                onChange={setPassword}
                validatorFunc={validatePassword}
                validationError={passwordError}
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
