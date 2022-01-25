import React, { useEffect, useState } from "react";

export interface AppState {
    isLoading: boolean;
    hasSeenTutorial: boolean;
    jwt: string | undefined;

    setIsLoading: (loading: boolean) => void;
    setHasSeenTutorial: (hasSeenTutorial: boolean) => void;
    setJwt: (isLoggedIn: string | undefined) => void;

    reset: () => void;
}

const initialAppState: AppState = {
    isLoading: true,
    hasSeenTutorial: false,
    jwt: undefined,

    setIsLoading: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
    setHasSeenTutorial: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
    setJwt: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function

    reset: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
};

const AppStateContext = React.createContext<AppState>(initialAppState);

const AppStateContextProvider: React.FC = ({ children }) => {
    const [appState, setAppState] = useState<AppState>(initialAppState);

    const setIsLoading = (isLoading: boolean) => {
        setAppState(currentState => ({
            ...currentState,
            isLoading,
        }));
    };

    const setHasSeenTutorial = (hasSeenTutorial: boolean) => {
        setAppState(currentState => ({
            ...currentState,
            hasSeenTutorial,
        }));
    };

    const setJwt = (jwt: string | undefined) => {
        setAppState(currentState => ({
            ...currentState,
            jwt,
        }));
    };

    const reset = () => {
        setAppState(currentState => ({
            ...currentState,
            jwt: initialAppState.jwt,
            hasSeenTutorial: initialAppState.hasSeenTutorial,
        }));
    };

    // Load state from memory on load
    useEffect(() => {
        setIsLoading(true);
        // TODO: Load this from AsyncStorage
        setIsLoading(false);
    }, []);

    // Load setters into the state on load
    useEffect(() => {
        setAppState(currentState => ({
            ...currentState,
            setIsLoading,
            setHasSeenTutorial,
            setJwt,
            reset,
        }));
    }, []);

    return (
        <AppStateContext.Provider value={appState}>
            {children}
        </AppStateContext.Provider>
    );
};

const AppStateContextProviderTest: React.FC<{ appState: AppState }> = ({
    appState,
    children,
}) => {
    return (
        <AppStateContext.Provider value={appState}>
            {children}
        </AppStateContext.Provider>
    );
};

export {
    AppStateContextProvider,
    AppStateContextProviderTest,
    initialAppState,
};
export default AppStateContext;
