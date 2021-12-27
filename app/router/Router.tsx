import * as React from "react";
import {NavigationContainer} from "@react-navigation/native";
import {createNativeStackNavigator} from "@react-navigation/native-stack";
import Home from "../pages/Home/Home";
import Landing from "../pages/Landing/Landing";
import {TutorialScreens} from "./types";
import {RouterProps} from "./Router.types";

const Stack = createNativeStackNavigator();

const Router: React.FC<RouterProps> = ({appState}) => {
    return (
        <NavigationContainer>
            <Stack.Navigator
                screenOptions={{
                    headerShown: false,
                }}
            >
                {tutorialStack(appState.hasSeenTutorial)}
                <Stack.Screen name="Home" component={Home} />
            </Stack.Navigator>
        </NavigationContainer>
    );
};

export default Router;

const tutorialStack = (hasSeenTutorial: boolean) => {
    if (hasSeenTutorial) {
        return null;
    }
    return <Stack.Screen name={TutorialScreens.Landing} component={Landing} />;
};
