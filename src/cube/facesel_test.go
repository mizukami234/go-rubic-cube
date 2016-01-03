package cube

import "testing"

func TestFaceSelAdd(t *testing.T) {
	cases := []struct {
		x, y, expected FaceSel
	}{
		{ FaceSel{N, 1}, FaceSel{0, 2}, FaceSel{N, 0} },
		{ FaceSel{N, 1}, FaceSel{1, 0}, FaceSel{S, 1} },
	}

	for _, c := range cases {
		got := c.x.Add(c.y)
		if (got[0] != c.expected[0]) || (got[1] != c.expected[1]) {
			t.Errorf("FaceSel{%v}.Add(FaceSel{%v}) == FaceSel{%v}, expecting FaceSel{%v}.", c.x, c.y, got, c.expected)
		}
	}
}

func TestFaceSelEqual(t *testing.T) {
	cases := []struct {
		x, y FaceSel
		expected bool
	}{
		{ FaceSel{N, 1}, FaceSel{N, 2}, false },
		{ FaceSel{S, 2}, FaceSel{S, 2}, true },
	}

	for _, c := range cases {
		got := c.x.Equal(c.y)
		if (got != c.expected) {
			t.Errorf("FaceSel{%v}.Equal(FaceSel{%v}) == %v, expecting %v.", c.x, c.y, got, c.expected)
		}
	}
}

func TestFaceSelSymmetry(t *testing.T) {
	x := FaceSel{N, 0}
	y := x.Symmetry()
	ex := FaceSel{S, 0}
	if (y[0] != ex[0]) || (y[1] != ex[1]) {
		t.Errorf("FaceSel{%v}.Symmetry() == FaceSel{%v}, expecting FaceSel{%v}.", x, y, ex)
	}
}
