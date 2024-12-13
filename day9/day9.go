package day9

import (
	"fmt"

	"github.com/quollveth/AdventOfGode/util"
)

type block struct {
	// id == -1 -> empty block
	lenght, id int
}

func Run() {
	input := util.ReadFileFull("day9/tinyin")
	inLen := len(input)
	_ = inLen

	// blocks before reordering
	preBlocks := []block{}

	isBlock := true
	id := -1

	// read blocks into array
	for _, c := range input {
		num := int(c - '0')
		value := func() int {
			if isBlock {
				id++
				return id
			}
			return -1
		}()

		b := block{
			lenght: num,
			id:     value,
		}

		preBlocks = append(preBlocks, b)

		isBlock = !isBlock
	}

	// reorder blocks
	/*
		go through blocks copying them to a new array
		if we reach an an empty one get the last element and place it here as much as possible
	*/

	var lastIndex int       // index of the last block, decreases by 1 each time we use it
	postBlocks := []block{} // blocks after reordering

	for {
		lastIndex = len(preBlocks) - 1
		if lastIndex == -1 {
			// this should never happen
			break
		}

		cblock := preBlocks[0]
		preBlocks = util.RemoveAtIndex(preBlocks, 0)
		lastIndex--

		// move blocks from preblocks to postblocks
		// once a block is moved it gets deleted
		// preblocks is modified after each move because im tired of this
		// if it's a non empty block move it and it's all good
		// if it's an empty block we need to fill it with data from the last block

		// non empty block
		if cblock.id != -1 {
			postBlocks = append(postBlocks, cblock)
			continue
		}

		// empty block
		// move it and take data from the last block
		// decrease the lenght of the last block
		// when it runs out delete it
		blockToFill := block{}
		blockToFill.lenght = cblock.lenght

		postBlocks = append(postBlocks, blockToFill)
	}
	fmt.Println(postBlocks)
}
