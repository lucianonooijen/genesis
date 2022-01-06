import React, { useState } from "react";
import TextInputProps, { InputFieldType } from "./TextInput.types";
import { InputContainer, InputElement, LabelText } from "./TextInput.styles";

// TODO: Check open source codebases for text input component structure
const TextInput: React.FC<TextInputProps> = ({
    type = InputFieldType.Default,
    label,
    placeholder,
    initialValue,
    onChange,
    testID,
}) => {
    const inputPropsForType = inputPropsForTypeLookup[type];

    const [value, setValue] = useState<string | undefined>(initialValue);
    const changeHandler = (eventValue: string) => {
        setValue(eventValue);
        onChange(eventValue);
    };

    const Label = React.useCallback(() => {
        if (!label) return null;
        return <LabelText testID="textinput_label">{label}</LabelText>;
    }, [label]);

    return (
        <InputContainer testID={testID}>
            <Label />
            <InputElement
                onChangeText={changeHandler}
                value={value}
                placeholder={placeholder}
                accessibilityLabel={label}
                {...inputPropsForType} // eslint-disable-line react/jsx-props-no-spreading
            />
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
