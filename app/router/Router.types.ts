export interface AppState {
    hasSeenTutorial: boolean;
    isLoggedIn: boolean;
}
export interface RouterProps {
    appState: AppState;
}
