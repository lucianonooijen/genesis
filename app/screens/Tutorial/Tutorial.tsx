import React, { useContext } from "react";
import { StackNavigationProps } from "../../types/Navigation";
import TutorialLayout from "../../layouts/TutorialLayout/TutorialLayout";
import tutorialData from "./data";
import { TutorialScreens } from "../../router/types";
import appState from "../../data/AppState/AppState";

export const TutorialScreenOne: React.FC<StackNavigationProps> = ({
    navigation,
}) => {
    const pageDataBase = tutorialData[0];
    const onPressContinue = () =>
        navigation.navigate(TutorialScreens.ScreenTwo);
    const pageData = { ...pageDataBase, onPressContinue };

    return <TutorialLayout pageData={pageData} />;
};

export const TutorialScreenTwo: React.FC<StackNavigationProps> = ({
    navigation,
}) => {
    const pageDataBase = tutorialData[1];
    const onPressContinue = () =>
        navigation.navigate(TutorialScreens.ScreenThree);
    const pageData = { ...pageDataBase, onPressContinue };

    return <TutorialLayout pageData={pageData} />;
};

export const TutorialScreenThree: React.FC<StackNavigationProps> = () => {
    const { setHasSeenTutorial } = useContext(appState);
    const pageDataBase = tutorialData[2];
    const onPressContinue = () => setHasSeenTutorial(true);
    const pageData = { ...pageDataBase, onPressContinue };

    return <TutorialLayout pageData={pageData} />;
};
