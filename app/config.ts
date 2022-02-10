import packageJson from "./package.json";
import AppConfig from "./types/AppConfig";

const config: AppConfig = {
    appVersion: packageJson.version,
    baseUrl: "http://localhost:5000/v1",
};

export default config;
