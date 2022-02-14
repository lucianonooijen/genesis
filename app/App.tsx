import React, { useContext, useEffect } from "react";
import { useNetInfo } from "@react-native-community/netinfo";
import Router from "router/Router";
import AppStateContext, {
    AppStateContextProvider,
} from "data/AppState/AppState";
import { UserProfileStateContextProvider } from "data/UserProfileState/UserProfileState";
import { configurePushNotifications } from "data/pushNotifications/pushNotifications";
import { useAppVersionCheck } from "data/api/versioncheck";
import { logGeneralAppStart } from "data/analytics/analytics";
import { initAnalytics } from "data/analytics/init";
import FatalError from "screens/FatalError/FatalError";

if (!process.env.JEST) {
    // DO NOT USE .configure() INSIDE A COMPONENT, EVEN App
    configurePushNotifications();
}

export const App = () => {
    const netInfo = useNetInfo();
    const appState = useContext(AppStateContext);
    const [hasCheckedVersion, versionError, checkVersion] =
        useAppVersionCheck();

    useEffect(() => {
        if (!process.env.JEST) {
            checkVersion();
        }
    }, [checkVersion]);

    useEffect(logGeneralAppStart, []);

    useEffect(initAnalytics, []);

    if (netInfo.isInternetReachable === false) {
        // must check for === false, as it can also be null if the state is unknown
        return (
            <FatalError
                title="Geen internetverbinding"
                description="Genesis werkt alleen met internetverbinding. Verbind je telefoon met het internet om de app te gebruiken."
            />
        );
    }

    if (appState.isLoading || !hasCheckedVersion) {
        // TODO: Add loading animation or something
        return null;
    }

    if (versionError) {
        return (
            <FatalError
                title={versionError.title}
                description={versionError.description}
            />
        );
    }

    if (appState.fatalError) {
        return (
            <FatalError
                title={appState.fatalError.title}
                description={appState.fatalError.description}
            />
        );
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
