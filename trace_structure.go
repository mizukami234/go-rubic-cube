package main

import (
	"github.com/cheggaaa/pb"
	"fmt"
	"sort"
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
	var known_states = tr.CubeSet{INIT}
	var known_identities = []tr.Path{}
	return search_moves_rec([]tr.Move{tr.Move{tr.Path{}, INIT}}, known_states, known_identities)
}

func search_moves_rec(prev_moves []tr.Move, known_states tr.CubeSet, known_identities []tr.Path) tr.CubeSet {
	fmt.Printf("----\n")
	depth++
	fmt.Printf("DEPTH: %d\n", depth)

	search_count := len(prev_moves)
	fmt.Printf("searching new steps: %d\n", (search_count * 6))

	var count_ids, count_knowns int

	bar := pb.StartNew(search_count)

	moves := []tr.Move{}
	new_known_states := tr.CubeSet{}
	for _, prev_move := range prev_moves {
		bar.Increment()
		new_steps := prev_move.NextSteps()
		this_known_states := tr.CubeSet{}
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
			if known_states.HasCube(new_step.Cube) {
				count_knowns++
				continue
			}
			// new known state
			if new_known_states.HasCube(new_step.Cube) {
				count_knowns++
				continue
			}
			// otherwise register known state and new move
			moves = append(moves, new_step)
			this_known_states = append(this_known_states, new_step.Cube)
			//sort.Sort(known_states)
		}
		new_known_states = append(new_known_states, this_known_states...)
		sort.Sort(new_known_states)
	}

	known_states = append(known_states, new_known_states...)
	sort.Sort(known_states)

	bar.FinishPrint("finished.")

	fmt.Printf("identity paths:       %d\n", count_ids)
	fmt.Printf("found known states:   %d\n", count_knowns)
	fmt.Printf("found new paths:      %d\n", len(moves))
	fmt.Printf("percent uncultivated: %f\n", float64(len(moves))/float64(len(prev_moves)*6))
	fmt.Printf("current known ids:    %d\n", len(known_identities))
	fmt.Printf("current known states: %d\n", len(known_states))
	if (len(moves) > 0) {
		return search_moves_rec(moves, known_states, known_identities)
	} else {
		return known_states
	}
}

func main(){
	search_all_states()
}
