import React, { useState } from "react";
import { StackNavigationProps } from "types/Navigation";
import { LoginRegisterScreens, MainScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "components/Input/TextInput/TextInput";
import { InputFieldType } from "../../../components/Input/TextInput/TextInput.types";

const Login: React.FC<StackNavigationProps> = ({ navigation }) => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const canSubmitForm = email && password;

    return (
        <PaddedEmptyLayout>
            <Title>Log in</Title>
            <SubTitle>Enter your account details to log in</SubTitle>
            <TextInput
                type={InputFieldType.Email}
                label="Email"
                onChange={setEmail}
                testID="input-email"
            />
            <TextInput
                type={InputFieldType.Password}
                label="Password"
                onChange={setPassword}
                testID="input-password"
            />
            <ButtonPrimary
                title="Log in"
                onPress={() => navigation.navigate(MainScreens.Home)}
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
        </PaddedEmptyLayout>
    );
};

export default Login;
