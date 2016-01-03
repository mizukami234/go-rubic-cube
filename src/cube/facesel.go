package cube

/** 面のセレクタ(face selector)集合

  cube[i][j]でアクセスできる面に対し、[i, j]を面のセレクタと呼ぶことにする。
  sel.Select(cube)で面を取得する

  説明上はセレクタsに対し、次の面、前の面をs+、s-、対称の面をs*と表すことにする。

  (群論的にはセレクタは2x3の群になる)

 */

type FaceSel [2]int

var (
	N_FACESELS = [3]FaceSel{ FaceSel{N, 0}, FaceSel{N, 1}, FaceSel{N, 2}, }
	S_FACESELS = [3]FaceSel{ FaceSel{S, 0}, FaceSel{S, 1}, FaceSel{S, 2}, }
	FACESELS = [2][3]FaceSel{ N_FACESELS, S_FACESELS }
)

func (fs FaceSel) Add(ft FaceSel) FaceSel {
	return FACESELS[(fs[0] + ft[0]) % 2][(fs[1] + ft[1]) % 3]
}

/* セレクタの同値 */
func (fs FaceSel) Equal(ft FaceSel) bool {
	return ((fs[0] == ft[0]) && (fs[1] == ft[1]))
}

/* 対称の面のセレクタを取得 */
func (fs FaceSel) Symmetry() FaceSel {
	return fs.Add(FaceSel{1, 0})
}

/* ルービックキューブとセレクタから面を取得 */
func (fs FaceSel) Select(c Cube) Face {
	return c[fs[0]][fs[1]]
}

