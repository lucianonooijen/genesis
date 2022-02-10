import React, { useContext } from "react";
import { useNetInfo } from "@react-native-community/netinfo";
import Router from "router/Router";
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
    const netInfo = useNetInfo();
    const appState = useContext(AppStateContext);

    if (netInfo.isInternetReachable === false) {
        // must check for === false, as it can also be null if the state is unknown
        return (
            <FatalError
                title="Geen internetverbinding"
                description="Genesis werkt alleen met internetverbinding. Verbind je telefoon met het internet om de app te gebruiken."
            />
        );
    }

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
