import React from "react";
import { fireEvent, render, waitFor } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import PasswordForgotComplete from "./PasswordForgotComplete";
import {
    AppStateContextProviderTest,
    initialAppState,
} from "../../../data/AppState/AppState";

describe("PasswordForgotComplete", () => {
    it("should render without throwing", () => {
        render(<PasswordForgotComplete navigation={useMockNavigation()} />);
    });

    it("should only complete password reset after filling in account details", async () => {
        const nav = useMockNavigation();
        const appState = initialAppState;
        appState.setJwt = jest.fn();
        const testJwt = "testing_jwt";
        const apiCall = jest.fn().mockResolvedValue({ jwt: testJwt });

        const r = render(
            <AppStateContextProviderTest appState={appState}>
                <PasswordForgotComplete navigation={nav} apiCall={apiCall} />
            </AppStateContextProviderTest>,
        );

        // Expect button to be disabled and not call navigate
        const saveButton = r.getByA11yLabel("Save my new password");
        expect(saveButton.props.accessibilityState.disabled).toBeTruthy();
        expect(appState.setJwt).toHaveBeenCalledTimes(0);
        fireEvent.press(saveButton);
        expect(appState.setJwt).toHaveBeenCalledTimes(0);
        // Change data

        fireEvent.changeText(r.getByA11yLabel("Reset code"), "reset code");
        expect(saveButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Password"), "newpassword");

        // Expect login to be possible now
        expect(saveButton.props.accessibilityState.disabled).toBeFalsy();
        expect(appState.setJwt).toHaveBeenCalledTimes(0);
        fireEvent.press(saveButton);
        await waitFor(() => expect(appState.setJwt).toHaveBeenCalledTimes(1));
        expect(appState.setJwt).toHaveBeenCalledWith(testJwt);
        expect(nav.navigate).toHaveBeenCalledTimes(0);
    });
});
