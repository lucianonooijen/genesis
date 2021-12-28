export enum InputFieldType {
    Default = "default",
    TextArea = "textarea",
    Email = "email",
    Password = "password",
    Phone = "phone",
}

export default interface TextInputProps {
    type?: InputFieldType;
    label?: string;
    placeholder?: string;
    initialValue?: string;
    onChange: (value: string) => void;
    testID?: string;
}
