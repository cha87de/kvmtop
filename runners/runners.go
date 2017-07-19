package runners

import "fmt"

// InitializeRunners starts necessary runners as threads
func InitializeRunners() error {

	initializeLookup()
	//go initializeCollect()

	fmt.Print("done with runners")

	return nil
}
