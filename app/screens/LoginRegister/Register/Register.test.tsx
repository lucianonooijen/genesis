import React from "react";
import { fireEvent, render, waitFor } from "@testing-library/react-native";
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

    it("should only allow register after filling in account details", async () => {
        const nav = useMockNavigation();
        const appState = initialAppState;
        appState.setJwt = jest.fn();
        const testJwt = "testing_jwt";
        const apiCall = jest.fn().mockResolvedValue({ jwt: testJwt });

        const r = render(
            <AppStateContextProviderTest appState={appState}>
                <Register navigation={nav} registerApiCall={apiCall} />
            </AppStateContextProviderTest>,
        );

        // Expect button to be disabled and not call navigate
        const registerButton = r.getByA11yLabel("Register");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        expect(appState.setJwt).toHaveBeenCalledTimes(0);
        fireEvent.press(registerButton);
        expect(appState.setJwt).toHaveBeenCalledTimes(0);

        // Change data
        fireEvent.changeText(r.getByA11yLabel("Name"), "example");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Email"), "email@example.com");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Password"), "verysecret");

        // Expect login to be possible now
        expect(registerButton.props.accessibilityState.disabled).toBeFalsy();
        expect(appState.setJwt).toHaveBeenCalledTimes(0);
        fireEvent.press(registerButton);
        await waitFor(() => expect(appState.setJwt).toHaveBeenCalledTimes(1));
        expect(nav.navigate).toHaveBeenCalledTimes(0);
    });
});
