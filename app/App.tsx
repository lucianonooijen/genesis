import * as React from "react";
import Router from "router/Router";
import { useContext } from "react";
import AppStateContext, {
    AppStateContextProvider,
} from "data/AppState/AppState";
import { UserProfileStateContextProvider } from "data/UserProfileState/UserProfileState";
import { configurePushNotifications } from "data/pushNotifications/pushNotifications";
import FatalError from "./screens/FatalError/FatalError";

if (!process.env.JEST) {
    // DO NOT USE .configure() INSIDE A COMPONENT, EVEN App
    configurePushNotifications();
}

export const App = () => {
    const appState = useContext(AppStateContext);

    if (appState.isLoading) {
        // TODO: Add loading animation or something
        return null;
    }

    if (appState.fatalError) {
        const fe = appState.fatalError;
        return <FatalError title={fe.title} description={fe.description} />;
    }

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
