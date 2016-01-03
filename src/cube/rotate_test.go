package cube

import "testing"

func TestCubeRotate(t *testing.T) {

	cube := CreateInitialCube()
	rotated := cube.Rotate(FaceSel{N, 0}, 0)
	expected := Cube{
		HalfCube{
			Face{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Face{2, 1, 1, 2, 1, 1, 2, 1, 1},
			Face{4, 4, 4, 2, 2, 2, 2, 2, 2},
		},
		HalfCube{
			Face{3, 3, 3, 3, 3, 3, 3, 3, 3},
			Face{4, 4, 5, 4, 4, 5, 4, 4, 5},
			Face{5, 5, 5, 5, 5, 5, 1, 1, 1},
		},
	}

	if !rotated.Equal(expected) {
		t.Errorf("Original Cube\n%vRotated Cube for (FaceSel{N, 0}, 0)\n%vExpecting Cube\n%v", cube.String(), rotated.String(), expected.String())
	}

}
