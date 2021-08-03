import React from 'react';
import {
  SafeAreaView, StatusBar, useColorScheme,
} from 'react-native';

const App = () => {

  const AppStatusBar = () => {
    const isDarkMode = useColorScheme() === 'dark';
    return (
      <StatusBar
        barStyle={isDarkMode ? 'light-content' : 'dark-content'}
      />)
  }

  return (
    <SafeAreaView>
      <AppStatusBar />
    </SafeAreaView>
  );
};

export default App;
