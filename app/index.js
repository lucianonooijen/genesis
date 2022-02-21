import { AppRegistry } from "react-native";
import * as Sentry from "@sentry/react-native";
import App from "./App";
import { name as appName } from "./app.json";
import config from "./config";
// import Storybook from "./storybook";

Sentry.init({
    dsn: config.sentryDsn,
    environment: config.environment,

    // Set tracesSampleRate to 1.0 to capture 100% of transactions for performance monitoring.
    // We recommend adjusting this value in production.
    tracesSampleRate: 1.0,
});

const SentrifiedApp = Sentry.wrap(App);

AppRegistry.registerComponent(appName, () => SentrifiedApp);
// AppRegistry.registerComponent(appName, () => Storybook);
