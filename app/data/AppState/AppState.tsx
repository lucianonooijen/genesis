import { appStateStorage } from "data/asyncStorage/asyncStorage";
import React, { useEffect, useState } from "react";

export interface AppState {
    isLoading: boolean;
    fatalError?: { title: string; description: string };
    hasSeenTutorial: boolean;
    jwt: string | null;

    setIsLoading: (loading: boolean) => void;
    setFatalError: (title: string, description: string) => void;
    setHasSeenTutorial: (hasSeenTutorial: boolean) => void;
    setJwt: (jwt: string | null) => void;

    reset: () => void;
}

const initialAppState: AppState = {
    isLoading: true,
    fatalError: undefined,
    hasSeenTutorial: false,
    jwt: null,

    setIsLoading: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
    setFatalError: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
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

        appStateStorage.set({ hasSeenTutorial, jwt: appState.jwt });
    };

    const setFatalError = (title: string, description: string) => {
        setAppState(currentState => ({
            ...currentState,
            fatalError: { title, description },
        }));
    };

    const setJwt = (jwt: string | null) => {
        setAppState(currentState => ({
            ...currentState,
            jwt,
        }));

        appStateStorage.set({ jwt, hasSeenTutorial: appState.hasSeenTutorial });
    };

    const reset = () => {
        setAppState(currentState => ({
            ...currentState,
            jwt: initialAppState.jwt,
            hasSeenTutorial: initialAppState.hasSeenTutorial,
        }));

        appStateStorage.remove();
    };

    useEffect(() => {
        setIsLoading(true);

        // Load setters into the state on load
        setAppState(currentState => ({
            ...currentState,
            setIsLoading,
            setHasSeenTutorial,
            setFatalError,
            setJwt,
            reset,
        }));

        // Load state from memory on load
        appStateStorage
            .get()
            .then(diskState => {
                setAppState(currentState => ({
                    ...currentState,
                    jwt: diskState.jwt,
                    hasSeenTutorial: diskState.hasSeenTutorial,
                    isLoading: false, // instead of `setIsLoading(false)`
                }));
            })
            .catch(err => {
                console.error(err); // eslint-disable-line no-console

                setIsLoading(false);
            });
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <AppStateContext.Provider value={appState}>
            {children}
        </AppStateContext.Provider>
    );
};

const AppStateContextProviderTest: React.FC<{ appState?: AppState }> = ({
    appState = initialAppState,
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
