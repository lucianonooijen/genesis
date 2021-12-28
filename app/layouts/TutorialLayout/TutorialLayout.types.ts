import { ImageURISource } from "react-native";
import { ScreenTitle } from "../../router/types";

export interface TutorialPageData {
    id: string;
    image: ImageURISource;
    title: string;
    text: string;
    nextScreen: ScreenTitle;
}

export interface TutorialLayoutProps {
    pageData: TutorialPageData;
    navigateFunction: (screen: ScreenTitle) => void;
}
