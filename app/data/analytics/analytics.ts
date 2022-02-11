import * as Sentry from "@sentry/react-native";

enum AnalyticsEvent {
    // General
    GeneralAppStart = "general_app_start",
    GeneralAppReset = "general_app_reset",
    GeneralFatalError = "general_app_start",

    // Tutorial
    TutorialComplete = "tutorial_complete",

    // Register
    RegisterOpen = "register_open",
    RegisterError = "register_error",
    RegisterComplete = "register_complete",

    // Login
    LoginOpen = "login_open",
    LoginError = "login_error",
    LoginComplete = "login_complete",

    // Password Reset
    PasswordResetOpen = "passwordreset_open",
    PasswordResetStarted = "passwordreset_started",
    PasswordResetStartError = "passwordreset_start_error",
    PasswordResetCompleteError = "passwordreset_complete_error",
    PasswordResetCompleteFinish = "passwordreset_complete_finish",

    // Push notifications
    PushNotificationsRequest = "pushnoticiations_request_permission",
    PushNotificationsSaveToken = "pushnotifications_save_token",

    // Profile
    ProfileChangeProfileStart = "profile_change_profile_start",
    ProfileChangeProfileFinish = "profile_change_profile_finish",
    ProfileChangeProfileError = "profile_change_profile_error",
    ProfileDeleteStart = "profile_delete_start",
    ProfileDeleteError = "profile_delete_error",
    ProfileDeleteComplete = "profile_delete_complete",
}

const logAnalyticsEvent = (type: AnalyticsEvent, data?: object) => {
    // eslint-disable-next-line
    if (process.env.JEST) {
        return;
    }

    // TODO: choose analytics platform and integrate here
    // eslint-disable-next-line no-console
    console.log(`[ANALYTICS]: type=${type} - data=${JSON.stringify(data)}`);

    Sentry.addBreadcrumb({
        type: "analytics",
        message: type,
        data,
    });
};

/* eslint-disable @typescript-eslint/no-explicit-any */

export const logGeneralAppStart = () =>
    logAnalyticsEvent(AnalyticsEvent.GeneralAppStart);
export const logGeneralAppReset = () =>
    logAnalyticsEvent(AnalyticsEvent.GeneralAppReset);
export const logGeneralFatalError = (title: string, description: string) =>
    logAnalyticsEvent(AnalyticsEvent.GeneralFatalError, { title, description });

export const logTutorialComplete = () =>
    logAnalyticsEvent(AnalyticsEvent.TutorialComplete);

export const logRegisterOpen = () =>
    logAnalyticsEvent(AnalyticsEvent.RegisterOpen);
export const logRegisterError = (error: any) =>
    logAnalyticsEvent(AnalyticsEvent.RegisterError, { error });
export const logRegisterComplete = (email: string) =>
    logAnalyticsEvent(AnalyticsEvent.RegisterComplete, { email });

export const logLoginOpen = () => logAnalyticsEvent(AnalyticsEvent.LoginOpen);
export const logLoginError = (error: any) =>
    logAnalyticsEvent(AnalyticsEvent.LoginError, { error });
export const logLoginComplete = (email: string) =>
    logAnalyticsEvent(AnalyticsEvent.LoginComplete, { email });

export const logPasswordResetOpen = () =>
    logAnalyticsEvent(AnalyticsEvent.PasswordResetOpen);
export const logPasswordResetStarted = (email: string) =>
    logAnalyticsEvent(AnalyticsEvent.PasswordResetStarted, { email });
export const logPasswordResetStartError = (error: any) =>
    logAnalyticsEvent(AnalyticsEvent.PasswordResetStartError, { error });
export const logPasswordResetCompleteError = (error: any) =>
    logAnalyticsEvent(AnalyticsEvent.PasswordResetCompleteError, { error });
export const logPasswordResetCompleteFinish = () =>
    logAnalyticsEvent(AnalyticsEvent.PasswordResetCompleteFinish);

export const logPushNotificationsRequest = () =>
    logAnalyticsEvent(AnalyticsEvent.PushNotificationsRequest);
export const logPushNotificationsSaveToken = (token: string) =>
    logAnalyticsEvent(AnalyticsEvent.PushNotificationsSaveToken, { token });

export const logProfileChangeProfileStart = () =>
    logAnalyticsEvent(AnalyticsEvent.ProfileChangeProfileStart);
export const logProfileChangeProfileFinish = () =>
    logAnalyticsEvent(AnalyticsEvent.ProfileChangeProfileFinish);
export const logProfileChangeProfileError = (error: any) =>
    logAnalyticsEvent(AnalyticsEvent.ProfileChangeProfileError, { error });
export const logProfileDeleteStart = () =>
    logAnalyticsEvent(AnalyticsEvent.ProfileDeleteStart);
export const logProfileDeleteError = (error: any) =>
    logAnalyticsEvent(AnalyticsEvent.ProfileDeleteError, { error });
export const logProfileDeleteComplete = () =>
    logAnalyticsEvent(AnalyticsEvent.ProfileDeleteComplete);
