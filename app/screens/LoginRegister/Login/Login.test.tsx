import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import Login from "./Login";
import {
    AppStateContextProviderTest,
    initialAppState,
} from "../../../data/AppState/AppState";
import { LoginRegisterScreens } from "../../../router/types";

describe("Login", () => {
    it("should render without throwing", () => {
        render(<Login navigation={useMockNavigation()} />);
    });

    it("should only allow log in after filling in account details", () => {
        const nav = useMockNavigation();
        const appState = initialAppState;
        appState.setIsLoggedIn = jest.fn();
        const r = render(
            <AppStateContextProviderTest appState={appState}>
                <Login navigation={nav} />
            </AppStateContextProviderTest>,
        );

        // Expect button to be disabled and not call navigate
        const loginButton = r.getByA11yLabel("Log in");
        expect(loginButton.props.accessibilityState.disabled).toBeTruthy();
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(0);
        fireEvent.press(loginButton);
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(0);

        // Change email and password
        fireEvent.changeText(r.getByA11yLabel("Email"), "email@example.com");
        expect(loginButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Password"), "verysecret");

        // Expect login to be possible now
        expect(loginButton.props.accessibilityState.disabled).toBeFalsy();
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(0);
        fireEvent.press(loginButton);
        expect(appState.setIsLoggedIn).toHaveBeenCalledTimes(1);
        expect(appState.setIsLoggedIn).toHaveBeenCalledWith(true);
        expect(nav.navigate).toHaveBeenCalledTimes(0);
    });

    it("should navigate to password reset when corresponding button is pressed", () => {
        const nav = useMockNavigation();
        const r = render(<Login navigation={nav} />);

        const button = r.getByA11yLabel("Forgot password");
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(button);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
        expect(nav.navigate).toHaveBeenCalledWith(
            LoginRegisterScreens.PasswordForgotStart,
        );
    });

    it("should navigate to register when corresponding button is pressed", () => {
        const nav = useMockNavigation();
        const r = render(<Login navigation={nav} />);

        const button = r.getByA11yLabel("Create an account");
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(button);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
        expect(nav.navigate).toHaveBeenCalledWith(
            LoginRegisterScreens.Register,
        );
    });
});
