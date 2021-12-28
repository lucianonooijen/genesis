import styled from "styled-components/native";

export const ButtonContainer = styled.View<{disabled?: boolean}>`
    margin: 0;
    padding: 5px 10px;
    text-align: center;
    background-color: ${props => (props.disabled ? "#bbb" : "#222")}
    border-radius: 24px;
    height: 64px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

`;

export const ButtonText = styled.Text`
    font-style: normal;
    font-size: 20px;
    line-height: 24px;
    color: #fff;
`;
