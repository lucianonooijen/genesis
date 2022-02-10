import PushNotification from "react-native-push-notification";
import PushNotificationIOS from "@react-native-community/push-notification-ios";
import { appStateStorage } from "data/asyncStorage/asyncStorage";
import { pushNotificationRegisterToken } from "@genesis/api";
import { getApiConfigWithJWT } from "data/api/api";

const tokenOsToServerOs = (tokenOs: string): string => {
    if (tokenOs === "ios") {
        return "iOS";
    }
    if (tokenOs === "android") {
        return "Android";
    }
    return tokenOs;
};

export const configurePushNotifications = () => {
    PushNotification.configure({
        async onRegister(tokenObject) {
            appStateStorage.get().then(appState => {
                if (appState.jwt) {
                    const apiConfig = getApiConfigWithJWT(appState.jwt);
                    pushNotificationRegisterToken(apiConfig, {
                        platform: tokenOsToServerOs(tokenObject.os),
                        token: tokenObject.token,
                    });
                }
            });
        },
        onNotification(notification) {
            console.warn("Received notification: ", notification); // eslint-disable-line no-console
            notification.finish(PushNotificationIOS.FetchResult.NoData);
        },
        permissions: {
            alert: true,
            badge: true,
            sound: true,
        },
        popInitialNotification: true,
        requestPermissions: false, // Note: need to do this in the app through requestPushPermissions() below;
    });
};

// Note: only sends out an API request if JWT is set in local storage (see onRegister in configurePushNotifications).
export const requestPushPermissions = () => {
    PushNotification.requestPermissions();
};
