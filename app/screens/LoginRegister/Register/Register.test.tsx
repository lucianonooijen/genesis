import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import Register from "./Register";

describe("Register", () => {
    it("should render without throwing", () => {
        render(<Register navigation={useMockNavigation()} />);
    });

    it("should only allow register after filling in account details", () => {
        const nav = useMockNavigation();
        const r = render(<Register navigation={nav} />);

        // Expect button to be disabled and not call navigate
        const registerButton = r.getByA11yLabel("Register");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(registerButton);
        expect(nav.navigate).toHaveBeenCalledTimes(0);

        // Change data
        fireEvent.changeText(r.getByA11yLabel("Name"), "example");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Email"), "email@example.com");
        expect(registerButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Password"), "verysecret");

        // Expect login to be possible now
        expect(registerButton.props.accessibilityState.disabled).toBeFalsy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(registerButton);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
    });
});
