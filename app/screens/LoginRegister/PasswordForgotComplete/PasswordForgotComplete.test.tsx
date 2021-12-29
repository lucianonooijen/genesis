import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import PasswordForgotComplete from "./PasswordForgotComplete";

describe("PasswordForgotComplete", () => {
    it("should render without throwing", () => {
        render(<PasswordForgotComplete navigation={useMockNavigation()} />);
    });

    it("should only complete password reset after filling in account details", () => {
        const nav = useMockNavigation();
        const r = render(<PasswordForgotComplete navigation={nav} />);

        // Expect button to be disabled and not call navigate
        const saveButton = r.getByA11yLabel("Save my new password");
        expect(saveButton.props.accessibilityState.disabled).toBeTruthy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(saveButton);
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        // Change data

        fireEvent.changeText(r.getByA11yLabel("Reset code"), "reset code");
        expect(saveButton.props.accessibilityState.disabled).toBeTruthy();
        fireEvent.changeText(r.getByA11yLabel("Password"), "newpassword");

        // Expect login to be possible now
        expect(saveButton.props.accessibilityState.disabled).toBeFalsy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(saveButton);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
    });
});
