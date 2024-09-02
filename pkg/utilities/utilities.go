package utilities

import "strconv"

func ConvertToInt(s string) (int, error) {
	newInst, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return newInst, nil
}
