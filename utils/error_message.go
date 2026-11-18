package utils

import "fmt"

func RequiredMessage(entityName, coloumn string) error {
	return fmt.Errorf("Please input %s for entity %s", coloumn, entityName)
}

func EnumMessage(entityName, coloumn string) error {
	return fmt.Errorf("Input value for %s for entity %s doesn't match", coloumn, entityName)
}
