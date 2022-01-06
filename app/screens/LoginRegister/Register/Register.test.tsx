import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import Register from "./Register";
import {
    AppStateContextProviderTest,
    initialAppState,
} from "../../../data/AppState/AppState";

describe("Register", () => {
    it("should render without throwing", () => {
        render(<Register navigation={useMockNavigation()} />);
    });

    it("should only allow register after filling in account details", () => {
        const nav = useMockNavigation();
        const appState = initialAppState;
        appState.setIsLoggedIn = jest.fn();

        const r = render(
            <AppStateContextProviderTest appState={appState}>
                <Register navigation={nav} />
            </AppStateContextProviderTest>,
        );

        // Expect button to be disabled and not call navigate
        const registerButton = r.getByA11yLabel("Register");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(0);
        fireEvent.press(registerButton);
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(0);

        // Change data
        fireEvent.changeText(r.getByA11yLabel("Name"), "example");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Email"), "email@example.com");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Password"), "verysecret");

        // Expect login to be possible now
        expect(registerButton.props.accessibilityState.disabled).toBeFalsy();
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(0);
        fireEvent.press(registerButton);
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(1);
        expect(appState.setIsLoggedIn).toHaveBeenCalledWith(true);
        expect(nav.navigate).toHaveBeenCalledTimes(0);
    });
});
