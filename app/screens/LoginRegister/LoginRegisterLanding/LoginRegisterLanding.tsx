import React from "react";
import { StackNavigationProps } from "types/Navigation";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";

const LoginRegisterLanding: React.FC<StackNavigationProps> = ({
    navigation,
}) => {
    return (
        <PaddedEmptyLayout>
            <Title>Welcome</Title>
            <SubTitle>
                Would you like to log in, or would you like to create an
                account?
            </SubTitle>
            <ButtonPrimary
                title="Register"
                onPress={() =>
                    navigation.navigate(LoginRegisterScreens.Register)
                }
            />
            <ButtonPrimary
                title="Login"
                onPress={() => navigation.navigate(LoginRegisterScreens.Login)}
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

export default LoginRegisterLanding;
