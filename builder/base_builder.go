package builder

import (
	"fmt"
	"strings"
)

type baseBuilder struct {
	client APIClient
	errors []error
}

func (b *baseBuilder) toStringList(values []int64) []string {
	keys := []string{}
	for _, k := range values {
		keys = append(keys, fmt.Sprintf("%d", k))
	}
	return keys
}

func (b *baseBuilder) getFlattenErrors() error {
	var list = make([]string, 0)
	for _, str := range b.errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}
