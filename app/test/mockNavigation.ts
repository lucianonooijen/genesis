import {NativeStackNavigationProp} from "@react-navigation/native-stack/lib/typescript/src/types";
import {ParamListBase} from "@react-navigation/native";

const useMockNavigation: () => NativeStackNavigationProp<ParamListBase> =
    () => {
        return {
            addListener: jest.fn(),
            canGoBack: () => false,
            dispatch: jest.fn(),
            getParent: jest.fn(),
            getState: () => ({
                key: "key",
                index: 0,
                routeNames: [],
                routes: [],
                type: "stack",
                stale: false,
            }),
            goBack: jest.fn(),
            isFocused: () => true,
            pop: jest.fn(),
            popToTop: jest.fn(),
            push: jest.fn(),
            removeListener: jest.fn(),
            replace: jest.fn(),
            reset: jest.fn(),
            setOptions: jest.fn(),
            setParams: jest.fn(),
            navigate: jest.fn(),
        };
    };

export default useMockNavigation;
