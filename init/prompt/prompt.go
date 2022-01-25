package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func GetReplaceValues() (ReplaceValues, error) {
	replaceValues := replaceValuesBase

	for index, rv := range replaceValues {
		prompt := &survey.Input{
			Message: rv.Description,
			Default: rv.OldValue,
		}
		answer := ""

		err := survey.AskOne(prompt, &answer)
		if err != nil {
			return nil, err
		}

		if answer == "" {
			return nil, fmt.Errorf(`answer for "%s" is empty (default: %s)`, rv.Description, rv.OldValue)
		}

		replaceValues[index].NewValue = answer
	}

	return replaceValues, nil
}
