import React, { useState } from "react";
import { StackNavigationProps } from "types/Navigation";
import { MainScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "../../../components/Input/TextInput/TextInput";
import { InputFieldType } from "../../../components/Input/TextInput/TextInput.types";

const PasswordForgotComplete: React.FC<StackNavigationProps> = ({
    navigation,
}) => {
    const [resetCode, setResetCode] = useState("");
    const [password, setPassword] = useState("");
    const canSubmit = resetCode && password;
    const submit = () => {
        console.log("password reset complete:");
        console.log(`resetCode: ${resetCode}`);
        console.log(`password: ${password}`);
        navigation.navigate(MainScreens.Home);
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
                type={InputFieldType.Email}
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
