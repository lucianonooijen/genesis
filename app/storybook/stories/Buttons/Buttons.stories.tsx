import React from "react";
import { Alert, View } from "react-native";
import { storiesOf } from "@storybook/react-native";
import { ButtonPrimary } from "../../../components/Buttons/ButtonRegular/ButtonRegular";

const onPress = () =>
    Alert.alert("Alert Title", "My Alert Msg", [
        {
            text: "Cancel",
            onPress: () => console.log("Cancel Pressed"),
            style: "cancel",
        },
        { text: "OK", onPress: () => console.log("OK Pressed") },
    ]);

storiesOf("Buttons", module)
    .addDecorator(story => <View>{story()}</View>)
    .add("Primary", () => (
        <ButtonPrimary title="Log in" onPress={onPress} disabled={false} />
    ))
    .add("Primary Disabled", () => (
        <ButtonPrimary title="Log in" onPress={onPress} disabled />
    ));
