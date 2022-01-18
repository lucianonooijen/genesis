import React, { useContext, useState } from "react";
import { passwordResetComplete } from "@genesis/api";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "components/Input/TextInput/TextInput";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import AppStateContext from "data/AppState/AppState";
import { getApiConfig } from "data/api/api";
import { PasswordForgotCompleteProps } from "./PasswordForgotComplete.types";

const PasswordForgotComplete: React.FC<PasswordForgotCompleteProps> = ({
    apiCall = passwordResetComplete,
}) => {
    const appState = useContext(AppStateContext);
    const [resetCode, setResetCode] = useState("");
    const [password, setPassword] = useState("");
    const canSubmit = resetCode && password;

    const submit = async () => {
        const config = getApiConfig(appState);
        try {
            const res = await apiCall(config, {
                resetToken: resetCode,
                password,
            });
            appState.setJwt(res.jwt);
        } catch (e) {
            console.warn(e); // eslint-disable-line no-console
        }
    };

    return (
        <PaddedEmptyLayout>
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
