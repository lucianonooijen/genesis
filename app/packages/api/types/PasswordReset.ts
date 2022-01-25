export interface PasswordResetStartRequest {
    email: string;
}

export interface PasswordResetCompleteRequest {
    resetToken: string;
    password: string;
}
