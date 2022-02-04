// if you use expo remove this line
import { AppRegistry } from "react-native";
import {
    getStorybookUI,
    configure,
    addDecorator,
} from "@storybook/react-native";
import { withKnobs } from "@storybook/addon-knobs";
import { name as appName } from "../app.json";
import * as stories from "./stories";
import "./rn-addons";

// enables knobs for all stories
addDecorator(withKnobs);

// import stories
configure(() => {
    /* eslint-disable no-unused-expressions */
    stories;
}, module);

const StorybookUI = getStorybookUI({
    port: 7007,
    host: "localhost",
    asyncStorage: null,
});

AppRegistry.registerComponent(appName, () => StorybookUI);

export default StorybookUI;
