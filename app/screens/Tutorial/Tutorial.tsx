import React from "react";
import {
    StackNavigationProps,
    WithStackNavigationProps,
} from "../../types/Navigation";
import {TutorialProps} from "./Tutorial.types";
import TutorialLayout from "../../layouts/TutorialLayout/TutorialLayout";
import {TutorialPageData} from "../../layouts/TutorialLayout/TutorialLayout.types";
import tutorialData from "./data";

const Tutorial: React.FC<WithStackNavigationProps<TutorialProps>> = ({
    pageData,
    navigation,
}) => {
    return (
        <TutorialLayout
            pageData={pageData}
            navigateFunction={navigation.navigate}
        />
    );
};

const generateTutorialPage =
    (pageData: TutorialPageData): React.FC<StackNavigationProps> =>
    ({navigation}) =>
        <Tutorial pageData={pageData} navigation={navigation} />;

export const TutorialScreenOne = generateTutorialPage(tutorialData[0]);
export const TutorialScreenTwo = generateTutorialPage(tutorialData[1]);
export const TutorialScreenThree = generateTutorialPage(tutorialData[2]);
