import * as React from "react";
import Router from "router/Router";
import { AppState } from "./router/Router.types";

const App = () => {
    const appState: AppState = {
        hasSeenTutorial: false,
    };

    return <Router appState={appState} />;
};

export default App;
