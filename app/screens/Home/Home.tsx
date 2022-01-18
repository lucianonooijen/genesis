import React from "react";
import { Text, View } from "react-native";
import { StackNavigationProps } from "types/Navigation";

const Home: React.FC<StackNavigationProps> = () => {
    return (
        <View
            style={{ flex: 1, alignItems: "center", justifyContent: "center" }}
        >
            <Text>Home Screen</Text>
        </View>
    );
};

export default Home;
