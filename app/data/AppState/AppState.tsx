import React, { useEffect, useState } from "react";

export interface AppState {
    isLoading: boolean;
    hasSeenTutorial: boolean;
    isLoggedIn: boolean;

    setIsLoading: (loading: boolean) => void;
    setHasSeenTutorial: (hasSeenTutorial: boolean) => void;
    setIsLoggedIn: (isLoggedIn: boolean) => void;
}

const initialAppState: AppState = {
    isLoading: true,
    hasSeenTutorial: false,
    isLoggedIn: false,

    setIsLoading: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
    setHasSeenTutorial: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
    setIsLoggedIn: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
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

    const setIsLoggedIn = (isLoggedIn: boolean) => {
        setAppState(currentState => ({
            ...currentState,
            isLoggedIn,
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
            setIsLoggedIn,
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
