/* global jest */

// https://stackoverflow.com/questions/59587799/how-to-resolve-animated-usenativedriver-is-not-supported-because-the-native
jest.mock("react-native/Libraries/Animated/NativeAnimatedHelper");
