import * as React from "react";
import Router from "router/Router";
import { useContext } from "react";
import AppStateContext, {
    AppStateContextProvider,
} from "./data/AppState/AppState";
import { UserProfileStateContextProvider } from "./data/UserProfileState/UserProfileState";

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
