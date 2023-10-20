package controller


import (
	"fmt"
	"cli/models"
)

func ApiRunController() error {
	output, err := models.CreateArchive()

	// TODO: Remove after debug 
	fmt.Println("The output value is:", output)
	fmt.Println("The err value is:", err)

	return err
}


