import React, { useContext } from "react";
import { render } from "@testing-library/react-native";
import { View, Text } from "react-native";
import AppStateContext, {
    AppState,
    AppStateContextProviderTest,
    initialAppState,
} from "./AppState";

describe("AppStateContextProviderTest", () => {
    it("should render the children", () => {
        const Child = () => (
            <View>
                <Text>Example</Text>
            </View>
        );
        const r = render(
            <AppStateContextProviderTest appState={initialAppState}>
                <Child />
            </AppStateContextProviderTest>,
        );
        const child = r.getByText(/Example/);
        expect(child).toBeTruthy();
    });

    it("should pass the appState without modifying anything with custom data", () => {
        const appState: AppState = {
            isLoading: false,
            hasSeenTutorial: true,
            jwt: "tester",

            setIsLoading: jest.fn(),
            setFatalError: jest.fn(),
            setHasSeenTutorial: jest.fn(),
            setJwt: jest.fn(),

            reset: jest.fn(),
        };

        const Child = () => {
            const contextData = useContext(AppStateContext);
            expect(contextData).toEqual(appState);
            return (
                <View>
                    <Text>Example</Text>
                </View>
            );
        };

        render(
            <AppStateContextProviderTest appState={appState}>
                <Child />
            </AppStateContextProviderTest>,
        );
    });
});
