import React, { useContext, useEffect, useState } from "react";
import { register } from "@genesis/api";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "components/Input/TextInput/TextInput";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import ErrorBanner from "components/ErrorBanner/ErrorBanner";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import {
    logRegisterComplete,
    logRegisterError,
    logRegisterOpen,
} from "data/analytics/analytics";
import {
    useEmailValidation,
    usePasswordValidation,
} from "components/Input/TextInput/TextInput.validation";
import { RegisterProps } from "./Register.types";

const Register: React.FC<RegisterProps> = ({
    navigation,
    registerApiCall = register,
}) => {
    const appState = useContext(AppStateContext);
    const [error, setError] = useState<Error | null>(null);
    const [name, setName] = useState("");

    const [email, setEmail] = useState("");
    const [emailError, validateEmail] = useEmailValidation();

    const [password, setPassword] = useState("");
    const [passwordError, validatePassword] = usePasswordValidation();

    const canSubmitForm =
        name && email && password && !emailError && !passwordError;

    useEffect(logRegisterOpen, []);

    const submit = async () => {
        setError(null);
        const config = getApiConfig(appState);
        try {
            const res = await registerApiCall(config, {
                email,
                password,
                firstName: name,
            });
            appState.setJwt(res.jwt);
            logRegisterComplete(email);
        } catch (e) {
            logRegisterError(e);
            setError(e as Error);
        }
    };

    return (
        <PaddedEmptyLayout>
            <ErrorBanner error={error} />
            <Title>Register</Title>
            <SubTitle>Create an account using the form below</SubTitle>
            <TextInput label="Name" onChange={setName} />
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
                title="Register"
                onPress={submit}
                disabled={!canSubmitForm}
            />
            <ButtonPrimary
                title="I already have an account"
                onPress={() => navigation.navigate(LoginRegisterScreens.Login)}
            />
        </PaddedEmptyLayout>
    );
};

export default Register;
