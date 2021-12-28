import React from "react";
import { StackNavigationProps } from "types/Navigation";
import { LoginRegisterScreens, MainScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";

const Login: React.FC<StackNavigationProps> = ({ navigation }) => {
    return (
        <PaddedEmptyLayout>
            <Title>Log in</Title>
            <SubTitle>Enter your account details to log in</SubTitle>
            <ButtonPrimary
                title="Log in"
                onPress={() => navigation.navigate(MainScreens.Home)}
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
