package algorithms

import (
	"fmt"
)


func ShareWith03(name string) string {
    if len(name) == 0 {
        name = "you"
    }
	return fmt.Sprintf("One for %s, one for me.", name)
}