module.exports = {
    options: {
        Component: {
            path: "src/components",
            template: "reactNativeNoRedux",
            additionalInstructions:
                "Note: components are not connected to Redux by default. If you need a Redux connection, use a container or section",
        },
        Container: { path: "src/containers", template: "reactNativeWithRedux" },
        Section: { path: "src/sections", template: "reactNativeWithRedux" },
        Page: {
            path: "src/screens",
            template: "reactNativeNoRedux",
            additionalInstructions:
                `Notes:
    * Pages are not connected to Redux by default. If you need a Redux connection, use a container or section.
    * The template does not yet wrap the components in the layout, if needed, add this by hand.
    * Make sure to add the page to the routes config file.`,
        },
        ReduxDomain: {
            path: "src/store",
            template: "reduxDomain",
            additionalInstructions:
                "Note: make sure to add the new domain to the rootReducer",
        },
    },
};
