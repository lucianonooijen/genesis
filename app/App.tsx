import React, { useContext } from "react";
import Router from "router/Router";
import AppStateContext, {
    AppStateContextProvider,
} from "data/AppState/AppState";
import { UserProfileStateContextProvider } from "data/UserProfileState/UserProfileState";
import { configurePushNotifications } from "data/pushNotifications/pushNotifications";

if (!process.env.JEST) {
    // DO NOT USE .configure() INSIDE A COMPONENT, EVEN App
    configurePushNotifications();
}

const App = () => {
    const appState = useContext(AppStateContext);

    return <Router appState={appState} />;
};

const ConnectedApp = () => (
    <AppStateContextProvider>
        <UserProfileStateContextProvider>
            <App />
        </UserProfileStateContextProvider>
    </AppStateContextProvider>
);

export default ConnectedApp;
