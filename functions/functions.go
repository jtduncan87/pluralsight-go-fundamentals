package functions

import (
	"strings"
)

// Converter module to titleCase and author to Upper
func Converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
	return module, author
}
