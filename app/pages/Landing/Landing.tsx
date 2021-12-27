import React from "react";
import {Text, View} from "react-native";
import {ButtonPrimary} from "../../components/Buttons/ButtonRegular/ButtonRegular";
import {TutorialNavigationProps} from "../../types/Navigation";
import {MainScreens} from "../../router/types";

const Landing: React.FC<TutorialNavigationProps> = ({navigation}) => {
    return (
        <View>
            <Text>Landing</Text>
            <ButtonPrimary
                title="test"
                onPress={() => navigation.navigate(MainScreens.Home)}
            />
        </View>
    );
};

export default Landing;
