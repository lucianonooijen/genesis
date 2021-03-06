import * as React from "react";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";

// Tutorial screens
import {
    TutorialScreenOne,
    TutorialScreenThree,
    TutorialScreenTwo,
} from "screens/Tutorial/Tutorial";

// Login and register routes
import LoginRegisterLanding from "screens/LoginRegister/LoginRegisterLanding/LoginRegisterLanding";
import Login from "screens/LoginRegister/Login/Login";
import Register from "screens/LoginRegister/Register/Register";
import PasswordForgotStart from "screens/LoginRegister/PasswordForgotStart/PasswordForgotStart";
import PasswordForgotComplete from "screens/LoginRegister/PasswordForgotComplete/PasswordForgotComplete";

// Main screens
import Home from "screens/Home/Home";
import Account from "screens/Account/Account";
import { RouterProps } from "./Router.types";
import { LoginRegisterScreens, MainScreens, TutorialScreens } from "./types";

const Stack = createNativeStackNavigator();
const Tab = createBottomTabNavigator();

const Router: React.FC<RouterProps> = ({ appState }) => {
    return (
        <NavigationContainer>
            <Stack.Navigator
                screenOptions={{
                    headerShown: false,
                }}
            >
                {tutorialStack(appState.hasSeenTutorial)}
                {loginRegisterStack(appState.jwt)}
                <Stack.Screen name={MainScreens.Home} component={MainTabs} />
            </Stack.Navigator>
        </NavigationContainer>
    );
};

export default Router;

const tutorialStack = (hasSeenTutorial: boolean) => {
    if (hasSeenTutorial) {
        return null;
    }
    return (
        <>
            <Stack.Screen
                name={TutorialScreens.ScreenOne}
                component={TutorialScreenOne}
            />
            <Stack.Screen
                name={TutorialScreens.ScreenTwo}
                component={TutorialScreenTwo}
            />
            <Stack.Screen
                name={TutorialScreens.ScreenThree}
                component={TutorialScreenThree}
            />
        </>
    );
};

const loginRegisterStack = (jwt: string | null) => {
    if (jwt) {
        return null;
    }
    return (
        <>
            <Stack.Screen
                name={LoginRegisterScreens.LoginRegisterLanding}
                component={LoginRegisterLanding}
            />
            <Stack.Screen name={LoginRegisterScreens.Login} component={Login} />
            <Stack.Screen
                name={LoginRegisterScreens.Register}
                component={Register}
            />
            <Stack.Screen
                name={LoginRegisterScreens.PasswordForgotStart}
                component={PasswordForgotStart}
            />
            <Stack.Screen
                name={LoginRegisterScreens.PasswordForgotComplete}
                component={PasswordForgotComplete}
            />
        </>
    );
};

const MainTabs = () => (
    <Tab.Navigator>
        <Tab.Screen name="Home" component={Home} />
        <Tab.Screen name="Account" component={Account} />
    </Tab.Navigator>
);
