import React from "react";
import { TouchableOpacity } from "react-native";
import ButtonRegularProps from "./ButtonRegular.types";
import { ButtonContainer, ButtonText } from "./ButtonRegular.styles";

export const ButtonPrimary: React.FC<ButtonRegularProps> = ({
    title,
    onPress,
    disabled,
    testID,
}) => {
    const onPressHandler = () => {
        if (disabled) return;
        if (!onPress) return;
        onPress();
    };
    return (
        <TouchableOpacity testID={testID} onPress={onPressHandler}>
            <ButtonContainer testID="button-container" disabled={disabled}>
                <ButtonText testID="button-title">{title}</ButtonText>
            </ButtonContainer>
        </TouchableOpacity>
    );
};
