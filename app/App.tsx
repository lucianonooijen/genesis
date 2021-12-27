import React from "react";
import {
    SafeAreaView,
    StatusBar,
    useColorScheme,
} from "react-native";
import Router from "./router/Router";

const App: React.FC = () => {
    const isDarkMode = useColorScheme() === "dark";
    const appState = {
        hasSeenTutorial: false
    }

    return (
        <SafeAreaView>
            <StatusBar
                barStyle={isDarkMode ? "light-content" : "dark-content"}
            />
            <Router appState={appState}/>
        </SafeAreaView>
    );
};

export default App;
