// Package twofer shows how to format string
package twofer

import ( 
    "fmt"
    "strings"
)

// ShareWith generates a formatted sentence
func ShareWith(name string) string {
    if name = strings.TrimSpace(name); name == "" {
        name = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", name)
}
