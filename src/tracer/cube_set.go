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

type ManagedCubeSet struct{
	sortedSet CubeSet
	unsortedSet CubeSet
}

func (mcs *ManagedCubeSet) AddCube(c Cube) {
	mcs.unsortedSet = append(mcs.unsortedSet, c)
	if len(mcs.unsortedSet) == 10000 {
		mcs.ForceSort()
	}
}

func (mcs *ManagedCubeSet) ForceSort() {
	mcs.sortedSet = append(mcs.sortedSet, mcs.unsortedSet...)
	sort.Sort(mcs.sortedSet)
	mcs.unsortedSet = CubeSet{}
}

func (mcs *ManagedCubeSet) Len() int {
	return len(mcs.sortedSet)+len(mcs.unsortedSet)
}

func (mcs *ManagedCubeSet) HasCube(c Cube) bool {
	if mcs.sortedSet.HasCube(c) { return true }
	if mcs.unsortedSet.include(c) { return true }
	return false
}

func (mcs *ManagedCubeSet) GetCubes() CubeSet {
	mcs.ForceSort()
	return mcs.sortedSet
}
