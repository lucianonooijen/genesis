import React from "react";
import { Text, View } from "react-native";
import { ButtonPrimary } from "components/Buttons/ButtonRegular/ButtonRegular";
import { TutorialScreens } from "router/types";
import { StackNavigationProps } from "types/Navigation";

const Home: React.FC<StackNavigationProps> = ({ navigation }) => {
    return (
        <View
            style={{ flex: 1, alignItems: "center", justifyContent: "center" }}
        >
            <Text>Home Screen</Text>
            <ButtonPrimary
                title="Back to tutorials"
                onPress={() => navigation.navigate(TutorialScreens.ScreenOne)}
            />
        </View>
    );
};

export default Home;
