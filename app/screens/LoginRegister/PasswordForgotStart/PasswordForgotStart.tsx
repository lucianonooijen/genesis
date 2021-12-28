import React from "react";
import { StackNavigationProps } from "types/Navigation";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";

const PasswordForgotStart: React.FC<StackNavigationProps> = ({
    navigation,
}) => {
    return (
        <PaddedEmptyLayout>
            <Title>Reset password</Title>
            <SubTitle>Enter your email to reset your password</SubTitle>
            <ButtonPrimary
                title="Login"
                onPress={() =>
                    navigation.navigate(
                        LoginRegisterScreens.PasswordForgotComplete,
                    )
                }
            />
        </PaddedEmptyLayout>
    );
};

export default PasswordForgotStart;
