import React from "react";
import { StackNavigationProps } from "types/Navigation";
import { MainScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";

const Register: React.FC<StackNavigationProps> = ({ navigation }) => {
    return (
        <PaddedEmptyLayout>
            <Title>Register</Title>
            <SubTitle>Create an account using the form below</SubTitle>
            <ButtonPrimary
                title="Register"
                onPress={() => navigation.navigate(MainScreens.Home)}
            />
        </PaddedEmptyLayout>
    );
};

export default Register;
