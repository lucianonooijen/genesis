import React from "react";
import {View} from "react-native";
import {Paragraph, Title} from "components/Typography/Typography";
import {TutorialLayoutProps, TutorialPageData} from "./TutorialLayout.types";
import {
    ImageHeader,
    NextButton,
    TutorialTextContainer,
} from "./TutorialLayout.styles";

const TutorialLayout: React.FC<TutorialLayoutProps> = ({
    pageData,
    navigateFunction,
}) => {
    const {id, image, title, text, nextScreen} = pageData;
    const goToNextPage = () => navigateFunction(nextScreen);

    return (
        <View key={id} testID={id}>
            <ImageHeader source={image} />
            <TutorialTextContainer>
                <Title testID="tutorial-title">{title}</Title>
                <Paragraph testID="tutorial-text">{text}</Paragraph>
            </TutorialTextContainer>
            <NextButton
                title="Continue"
                onPress={goToNextPage}
                testID="tutorial-nextbutton"
            />
        </View>
    );
};

export default TutorialLayout;
