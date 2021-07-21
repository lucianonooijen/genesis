import React from "react";
import { Platform, StatusBar, Text, View } from "react-native";

export default function App() {
  const StatusBarForPlatform = () => {
    if (Platform.OS === "ios") {
      return <StatusBar barStyle="dark-content" hidden={false} translucent />;
    }
    return null;
  };

  return (
    <View>
      <Text>Welcome to Genesis</Text>
      <StatusBarForPlatform />
    </View>
  );
}
