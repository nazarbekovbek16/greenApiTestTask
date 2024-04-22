package service

import (
	"encoding/json"
	"strconv"
)

func handleMQ(msg []byte) ([]byte, error) {
	var data map[string]string
	if err := json.Unmarshal(msg, &data); err != nil {
		return nil, err
	}

	param := data["param"]
	result := double(param)

	return []byte(result), nil
}

func double(param string) string {
	doubled := 2 * atoi(param)
	return strconv.Itoa(doubled)
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
