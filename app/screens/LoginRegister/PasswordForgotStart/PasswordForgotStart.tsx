import React, { useState } from "react";
import { StackNavigationProps } from "types/Navigation";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import { InputFieldType } from "../../../components/Input/TextInput/TextInput.types";
import TextInput from "../../../components/Input/TextInput/TextInput";

const PasswordForgotStart: React.FC<StackNavigationProps> = ({
    navigation,
}) => {
    const [email, setEmail] = useState("");
    const submit = () => {
        console.log("password reset start:");
        console.log(`email: ${email}`);
        navigation.navigate(LoginRegisterScreens.PasswordForgotComplete);
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
