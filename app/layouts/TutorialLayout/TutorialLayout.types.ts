import { ImageURISource } from "react-native";

export interface TutorialPageData {
    id: string;
    image: ImageURISource;
    title: string;
    text: string;
    onPressContinue: () => void;
}

export interface TutorialLayoutProps {
    pageData: TutorialPageData;
}
