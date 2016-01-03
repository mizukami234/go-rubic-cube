package cube

import (
	"bytes"
	"strconv"
)

/**

ルービックキューブの8頂点のうち点対称な2点をとり（極ということにする）、N、Sとする。
このとき各面はN,Sいずれかの極を含んでいる。
N極を含む3つの面をN0, N1, N2とし、その対称となる面をS0, S1, S2とする。
但し番号の付け方は、極を上から見た時にNの場合反時計回り、Sの場合時計回りになるようにする。

面の一色で塗られる小さい面(小面)を以下のように番号付けする

[N]____________  [S]____________
|  0|  1|  2|    |  0|  3|  6|
|---|---|---|    |---|---|---|
|  3|  4|  5|    |  1|  4|  7|
|---|---|---|    |---|---|---|
|  6|  7|  8|    |  2|  5|  8|
^^^^^^^^^^^^^    ^^^^^^^^^^^^^

さらに、例えば面N2の場合には

N2_____________
|N20|N21|N22|
|---|---|---|
|N23|N24|N25|
|---|---|---|
|N26|N27|N28|
^^^^^^^^^^^^^

と表すようにして、各小面をすべて番号付ける。
以上により、2*3*9次元の行列表現とする。

*/
type Color int
type Face [9]Color
type HalfCube [3]Face
type Cube [2]HalfCube

const (
	N = 0
	S = 1
)

func CreateInitialCube() Cube {
	return Cube{
		// N0-2
		HalfCube{
			Face{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Face{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Face{2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
		// S0-2
		HalfCube{
			Face{3, 3, 3, 3, 3, 3, 3, 3, 3},
			Face{4, 4, 4, 4, 4, 4, 4, 4, 4},
			Face{5, 5, 5, 5, 5, 5, 5, 5, 5},
		},
	}
}

func (c Cube) Equal(d Cube) bool {
	polar := 0
	for polar < 2 {
		n := 0
		for n < 3 {
			i := 0
			for i < 9 {
				if (c[polar][n][i] != d[polar][n][i]) { return false }
				i++
			}
			n++
		}
		polar++
	}
	return true
}

func (c Cube) Less(d Cube) bool {
	polar := 0
	for polar < 2 {
		n := 0
		for n < 3 {
			i := 0
			for i < 9 {
				if (c[polar][n][i] < d[polar][n][i]) { return true }
				i++
			}
			n++
		}
		polar++
	}
	return false
}

func (c Cube) MapFaces(f func(Face, FaceSel) Face) (new_cube Cube) {
	polar := 0
	for polar < 2 {
		n := 0
		for n < 3 {
			sel := FaceSel{polar, n}
			new_cube[polar][n] = f(sel.Select(c), sel)
			n++
		}
		polar++
	}
	return
}

var cube_map_str = [9][]int{
	[]int{ -1, -1, -1, 32, 31, 30 },
	[]int{ -1, -1, -1, 35, 34, 33 },
	[]int{ -1, -1, -1, 38, 37, 36 },
	[]int{ 28, 27, 26, 12, 15, 18, 52, 51, 50, 40, 43, 46 },
	[]int{ 25, 24, 23, 11, 14, 17, 55, 54, 53, 41, 44, 47 },
	[]int{ 22, 21, 20, 10, 13, 16, 58, 57, 56, 42, 45, 48 },
	[]int{ -1, -1, -1,  0,  1,  2 },
	[]int{ -1, -1, -1,  3,  4,  5 },
	[]int{ -1, -1, -1,  6,  7,  8 },
}

func (c Cube) String() string {
	var buffer bytes.Buffer
	for _, line := range cube_map_str {
		for _, magic := range line {
			if magic == -1 {
				buffer.WriteString("  ")
			} else {
				var (
					face_n = int(magic / 10)
					polar = int(face_n / 3)
					n = face_n % 3
					i = magic % 10
				)
				buffer.WriteString(" "+strconv.Itoa(int(c[polar][n][i])))
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

