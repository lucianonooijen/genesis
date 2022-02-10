export enum GetEndpoints {
    VersionCheck = "app-version",
    UserProfile = "user/profile",
}

export enum PostEndpoints {
    CreateAccount = "user/register",
    Login = "user/login",
    PassResetStart = "user/password-reset/start",
    PassResetComplete = "user/password-reset/complete",
    PushNotifications = "user/push-notifications",
}

export enum PutEndpoints {
    UserProfile = "user/profile",
}

export enum DeleteEndpoints {
    UserProfile = "user/profile",
}
