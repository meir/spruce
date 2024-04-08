package spruce

import (
	"github.com/meir/spruce/pkg/states"
	"github.com/meir/spruce/pkg/structure"
)

func Parse(file string) (*structure.File, error) {
	return states.Parse(file)
}
