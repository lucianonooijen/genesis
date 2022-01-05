import React, { useContext, useState } from "react";
import { StackNavigationProps } from "types/Navigation";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "components/Input/TextInput/TextInput";
import { InputFieldType } from "components/Input/TextInput/TextInput.types";
import AppStateContext from "data/AppState/AppState";

const Login: React.FC<StackNavigationProps> = ({ navigation }) => {
    const { setIsLoggedIn } = useContext(AppStateContext);
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const canSubmitForm = email && password;

    const submit = () => {
        console.log("login request:");
        console.log(`email: ${email}`);
        console.log(`password: ${password}`);
        setIsLoggedIn(true);
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
