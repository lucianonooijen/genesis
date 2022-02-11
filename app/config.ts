import packageJson from "./package.json";
import AppConfig from "./types/AppConfig";

const config: AppConfig = {
    appVersion: packageJson.version,
    baseUrl: "http://localhost:5000/v1",
    environment: "development",
    sentryDsn:
        "https://7fc62c48b6ae444db483df7f6b64c47d@o345560.ingest.sentry.io/6199409",
};

export default config;
