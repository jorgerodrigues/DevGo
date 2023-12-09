package snippets

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func GenerateReactComponent() {
	const componentTemplate = `
    const {{ ComponentName }} = () => {
      return (
        <div></div>
      )
    }`
  err := clipboard.WriteAll(componentTemplate)

	if err != nil {
		fmt.Printf("Failed to copy to clipboard: %s", err)
	} else {

		fmt.Println("Component template copied to clipboard")
	}
}
