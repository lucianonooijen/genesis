import React from "react";
import {Text, View} from "react-native";
import {ButtonPrimary} from "../../components/Buttons/ButtonRegular/ButtonRegular";
import {StackNavigationProps} from "../../types/Navigation";
import {MainScreens} from "../../router/types";

const Tutorial: React.FC<StackNavigationProps> = ({navigation}) => {
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

export default Tutorial;
