import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import Login from "./Login";

describe("Login", () => {
    it("should render without throwing", () => {
        render(<Login navigation={useMockNavigation()} />);
    });

    it("should only allow log in after filling in account details", () => {
        const nav = useMockNavigation();
        const r = render(<Login navigation={nav} />);

        // Expect button to be disabled and not call navigate
        const loginButton = r.getByA11yLabel("Log in");
        expect(loginButton.props.accessibilityState.disabled).toBeTruthy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(loginButton);
        expect(nav.navigate).toHaveBeenCalledTimes(0);

        // Change email and password
        fireEvent.changeText(r.getByA11yLabel("Email"), "email@example.com");
        expect(loginButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Password"), "verysecret");

        // Expect login to be possible now
        expect(loginButton.props.accessibilityState.disabled).toBeFalsy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(loginButton);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
    });
});
