package errorImpl

import "errors"

//ScenarioError - create custom scenario error
func ScenarioError(customText string, err error) error {
	return errors.New("Scenario: " + customText + ": " + err.Error())
}

//AdminsError - create custom admins error
func AdminsError(customText string, err error) error {
	return errors.New("Admins: " + customText + ": " + err.Error())
}

//AuthError - create custom auth error
func AuthError(customText string, err error) error {
	return errors.New("Auth: " + customText + ": " + err.Error())
}

//ServiceError - create custom auth error
func ServiceError(customText string, err error) error {
	return errors.New("Service: " + customText + ": " + err.Error())
}
