import React, { useEffect, useState, useContext } from "react";
import Router from "router/Router";
import StorybookUIRoot from "storybook";
import AppStateContext, {
    AppStateContextProvider,
} from "./data/AppState/AppState";
import { UserProfileStateContextProvider } from "./data/UserProfileState/UserProfileState";

const App = () => {
    const appState = useContext(AppStateContext);

    const [storybookActive, setStorybookActive] = useState(false);
    const toggleStorybook = React.useCallback(
        () => setStorybookActive(active => !active),
        [],
    );

    useEffect(() => {
        if (__DEV__) {
            /* eslint-disable-next-line global-require, @typescript-eslint/no-var-requires  */
            const DevMenu = require("react-native-dev-menu");
            DevMenu.addItem("Toggle Storybook", toggleStorybook);
        }
    }, [toggleStorybook]);

    return storybookActive ? (
        <StorybookUIRoot />
    ) : (
        <Router appState={appState} />
    );
};

const ConnectedApp = () => (
    <AppStateContextProvider>
        <UserProfileStateContextProvider>
            <App />
        </UserProfileStateContextProvider>
    </AppStateContextProvider>
);

export default ConnectedApp;
