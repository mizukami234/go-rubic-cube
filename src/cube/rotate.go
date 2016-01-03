package cube

/**

  回転操作

  回転操作は、次のように定める。
  まず、面に対応する軸ベクトルを考える。このベクトルは、面からキューブの中心に向かうベクトルとする。
  軸ベクトルに対して、その始点側からキューブの回転層を0, 1, 2と番号付ける。
  一回の"正の回転"を、この回転層を軸ベクトルに対し右ねじの方向に回す操作とする。

  以上から、面と回転層を得ることで原始的な回転操作を定義できる。

  次の二点に注意する。

  [1] 回転層は0, 2の場合、特定の面が回転する。その様子は以下の様な形である。

  [N]___________     [N]___________     [N]___________
    |  2|  5|  8|      |  0|  1|  2|      |  6|  3|  0|
    |---|---|---|  -   |---|---|---|  +   |---|---|---|
    |  1|  4|  7|  <=  |  3|  4|  5|  =>  |  7|  4|  1|
    |---|---|---|      |---|---|---|      |---|---|---|
    |  0|  3|  6|      |  6|  7|  8|      |  8|  5|  2|
    ^^^^^^^^^^^^^      ^^^^^^^^^^^^^      ^^^^^^^^^^^^^

  [S]___________     [S]___________     [S]___________
    |  6|  7|  8|      |  0|  3|  6|      |  2|  1|  0|
    |---|---|---|  -   |---|---|---|      |---|---|---|
    |  3|  4|  5|  <=  |  1|  4|  7|  =>  |  5|  4|  3|
    |---|---|---|      |---|---|---|      |---|---|---|
    |  0|  1|  2|      |  2|  5|  8|      |  8|  7|  6|
    ^^^^^^^^^^^^^      ^^^^^^^^^^^^^      ^^^^^^^^^^^^^

  Nを正に回転させる変換と、Sを負に回転させる変換が同値になる。逆も同様。
  よってNを正に回転させる操作を正回転、Nを負に回転させる操作を負回転とし、それらのみ定義する。

  [2] 軸ベクトルpがNのとき、pを通らない4面(周)が一列になるように展開すると以下のようになる。

       [ p* ]
  [ p- | p+ | p*-| p*+]
       [ p  ]

  ここで[ p- | p+ | p*-| p*+]の各面はそれぞれ、右下N、左下N、右上S、左上Sに極をもつ。
  さらにそのn層における番号は

  p-  [3n+2, 3n+1, 3n  ]
  p+  [n,    3+n,  6+n ]
  p*- [8-3n, 7-3n, 6-3n]
  p*+ [2-n,  5-n,  8-n ]

  pがSのときは、同様に

       [ p* ]
  [ p+ | p- | p*+| p*-]
       [ p  ]

  ここで[ p+ | p- | p*+| p*-]の各面はそれぞれ、右下S、左下S、右上N、左上Nに極をもつ。

  p+  [6+n,  3+n,  n   ]
  p-  [3n,   3n+1, 3n+2]
  p*+ [8-n,  5-n,  2-n ]
  p*- [6-3n, 7-3n, 8-3n]

*/

type FaceTransformer [9]int

func (ft FaceTransformer) Transform(face Face) (new_face Face) {
	for i, _ := range face {
		new_face[i] = face[ft[i]]
	}
	return
}

var (
	ROT_NEXT_N = FaceTransformer{6, 3, 0, 7, 4, 1, 8, 5, 2}
	ROT_PREV_N = FaceTransformer{2, 5, 8, 1, 4, 7, 0, 3, 6}
	ROT_NEXTS = [2]FaceTransformer{ROT_NEXT_N, ROT_PREV_N}
	ROT_PREVS = [2]FaceTransformer{ROT_PREV_N, ROT_NEXT_N}

	N_ROUND_FACESEL_OFFSETS = [4]FaceSel{FaceSel{0, 2}, FaceSel{0, 1}, FaceSel{1, 2}, FaceSel{1, 1}}
	N_ROUND_FACE_SCALS = [4]int{3, 1, -3, -1};
	N_ROUND_FACE_CONSTS = [4][3]int{[3]int{2, 1, 0}, [3]int{0, 3, 6}, [3]int{8, 7, 6}, [3]int{2, 5, 8}}
	S_ROUND_FACESEL_OFFSETS = [4]FaceSel{FaceSel{0, 1}, FaceSel{0, 2}, FaceSel{1, 1}, FaceSel{1, 2}}
	S_ROUND_FACE_SCALS = [4]int{1, 3, -1, -3}
	S_ROUND_FACE_CONSTS = [4][3]int{[3]int{6, 3, 0}, [3]int{0, 1, 2}, [3]int{8, 5, 2}, [3]int{6, 7, 8}}
	ROUND_FACESEL_OFFSETS = [2][4]FaceSel{N_ROUND_FACESEL_OFFSETS, S_ROUND_FACESEL_OFFSETS}
	ROUND_FACE_SCALS = [2][4]int{N_ROUND_FACE_SCALS, S_ROUND_FACE_SCALS}
	ROUND_FACE_CONSTS = [2][4][3]int{N_ROUND_FACE_CONSTS, S_ROUND_FACE_CONSTS}
)

type RoundFace struct {
	facesel FaceSel
	indices [3]int
}

type Round [4]RoundFace

func (fs FaceSel) getRound(layer int) (round Round) {
	for i, offset := range ROUND_FACESEL_OFFSETS[fs[0]] {
		scal := ROUND_FACE_SCALS[fs[0]][i]
		consts := ROUND_FACE_CONSTS[fs[0]][i]

		round[i].facesel = fs.Add(offset)
		for j, c := range consts {
			round[i].indices[j] = scal * layer + c
		}
	}
	return
}

func (r Round) indexOfFaceSel(fs FaceSel) int {
	for i, roundface := range r {
		if roundface.facesel.Equal(fs) {
			return i
		}
	}
	return -1
}

/* 周の面を回す */
func (cube Cube) rotateRound(pivot FaceSel, layer int) Cube {
	round := pivot.getRound(layer)
	return cube.MapFaces(func(target_face Face, sel FaceSel) (new_face Face) {
		index := round.indexOfFaceSel(sel)
		if index == -1 {
			return target_face // face is not one of round faces
		}
		source_roundface := round[(index + 3) % 4]
		target_roundface := round[index]
		source_face := source_roundface.facesel.Select(cube)
		for index, col := range target_face {
			ii := -1
			for jj, jndex := range target_roundface.indices {
				if index == jndex {
					ii = jj
					break
				}
			}
			if ii == -1 {
				new_face[index] = col
			} else {
				new_face[index] = source_face[source_roundface.indices[ii]]
			}
		}
		return
	})
}

/* 回転操作 */
func (cube Cube) Rotate(pivot FaceSel, layer int) (new_cube Cube) {
	base_polar := pivot[0]
	new_cube = cube.rotateRound(pivot, layer)
	if layer != 1 {
		var (
			rot_facesel FaceSel
			face_transformer FaceTransformer
		)
		if layer == 0 {
			rot_facesel = pivot
			face_transformer = ROT_NEXTS[base_polar]
		} else {
			rot_facesel = pivot.Symmetry()
			face_transformer = ROT_PREVS[base_polar]
		}
		new_cube = new_cube.MapFaces(func(face Face, sel FaceSel) Face {
			if rot_facesel.Equal(sel) {
				return face_transformer.Transform(face)
			} else {
				return face
			}
		})
	}
	return
}
