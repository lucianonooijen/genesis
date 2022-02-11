import React, { useContext, useState } from "react";
import { passwordResetComplete } from "@genesis/api";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "components/Input/TextInput/TextInput";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import { usePasswordValidation } from "components/Input/TextInput/TextInput.validation";
import ErrorBanner from "components/ErrorBanner/ErrorBanner";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import {
    logPasswordResetCompleteError,
    logPasswordResetCompleteFinish,
} from "data/analytics/analytics";
import { PasswordForgotCompleteProps } from "./PasswordForgotComplete.types";

const PasswordForgotComplete: React.FC<PasswordForgotCompleteProps> = ({
    passwordResetCompleteApiCall = passwordResetComplete,
}) => {
    const appState = useContext(AppStateContext);
    const [error, setError] = useState<Error | null>(null);
    const [resetCode, setResetCode] = useState("");
    const [password, setPassword] = useState("");
    const [passwordError, validatePassword] = usePasswordValidation();

    const canSubmit = resetCode && password && !passwordError;

    const submit = async () => {
        setError(null);
        const config = getApiConfig(appState);
        try {
            const res = await passwordResetCompleteApiCall(config, {
                resetToken: resetCode,
                password,
            });
            logPasswordResetCompleteFinish();
            appState.setJwt(res.jwt);
        } catch (e) {
            logPasswordResetCompleteError(e);
            setError(e as Error);
        }
    };

    return (
        <PaddedEmptyLayout>
            <ErrorBanner error={error} />
            <Title>Change your password</Title>
            <SubTitle>
                Fill the code from your email in this form and set your new
                password
            </SubTitle>
            <TextInput label="Reset code" onChange={setResetCode} />
            <TextInput
                type={InputFieldType.Password}
                label="Password"
                onChange={setPassword}
                validatorFunc={validatePassword}
                validationError={passwordError}
            />
            <ButtonPrimary
                title="Save my new password"
                onPress={submit}
                disabled={!canSubmit}
            />
        </PaddedEmptyLayout>
    );
};

export default PasswordForgotComplete;
