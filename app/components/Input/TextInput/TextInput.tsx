import React, { useState } from "react";
import TextInputProps, { InputFieldType } from "./TextInput.types";
import {
    ErrorText,
    InputContainer,
    InputElement,
    LabelText,
} from "./TextInput.styles";

// TODO: Check open source codebases for text input component structure
const TextInput: React.FC<TextInputProps> = ({
    type = InputFieldType.Default,
    label,
    placeholder,
    initialValue,
    onChange,
    testID,
    validatorFunc,
    validationError,
}) => {
    const inputPropsForType = inputPropsForTypeLookup[type];

    const [value, setValue] = useState<string | undefined>(initialValue);
    const changeHandler = (eventValue: string) => {
        if (validatorFunc) {
            validatorFunc("", true);
        }

        setValue(eventValue);
        onChange(eventValue);
    };

    const blurHandler = () => {
        if (validatorFunc) {
            validatorFunc(value || "");
        }
    };

    const hasError = !!validationError;

    const Label = React.useCallback(() => {
        if (!label) return null;
        return (
            <LabelText hasError={hasError} testID="textinput_label">
                {label}
            </LabelText>
        );
    }, [label, hasError]);

    const Error = React.useCallback(() => {
        if (!validationError) return null;
        return (
            <ErrorText testID="textinput_error">{validationError}</ErrorText>
        );
    }, [validationError]);

    return (
        <InputContainer testID={testID}>
            <Label />
            <InputElement
                onChangeText={changeHandler}
                onBlur={blurHandler}
                value={value}
                placeholder={placeholder}
                accessibilityLabel={label}
                hasError={hasError}
                {...inputPropsForType} // eslint-disable-line react/jsx-props-no-spreading
            />
            <Error />
        </InputContainer>
    );
};

const inputPropsForTypeLookup: Record<InputFieldType, object> = {
    [InputFieldType.Default]: {},
    [InputFieldType.TextArea]: {
        numberOfLines: 3,
        multiline: true,
    },
    [InputFieldType.Email]: {
        autoCapitalize: "none",
        autoCompleteType: "email",
        keyboardType: "email-address",
        textContentType: "emailAddress",
    },
    [InputFieldType.Password]: {
        autoCompleteType: "password",
        autoCapitalize: "none",
        secureTextEntry: true,
    },
    [InputFieldType.Phone]: {
        keyboardType: "phone-pad",
    },
};

export default TextInput;
