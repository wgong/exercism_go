// Package proverb generates sayings with input rhyme words
package proverb

import "fmt"

// Proverb produces templated proverb
func Proverb(rhyme []string) []string {
	sentences := []string{}
	templates := [2]string{"For want of a %s the %s was lost.",
		"And all for the want of a %s."}
	for i := range rhyme {
		if i < len(rhyme)-1 {
			sentences = append(sentences, fmt.Sprintf(templates[0], rhyme[i], rhyme[i+1]))
		} else {
			sentences = append(sentences, fmt.Sprintf(templates[1], rhyme[0]))
		}
	}
	return sentences
}
