package utilities

import (
	"os"
	"strconv"
)

func ConvertToInt(s string) (int, error) {
	newInst, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return newInst, nil
}

func IsLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}
