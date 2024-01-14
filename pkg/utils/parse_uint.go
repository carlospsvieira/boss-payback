package utils

import "strconv"

func ParseUint(id string) (uint, error) {
	parsedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(parsedId), nil
}
