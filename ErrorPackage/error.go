package errordict

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var errorMap map[string]string

func init() {
	if err := loadErrorMap("ErrorPackage/error.json"); err != nil {
		panic(err)
	}
}

func loadErrorMap(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	errMap := make(map[string]string)
	if err := json.Unmarshal(data, &errMap); err != nil {
		return err
	}

	errorMap = errMap
	return nil
}

func GetErrorMessage(code string) (string, error) {
	errMsg, ok := errorMap[code]
	if !ok {
		return "", fmt.Errorf("error code %s not found", code)
	}
	return errMsg, nil
}

// GetErrorMessageByCode is a function to retrieve the error message based on the error code provided.
func GetErrorMessageByCode(code string) (string, error) {
	errMsg, err := GetErrorMessage(code)
	if err != nil {
		return "", err
	}
	return errMsg, nil
}

func Test() string {
	return "hi"
}
