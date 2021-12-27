import * as React from 'react';
import { NavigationContainer} from "@react-navigation/native";
import {createBottomTabNavigator} from "@react-navigation/bottom-tabs";
import { createStackNavigator } from "@react-navigation/stack";

import {MainScreens, TutorialScreens} from "./types";

import Landing from "../pages/Landing/Landing";
import Home from "../pages/Home/Home";

const Stack = createStackNavigator();
const Tab = createBottomTabNavigator();

interface AppState {
    hasSeenTutorial: boolean
}
interface RouterProps {
    appState: AppState
}

const Router: React.FC<RouterProps> = ({ appState }) => {
    return (
        <NavigationContainer>
            <Stack.Navigator screenOptions={{
                headerShown: false,
            }}>
                {!appState.hasSeenTutorial && (
                    <Stack.Screen
                        name={TutorialScreens.Landing}
                        component={Landing}
                    />
                )}
                <Stack.Screen name={MainScreens.Home} component={Home} />
            </Stack.Navigator>
        </NavigationContainer>
    );
};

export default Router;
