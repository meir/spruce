package variables

import (
	"fmt"
	"strings"
)

type MapVariable struct {
	_map map[string]interface{}
}

func NewMapVariable(m map[string]interface{}) *MapVariable {
	return &MapVariable{_map: m}
}

func (m *MapVariable) Get() any {
	return m._map
}

func (m *MapVariable) Set(value any) {
	if value, ok := value.(map[string]interface{}); ok {
		m._map = value
	}
}

func (m *MapVariable) String() string {
	output := []string{}
	for k, v := range m._map {
		output = append(output, fmt.Sprintf("%s=\"%v\"", k, v))
	}
	return strings.Join(output, " ")
}
