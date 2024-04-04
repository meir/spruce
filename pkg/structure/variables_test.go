package structure

import (
	"fmt"
	"strings"
	"testing"

	"oss.terrastruct.com/util-go/assert"
)

type IntVariable struct {
	x int
}

func NewIntVariable(x int) *IntVariable {
	return &IntVariable{x: x}
}

func (i *IntVariable) Get() any {
	return i.x
}

func (i *IntVariable) Set(value any) {
	if value, ok := value.(int); ok {
		i.x = value
	}
}

func (i *IntVariable) String() string {
	return fmt.Sprint(i.x)
}

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
		output = append(output, fmt.Sprintf("%s=\"%s\"", k, v))
	}
	return strings.Join(output, " ")
}

func TestIntVariable(t *testing.T) {
	i := NewIntVariable(42)
	assert.Equal(t, 42, i.Get())
	i.Set(43)
	assert.Equal(t, 43, i.Get())
	assert.Equal(t, "43", i.String())
}

func TestScope(t *testing.T) {
	s := NewScope()
	i := NewIntVariable(42)
	s.Set("foo", i)
	assert.Equal(t, 42, s.Get("foo").Get())
	i.Set(43)
	s.Set("foo", i)
	assert.Equal(t, 43, s.Get("foo").Get())
	assert.Equal(t, "43", s.Get("foo").String())
}

func TestScopeParent(t *testing.T) {
	s := NewScope()
	i := NewIntVariable(42)
	s.Set("foo", i)

	s2 := NewScopeWithParent(s)
	assert.Equal(t, 42, s2.Get("foo").Get())

	i2 := s2.Get("foo").(*IntVariable)
	i2.Set(43)
	assert.Equal(t, 43, s.Get("foo").Get())
}

func SetMapVariable(s *Scope) {
	attributes := s.Get("foo")
	if m, ok := Get[map[string]any](attributes); ok {
		m["foo"] = 43
		Set(attributes, m)
	}
}

func TestMapVariableWithParent(t *testing.T) {
	s := NewScope()
	m := NewMapVariable(map[string]interface{}{"foo": 42})
	s.Set("foo", m)
	s2 := NewScopeWithParent(s)
	o := s2.Get("foo").Get().(map[string]any)
	assert.Equal(t, 42, o["foo"].(int))

	SetMapVariable(s2)
	o2 := s.Get("foo").Get().(map[string]any)
	assert.Equal(t, 43, o2["foo"])
}

func TestMapVariableWithParentInline(t *testing.T) {
	s := NewScope()
	m := NewMapVariable(map[string]interface{}{"foo": 42})
	s.Set("foo", m)
	s2 := NewScopeWithParent(s)
	o := s2.Get("foo").Get().(map[string]any)
	assert.Equal(t, 42, o["foo"].(int))

	attributes := s.Get("foo")
	if m, ok := Get[map[string]any](attributes); ok {
		m["foo"] = 43
		Set(attributes, m)
	}
	o2 := s.Get("foo").Get().(map[string]any)
	assert.Equal(t, 43, o2["foo"])
}

func TestMapVariableWithDoubleParents(t *testing.T) {
	scope1 := NewScope()
	scope2 := NewScopeWithParent(scope1)
	m := NewMapVariable(map[string]interface{}{"foo": 42})
	scope2.Set("foo", m)
	scope3 := NewScopeWithParent(scope2)

	output1, _ := Get[map[string]any](scope2.Get("foo"))
	assert.Equal(t, 42, output1["foo"].(int))

	SetMapVariable(scope3)
	output2, _ := Get[map[string]any](scope2.Get("foo"))
	assert.Equal(t, 43, output2["foo"].(int))
}

func TestMapVariableReverseParent(t *testing.T) {
	scope1 := NewScope()
	m := NewMapVariable(map[string]interface{}{"foo": 42})
	scope1.Set("foo", m)

	scope2 := NewScopeWithParent(scope1)
	scope3 := NewScopeWithParent(scope2)

	output1, _ := Get[map[string]any](scope1.Get("foo"))
	assert.Equal(t, 42, output1["foo"].(int))

	SetMapVariable(scope1)
	output2, _ := Get[map[string]any](scope3.Get("foo"))
	assert.Equal(t, 43, output2["foo"].(int))
}
