package validation

import "fmt"

func GetPostValidate(id int) error {
	if id < 0 {
		return fmt.Errorf("invalid id")
	}
	return nil
}
