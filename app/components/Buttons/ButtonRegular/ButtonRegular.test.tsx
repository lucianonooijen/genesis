import React from "react";
import {fireEvent, render} from "@testing-library/react-native";
import {ButtonPrimary} from "./ButtonRegular";

describe("ButtonRegular", () => {
    it("should render and show the title", () => {
        const title = "ExampleTitle";
        const r = render(<ButtonPrimary title={title} />);
        const text = r.getByTestId("button-title");
        expect(text).toBeTruthy();
        expect(text.children).toEqual([title]);
    });

    it("should call onPress when pressed and the button is enabled", () => {
        const onPress = jest.fn();
        const r = render(
            <ButtonPrimary title="ExampleTitle" onPress={onPress} />,
        );
        const touchableOpacity = r.getByTestId("button-touchableopacity");
        expect(touchableOpacity).toBeTruthy();
        expect(onPress.mock.calls.length).toBe(0);
        fireEvent.press(touchableOpacity);
        expect(onPress.mock.calls.length).toBe(1);
    });

    it("should not call onPress when pressed and the button is disabled", () => {
        const onPress = jest.fn();
        const r = render(
            <ButtonPrimary title="ExampleTitle" onPress={onPress} disabled />,
        );
        const touchableOpacity = r.getByTestId("button-touchableopacity");
        expect(touchableOpacity).toBeTruthy();
        expect(onPress.mock.calls.length).toBe(0);
        fireEvent.press(touchableOpacity);
        expect(onPress.mock.calls.length).toBe(0);
    });
});
