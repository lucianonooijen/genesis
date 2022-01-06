import styled from "styled-components/native";
import { Label } from "../../Typography/Typography";

export const InputContainer = styled.View`
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    margin: 20px 0;
`;

export const InputElement = styled.TextInput<{ withLabel?: boolean }>`
    padding: ${props => (props.withLabel ? "0 0 6px" : "18px 0 6px")};
    font-size: 18px;
    border-bottom-width: 1px;
    border-color: #e7e7e7;
    width: 100%;
`;

export const LabelText = styled(Label)`
    color: #333;
`;
