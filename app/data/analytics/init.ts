import { Amplitude } from "@amplitude/react-native";
import config from "../../config";
import { appStateStorage } from "../asyncStorage/asyncStorage";
import { parseJwt } from "./base64";

export const initAnalytics = () => {
    // eslint-disable-next-line
    if (process.env.JEST) {
        return;
    }

    initAmplitude();
};

/* eslint-disable no-console */

const initAmplitude = async () => {
    const ampInstance = Amplitude.getInstance();
    ampInstance
        .init(config.amplitudeApiKey)
        .then(r => console.log(`initialized Amplitude: ${r}`))
        .catch(e => console.error(`error initializing Amplitude ${e}`));

    loadAmplitudeUserDataFromJWT();
};

export const loadAmplitudeUserDataFromJWT = async () => {
    const ampInstance = Amplitude.getInstance();

    try {
        const appState = await appStateStorage.get();

        if (appState.jwt) {
            const jwt = parseJwt(appState.jwt);

            ampInstance
                .setUserId(jwt.aud)
                .then(() =>
                    console.log(`[ANALYTICS]: userid set to ${jwt.aud}`),
                )
                .catch(console.error);
        }
    } catch (e) {
        console.warn(`error in loadAmplitudeUserFromJWT: ${e}`);
    }
};
