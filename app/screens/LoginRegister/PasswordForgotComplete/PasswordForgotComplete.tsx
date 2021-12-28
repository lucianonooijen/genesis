import React from "react";
import { StackNavigationProps } from "types/Navigation";
import { MainScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";

const PasswordForgotComplete: React.FC<StackNavigationProps> = ({
    navigation,
}) => {
    return (
        <PaddedEmptyLayout>
            <Title>Change your password</Title>
            <SubTitle>
                Fill the code from your email in this form and set your new
                password
            </SubTitle>
            <ButtonPrimary
                title="Save password"
                onPress={() => navigation.navigate(MainScreens.Home)}
            />
        </PaddedEmptyLayout>
    );
};

export default PasswordForgotComplete;
