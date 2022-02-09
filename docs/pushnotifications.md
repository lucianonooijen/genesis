# Push notifications

To set up push notifications, a few manual steps are required to set up the Apple and Google services correctly.

## Apple instructions:

Related server variables:

* `APNS_TOPIC`: app package name, i.e. `nl.bytecode.genesis`
* `APNS_TEAM_ID`: see step 4
* `APNS_KEY_ID`: see step 3
* `APNS_KEY_BASE64`: see step 3

Steps:

1. Create app in App Store Connect with the correct app package name (`APNS_TOPIC` value, i.e. `nl.bytecode.genesis`)
2. Create key in https://developer.apple.com/account/resources/authkeys/add
    - Register a New Key
    - APNs
3. Download key, save the Key and Key ID.
   - Run `cat [keyname].p8 | base64` to get the `APNS_KEY_BASE64` value
   - The Key ID is the `APNS_KEY_ID` value
4. Find Team ID
    - Open a web browser to https://developer.apple.com/account. Login with your Apple Developer account credentials.
    - When you reach the Overview page, click the "Membership" item on the left menu.
    - On the Membership page, locate the "Team ID" field. This is the `APNS_TEAM_ID` value.

## Google instructions:

Related server variables:

* `FCM_CHANNEL_ID`: app package name, i.e. `nl.bytecode.genesis`
* `FCM_CREDENTIALS_BASE64`: see step 7

Steps:

1. Create a Google Developer Console project
2. Create a Firebase Console project linked to the Google Developer Console project
3. In Firebase, go to project settings > Cloud Messaging. Under project credentials is a server key.
4. Create a google-services.json in Firebase by going to:
   - Project settings (cog wheel in top left next to Project Overview)
   - There, create an Android app (scroll down to your apps, select Android icon)
   - Enter the matching package name of the app (`FCM_CHANNEL_ID` value, i.e. `nl.bytecode.genesis`)
   - Click on "download google-services.json"
5. Insert the google-services.json into the repo at `/app/android/app`
6. Go back to the Firebase dashboard
   - Click on the cog right of "Project overview"
   - Open "Project settings", subpage "Service accounts"
   - Open "Firebase Admin SDK" tab
   - Press "Generate new private key"
   - Download the key to your disk
7. Run `cat [keyname].json | base64` to get the `FCM_CREDENTIALS_BASE64` value
