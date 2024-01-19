// Conversion for various objects
package hiconverters

import "regexp"

type TextConversionRule interface {
	ConvertAll(input string) string
}

type TextConversionRuleDefinition struct {
	Matchingregexp *regexp.Regexp
	Replacement    string
}

func NewTextConversionRuleDefinition(matching string, replacement string) TextConversionRuleDefinition {
	return TextConversionRuleDefinition{
		regexp.MustCompile(matching),
		replacement,
	}
}
