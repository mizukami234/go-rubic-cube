package main

import (
	"github.com/cheggaaa/pb"
	"fmt"
	. "cube"
	tr "tracer"
)

/**

  [1] 雑に全通り検索する

  原始操作の列を順番に作りながら、知っている状態に戻るような操作の場合は探索をやめることで、全状態を取得する。

 */

var INIT = CreateInitialCube()
var depth = 0

func include_suffix(suffix_paths []tr.Path, path tr.Path) bool {
	for _, suffix_path := range suffix_paths {
		if path.HasSuffix(suffix_path) {
			return true
		}
	}
	return false
}

func search_all_states() tr.CubeSet {
	//var known_states = tr.CubeSet{INIT}
	var known_cubes = &tr.ManagedCubeSet{}
	var known_identities = []tr.Path{}
	return search_moves_rec([]tr.Move{tr.Move{tr.Path{}, INIT}}, known_cubes, known_identities)
}

func search_moves_rec(prev_moves []tr.Move, known_cubes *tr.ManagedCubeSet, known_identities []tr.Path) tr.CubeSet {
	fmt.Printf("----\n")
	depth++
	fmt.Printf("DEPTH: %d\n", depth)

	search_count := len(prev_moves)
	fmt.Printf("searching new steps: %d\n", (search_count * 6))

	var count_ids, count_knowns int

	bar := pb.StartNew(search_count)

	moves := []tr.Move{}
	for _, prev_move := range prev_moves {
		bar.Increment()
		new_steps := prev_move.NextSteps()
		for _, new_step := range new_steps {
			// new step returns initial state
			if new_step.Cube.Equal(INIT) {
				known_identities = append(known_identities, new_step.Path)
				count_ids++
				continue
			}
			// new step has loop path then it's a known state
			if include_suffix(known_identities, new_step.Path) {
				count_knowns++
				continue
			}
			// new step got known state
			if known_cubes.HasCube(new_step.Cube) {
				count_knowns++
				continue
			}
			// otherwise register known state and new move
			moves = append(moves, new_step)
			known_cubes.AddCube(new_step.Cube)
		}
	}

	known_cubes.ForceSort()

	bar.FinishPrint("finished.")

	fmt.Printf("identity paths:       %d\n", count_ids)
	fmt.Printf("found known cubes:    %d\n", count_knowns)
	fmt.Printf("found new paths:      %d\n", len(moves))
	fmt.Printf("percent uncultivated: %f\n", float64(len(moves))/float64(len(prev_moves)*6))
	fmt.Printf("current known ids:    %d\n", len(known_identities))
	fmt.Printf("current known cubes:  %d\n", known_cubes.Len())
	if (len(moves) > 0) {
		return search_moves_rec(moves, known_cubes, known_identities)
	} else {
		return known_cubes.GetCubes()
	}
}

func main(){
	search_all_states()
}
