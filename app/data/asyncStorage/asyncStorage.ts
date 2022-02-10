/* eslint-disable no-console */

import AsyncStorage from "@react-native-community/async-storage";
import { AppStateStorage } from "./asyncStorage.types";

enum AsyncStorageKey {
    Jwt = "jwt",
    HasSeenTutorial = "has_seen_tutorial",
}

enum BooleanString {
    True = "true",
    False = "false",
}

const boolToStorageString = (b: boolean): string => {
    if (b) {
        return BooleanString.True;
    }

    return BooleanString.False;
};

const storageStringToBool = (s: string): boolean => {
    if (s === BooleanString.True) {
        return true;
    }
    if (s === BooleanString.False) {
        return false;
    }

    console.warn(
        `storageStringToBool cannot convert value ${s}, defaulting to true`,
    );

    return false;
};

export const appStateStorage: AppStateStorage = {
    get: async () => {
        const jwt = await AsyncStorage.getItem(AsyncStorageKey.Jwt);
        const hasSeenTutorialString = await AsyncStorage.getItem(
            AsyncStorageKey.HasSeenTutorial,
        );

        const hasSeenTutorial =
            hasSeenTutorialString === null
                ? false
                : storageStringToBool(hasSeenTutorialString);

        const appState = {
            jwt,
            hasSeenTutorial,
        };

        console.log(`appStateStorage get: ${JSON.stringify(appState)}`);

        return appState;
    },
    set: async ({ jwt, hasSeenTutorial }) => {
        if (jwt === null) {
            await AsyncStorage.removeItem(AsyncStorageKey.Jwt);
        } else {
            await AsyncStorage.setItem(AsyncStorageKey.Jwt, jwt);
        }

        await AsyncStorage.setItem(
            AsyncStorageKey.HasSeenTutorial,
            boolToStorageString(hasSeenTutorial),
        );

        console.log("appStateStorage set, printing get for debug:");
        await appStateStorage.get(); // printing for debug
    },
    remove: async () => {
        await AsyncStorage.removeItem(AsyncStorageKey.Jwt);
        await AsyncStorage.removeItem(AsyncStorageKey.HasSeenTutorial);
    },
};
