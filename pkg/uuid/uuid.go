package uuid

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
)

func GenerateUUID() {
	u, err := uuid.NewRandom()

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	uuidString := u.String()

	// Copy to clipboard
	err = clipboard.WriteAll(uuidString)

	if err != nil {
		fmt.Printf("Failed to copy to clipboard: %s", err)
	} else {

		fmt.Println("UUID copied to clipboard:", uuidString)
	}
}
