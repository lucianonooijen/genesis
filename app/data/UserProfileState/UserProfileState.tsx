import React, { useEffect, useState } from "react";
import { UserProfile } from "@genesis/api/types/Profile";

export interface UserProfileState {
    profile: UserProfile;
    hasLoaded: boolean;
    setProfile: (userProfile: UserProfile) => void;
}

const initialUserProfileState: UserProfileState = {
    profile: { firstName: "" },
    hasLoaded: false,
    setProfile: () => {}, // eslint-disable-line @typescript-eslint/no-empty-function
};

const UserProfileStateContext = React.createContext<UserProfileState>(
    initialUserProfileState,
);

const UserProfileStateContextProvider: React.FC = ({ children }) => {
    const [appState, setAppState] = useState<UserProfileState>(
        initialUserProfileState,
    );

    const update = (profile: UserProfile) => {
        setAppState(currentState => ({
            ...currentState,
            profile,
            hasLoaded: true,
        }));
    };

    // Load setters into the state on load
    useEffect(() => {
        setAppState(currentState => ({
            ...currentState,
            setProfile: update,
        }));
    }, []);

    return (
        <UserProfileStateContext.Provider value={appState}>
            {children}
        </UserProfileStateContext.Provider>
    );
};

const UserProfileStateContextProviderTest: React.FC<{
    userProfileState: UserProfileState;
}> = ({ userProfileState, children }) => {
    return (
        <UserProfileStateContext.Provider value={userProfileState}>
            {children}
        </UserProfileStateContext.Provider>
    );
};

export {
    UserProfileStateContextProvider,
    UserProfileStateContextProviderTest,
    initialUserProfileState,
};
export default UserProfileStateContext;
