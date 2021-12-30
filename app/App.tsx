import * as React from "react";
import Router from "router/Router";
import { useContext } from "react";
import AppStateContext, {
    AppStateContextProvider,
} from "./data/AppState/AppState";

const App = () => {
    const appState = useContext(AppStateContext);

    return <Router appState={appState} />;
};

const ConnectedApp = () => (
    <AppStateContextProvider>
        <App />
    </AppStateContextProvider>
);

export default ConnectedApp;
