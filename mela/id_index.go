package mela

import (
	"fmt"

	"github.com/google/uuid"
)

type IDIndex map[string]string

func (i IDIndex) AddName(name string) {
	i[name] = uuid.NewString()
}

func (i IDIndex) ID(name string) (string, bool) {
	id, ok := i[name]
	return id, ok
}

func (i IDIndex) LinkName(name string) (string, bool) {
	id, ok := i.ID(name)
	if !ok {
		return "", false
	}

	return fmt.Sprintf("[%s](mela://recipe/%s)", name, id), true
}
