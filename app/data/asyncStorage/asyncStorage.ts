import AsyncStorage from "@react-native-community/async-storage";

export interface DataStorage<T> {
    get: () => Promise<T>;
    set: (data: T) => Promise<void>;
}

const jwtKey = "jwt";

export type AppStateStorage = DataStorage<{ jwt?: string | null }>;
export const appStateStorage: AppStateStorage = {
    get: async () => {
        const jwt = await AsyncStorage.getItem(jwtKey);
        return {
            jwt,
        };
    },
    set: async ({ jwt }) => {
        if (!jwt) {
            return;
        }
        AsyncStorage.setItem(jwtKey, jwt);
    },
};
