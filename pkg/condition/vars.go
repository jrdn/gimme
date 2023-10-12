package condition

import "strings"

func NewVars() *Vars {
	return &Vars{x: make(map[string]any)}
}

func NewVarsWithData(data map[string]any) *Vars {
	x := NewVars()
	for k, v := range data {
		x.Set(k, v)
	}
	return x
}

type Vars struct {
	x map[string]any
}

func (v *Vars) Get(key string) any {
	return v.x[strings.ToLower(key)]
}

func (v *Vars) Set(key string, value any) {
	v.x[strings.ToLower(key)] = value
}
