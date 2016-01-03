package tracer

import (
	. "cube"
)

/**

  操作
  原始回転の列を操作よぶことにする。操作はルービックキューブをある状態から別の状態に変換する。
  何もしないものは空の操作[]である。
  操作は左から順番に適用される。

 */

type Path []FaceSel

func (p Path) AddNext(prim FaceSel) Path {
	return append(p, prim)
}

func (p Path) HasSuffix(suffix_path Path) bool {
	p_len := len(p)
	offset := p_len - len(suffix_path)
	for i, prim := range suffix_path {
		if p[offset + i] != prim {
			return false
		}
	}
	return true
}

type Move struct {
	Path Path
	Cube Cube
}

func (move Move) NextSteps() (new_moves [6]Move) {
	for i, prim := range Primitives {
		state := PrimitiveRotate(move.Cube, prim)
		new_moves[i] = Move{move.Path.AddNext(prim), state}
	}
	return
}

