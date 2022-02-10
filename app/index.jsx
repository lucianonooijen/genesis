import React from "react";
import { AppRegistry } from "react-native";
import App from "./App";
import { name as appName } from "./app.json";
import StorybookUIRoot from "./storybook";

const AppSwitch = () => {
    const [storybookActive, setStorybookActive] = React.useState(false);
    const toggleStorybook = React.useCallback(
        () => setStorybookActive(active => !active),
        [],
    );

    React.useEffect(() => {
        /* eslint-siable-next-line no-undef */
        if (__DEV__) {
            /* eslint-disable-next-line global-require, @typescript-eslint/no-var-requires  */
            const DevMenu = require("react-native-dev-menu");
            DevMenu.addItem("Toggle Storybook", toggleStorybook);
        }
    }, [toggleStorybook]);

    return storybookActive ? <StorybookUIRoot /> : <App />;
};

AppRegistry.registerComponent(appName, () => AppSwitch);
