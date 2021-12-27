import React from "react"
import {render} from "@testing-library/react-native";
import Router from "./index"

describe("router", () => {
    it("should render if tutorialSeen is false", () => {
        render(<Router  tutorialSeen={false} />)
    })
})
