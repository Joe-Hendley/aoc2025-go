package file

import (
	"fmt"
	"os"
)

func DangerReadToString(path string) string {
	s, err := ReadToString(path)

	if err != nil {
		panic(err)
	}

	return s
}

func ReadToString(path string) (string, error) {
	f, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	s := string(f)
	if len(s) == 0 {
		return "", fmt.Errorf("empty file")
	}

	return s, nil
}
