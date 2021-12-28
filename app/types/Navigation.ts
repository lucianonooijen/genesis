import { ParamListBase } from "@react-navigation/native";
import { NativeStackNavigationProp } from "@react-navigation/native-stack/lib/typescript/src/types";

export interface StackNavigationProps {
    navigation: NativeStackNavigationProp<ParamListBase>;
}

export type WithStackNavigationProps<T> = T & StackNavigationProps;
