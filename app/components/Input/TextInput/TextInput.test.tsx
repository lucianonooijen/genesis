import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import TextInput from "./TextInput";
import { InputFieldType } from "./TextInput.types";

describe("TextInput", () => {
    it("should render", () => {
        render(<TextInput onChange={() => {}} />);
    });

    it("should not show the label when no label is given", () => {
        const r = render(<TextInput onChange={() => {}} />);
        expect(() => r.getByTestId("textinput_label")).toThrow();
    });

    it("should show the label when label is given", () => {
        const labelText = "LabelText";
        const r = render(<TextInput label={labelText} onChange={() => {}} />);
        const label = r.getByTestId("textinput_label");
        expect(label).toBeTruthy();
        expect(label.children).toEqual([labelText]);
    });

    it("should call the onChange handler when there is input", async () => {
        const onChange = jest.fn();
        const label = "testinputlabel";
        const r = render(<TextInput label={label} onChange={onChange} />);
        expect(onChange).toHaveBeenCalledTimes(0);
        const testInput = "test input";
        const inputElement = r.getByA11yLabel(label);
        expect(inputElement.props.value).toBeFalsy();
        fireEvent.changeText(inputElement, testInput);
        expect(inputElement.props.value).toBe(testInput);
        expect(onChange).toHaveBeenCalledTimes(1);
        expect(onChange).toHaveBeenCalledWith(testInput);
    });

    it("should render with type set to TextArea", () => {
        render(
            <TextInput onChange={() => {}} type={InputFieldType.TextArea} />,
        );
    });

    it("should render with type set to Email", () => {
        render(<TextInput onChange={() => {}} type={InputFieldType.Email} />);
    });

    it("should render with type set to Password", () => {
        render(
            <TextInput onChange={() => {}} type={InputFieldType.Password} />,
        );
    });

    it("should render with type set to Phone", () => {
        render(<TextInput onChange={() => {}} type={InputFieldType.Phone} />);
    });

    it("should not render the error when there is none", () => {
        const r = render(
            <TextInput onChange={() => {}} type={InputFieldType.Email} />,
        );
        expect(() => r.getByTestId("textinput_error")).toThrow();
    });

    it("should not render the error when it's an emptry string", () => {
        const r = render(
            <TextInput
                onChange={() => {}}
                type={InputFieldType.Email}
                validationError=""
            />,
        );
        expect(() => r.getByTestId("textinput_error")).toThrow();
    });

    it("should render the error when it's truthy", () => {
        const r = render(
            <TextInput
                onChange={() => {}}
                type={InputFieldType.Email}
                validationError="ErrorString"
            />,
        );
        r.getByTestId("textinput_error");
    });
});
