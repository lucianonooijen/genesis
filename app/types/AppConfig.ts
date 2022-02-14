export default interface AppConfig {
    appVersion: string;
    baseUrl: string;
    environment:
        | "development"
        | "test"
        | "acceptance"
        | "staging"
        | "production";
    sentryDsn: string;
    amplitudeApiKey: string;
}
