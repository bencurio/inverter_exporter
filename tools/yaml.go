package tools

import (
	"errors"
	"fmt"
	"os"
)

func ValidateFile(file string) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New(fmt.Sprintf("'%s' is a directory!", file))
	}
	return nil
}
