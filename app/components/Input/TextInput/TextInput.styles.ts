import styled from "styled-components/native";
import { Label } from "../../Typography/Typography";

export const InputContainer = styled.View`
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    margin: 20px 0;
`;

// TODO: load from theme instead of local variable
const errorColor = "#f00";

export const InputElement = styled.TextInput<{ hasError: boolean }>`
    padding: ${props => (props.hasError ? "18px 0 0;" : "18px 0 10px;")}
    font-size: 18px;
    border-bottom-width: 1px;
    border-color: ${props => (props.hasError ? errorColor : "#e7e7e7")};
    width: 100%;
`;

export const LabelText = styled(Label)<{ hasError: boolean }>`
    color: ${props => (props.hasError ? errorColor : "#333")};
`;

export const ErrorText = styled(Label)`
    color: ${errorColor};
    font-size: 8px;
    padding: 0;
    margin: 0;
    height: 10px;
`;
