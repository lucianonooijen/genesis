export enum MainScreens {
    Home = "Main_Home",
}

export enum TutorialScreens {
    ScreenOne = "Tutorial_ScreenOne",
    ScreenTwo = "Tutorial_ScreenTwo",
    ScreenThree = "Tutorial_ScreenThree",
}

export enum LoginRegisterScreens {
    LoginRegisterLanding = "LoginRegister_Landing",
    Login = "LoginRegister_Login",
    Register = "LoginRegister_Register",
    PasswordForgotStart = "LoginRegister_PasswordForgotStart",
    PasswordForgotComplete = "LoginRegister_PasswordForgotComplete",
}

export type ScreenTitle = MainScreens | TutorialScreens | LoginRegisterScreens;
