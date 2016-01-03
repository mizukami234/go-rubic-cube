package tracer

import (
	"sort"
	. "cube"
)

type CubeSet []Cube

func (cs CubeSet) Len() int {
	return len(cs)
}

func (cs CubeSet) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs CubeSet) Less(i, j int) bool {
	return cs[i].Less(cs[j])
}

func (cs CubeSet) HasCube(c Cube) bool {
	i := sort.Search(cs.Len(), func (i int) bool {
		return !cs[i].Less(c)
	})
	if i == cs.Len() {
		return false
	} else {
		return cs[i].Equal(c)
	}
}

func (cs CubeSet) include(c Cube) bool {
	for _, d := range cs {
		if c.Equal(d) {
			return true
		}
	}
	return false
}
