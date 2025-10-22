package intermediate

import (
	"fmt"
	"time"
)

// string - sequence of bytes, immutable, array of unicode characters(rune)

func runes() {
	message := "hello world!"
	rawMessage := "hello\n go"
	defer fmt.Println(message, rawMessage)

	// carriage return will keep on overwriting on the same line
	for i := 1; i <= 5; i++ {
		fmt.Printf("\rProcessing %d%%", i*20)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("\nDone!")

	fmt.Println(len(rawMessage))
}
