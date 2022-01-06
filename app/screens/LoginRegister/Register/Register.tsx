import React, { useContext, useState } from "react";
import { StackNavigationProps } from "types/Navigation";
import { LoginRegisterScreens } from "router/types";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { SubTitle, Title } from "components/Typography/Typography";
import TextInput from "../../../components/Input/TextInput/TextInput";
import { InputFieldType } from "../../../components/Input/TextInput/TextInput.types";
import AppStateContext from "../../../data/AppState/AppState";

const Register: React.FC<StackNavigationProps> = ({ navigation }) => {
    const { setIsLoggedIn } = useContext(AppStateContext);
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const canSubmitForm = name && email && password;

    const submit = () => {
        console.log("login request:");
        console.log(`name: ${name}`);
        console.log(`email: ${email}`);
        console.log(`password: ${password}`);
        setIsLoggedIn(true);
    };

    return (
        <PaddedEmptyLayout>
            <Title>Register</Title>
            <SubTitle>Create an account using the form below</SubTitle>
            <TextInput label="Name" onChange={setName} />
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
                title="Register"
                onPress={submit}
                disabled={!canSubmitForm}
            />
            <ButtonPrimary
                title="I already have an account"
                onPress={() => navigation.navigate(LoginRegisterScreens.Login)}
            />
        </PaddedEmptyLayout>
    );
};

export default Register;
