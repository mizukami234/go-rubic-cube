package tracer

import (
	. "cube"
)

/**

  原始回転
  ある面のlayer=2の回転は、ある面でのlayer=0の逆回転になる。
  したがって、layer=2の回転は最小表現には必要ない。
  同様に、layer=1の回転は、2, 0の回転で表現できる。このとき回転して同値な図形を同一視する。
  以上から、原始的な回転操作は面に対応する6通りしか存在しない。
  原始回転操作はセレクタによって表現する。

 */

var Primitives = [6]FaceSel{
	FaceSel{N, 0},
	FaceSel{N, 1},
	FaceSel{N, 2},
	FaceSel{S, 0},
	FaceSel{S, 1},
	FaceSel{S, 2},
}

func PrimitiveRotate(c Cube, fs FaceSel) Cube {
	return c.Rotate(fs, 0)
}
