package tools

import (
	"fmt"
	"os"
)

func ValidateFile(file string) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory!", file)
	}
	return nil
}
