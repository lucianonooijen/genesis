import ApiConfig from "@genesis/api/types/Config";
import AppConfig from "../../types/AppConfig";
import { AppState } from "../AppState/AppState";
import config from "../../config";

export const generateApiConfig = (
    appConfig: AppConfig,
    appState: AppState,
): ApiConfig => ({
    baseUrl: appConfig.baseUrl,
    jwt: appState.jwt,
});

export const getApiConfig = (appState: AppState): ApiConfig =>
    generateApiConfig(config, appState);

export const getApiConfigWithJWT = (jwt: string): ApiConfig => ({
    baseUrl: config.baseUrl,
    jwt,
});
