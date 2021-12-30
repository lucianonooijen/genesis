import React from "react";
import { View } from "react-native";
import { Paragraph, Title } from "components/Typography/Typography";
import { TutorialLayoutProps } from "./TutorialLayout.types";
import {
    ImageHeader,
    NextButton,
    TutorialTextContainer,
} from "./TutorialLayout.styles";

const TutorialLayout: React.FC<TutorialLayoutProps> = ({ pageData }) => {
    const { id, image, title, text, onPressContinue } = pageData;

    return (
        <View key={id} testID={id}>
            <ImageHeader source={image} />
            <TutorialTextContainer>
                <Title testID="tutorial-title">{title}</Title>
                <Paragraph testID="tutorial-text">{text}</Paragraph>
            </TutorialTextContainer>
            <NextButton
                title="Continue"
                onPress={onPressContinue}
                testID="tutorial-nextbutton"
            />
        </View>
    );
};

export default TutorialLayout;
