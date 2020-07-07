package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

func requitedIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {

		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}