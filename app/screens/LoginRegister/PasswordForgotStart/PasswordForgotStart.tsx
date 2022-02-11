import React, { useContext, useState } from "react";
import { passwordResetStart } from "@genesis/api";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import { useEmailValidation } from "components/Input/TextInput/TextInput.validation";
import TextInput from "components/Input/TextInput/TextInput";
import ErrorBanner from "components/ErrorBanner/ErrorBanner";
import { getApiConfig } from "data/api/api";
import AppStateContext from "data/AppState/AppState";
import { PasswordForgotStartProps } from "./PasswordForgotStart.types";

const PasswordForgotStart: React.FC<PasswordForgotStartProps> = ({
    navigation,
    passwordResetStartApiCall = passwordResetStart,
}) => {
    const appState = useContext(AppStateContext);
    const [error, setError] = useState<Error | null>(null);
    const [email, setEmail] = useState("");
    const [emailError, validateEmail] = useEmailValidation();

    const submit = async () => {
        setError(null);
        const config = getApiConfig(appState);
        try {
            await passwordResetStartApiCall(config, { email });
            navigation.navigate(LoginRegisterScreens.PasswordForgotComplete);
        } catch (e) {
            setError(e as Error);
        }
    };

    return (
        <PaddedEmptyLayout>
            <ErrorBanner error={error} />
            <Title>Reset password</Title>
            <SubTitle>Enter your email to reset your password</SubTitle>
            <TextInput
                type={InputFieldType.Email}
                label="Email"
                onChange={setEmail}
                validatorFunc={validateEmail}
                validationError={emailError}
            />
            <ButtonPrimary
                title="Reset my password"
                onPress={submit}
                disabled={!email || !!emailError}
            />
        </PaddedEmptyLayout>
    );
};

export default PasswordForgotStart;
