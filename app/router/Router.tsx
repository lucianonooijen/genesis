import * as React from "react";
import {NavigationContainer} from "@react-navigation/native";
import {createNativeStackNavigator} from "@react-navigation/native-stack";
import Home from "../pages/Home/Home";

const Stack = createNativeStackNavigator();

const Router: React.FC = () => {
    return (
        <NavigationContainer>
            <Stack.Navigator>
                <Stack.Screen name="Home" component={Home} />
            </Stack.Navigator>
        </NavigationContainer>
    );
};

export default Router;
