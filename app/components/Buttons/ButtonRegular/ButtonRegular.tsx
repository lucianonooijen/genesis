import React from "react";
import {TouchableOpacity} from "react-native";
import ButtonRegularProps from "./ButtonRegular.types";
import {ButtonContainer, ButtonText} from "./ButtonRegular.styles";

export const ButtonPrimary: React.FC<ButtonRegularProps> = ({
    title,
    onPress,
    disabled,
}) => {
    const onPressHandler = () => {
        if (disabled) return;
        if (!onPress) return;
        onPress();
    };
    return (
        <TouchableOpacity onPress={onPressHandler}>
            <ButtonContainer disabled={disabled}>
                <ButtonText>{title}</ButtonText>
            </ButtonContainer>
        </TouchableOpacity>
    );
};
