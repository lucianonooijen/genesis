import styled from "styled-components/native";
import { ButtonPrimary } from "../../components/Buttons/ButtonRegular/ButtonRegular";

export const ImageHeader = styled.Image`
    height: 50%;
`;

export const TutorialTextContainer = styled.View`
    margin: 40px 20px 20px;
`;

export const NextButton = styled(ButtonPrimary)`
    position: absolute;
    left: 20px;
    bottom: 20px;
    right: 20px;
`;
