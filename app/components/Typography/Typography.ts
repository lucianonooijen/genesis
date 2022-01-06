import styled from "styled-components/native";

// TODO: Move this to styling/theme package
const scalingFactor = 1.2;
const baseSizePx = 18;
const getFontSize = (scalingExponent: number): string => {
    const scalingNumber = scalingFactor ** scalingExponent;
    const fontSize = baseSizePx * scalingNumber;
    const fontSizeRounded = Math.round(fontSize);
    return `${fontSizeRounded}px`;
};

export const Title = styled.Text`
    font-size: ${getFontSize(4)};
`;

export const SubTitle = styled.Text`
    font-size: ${getFontSize(2)};
`;

export const Paragraph = styled.Text`
    font-size: ${getFontSize(1)};
`;

export const Label = styled.Text`
    font-size: ${getFontSize(-1)};
`;
