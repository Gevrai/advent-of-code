package main

import (
	"advent-of-code-2019/utils"
)

func main() {
	input := utils.ReadInputFileRelative()

	const width = 25
	const height = 6

	layers := extractLayers(input[0], width, height)

	layerWithFewerZeros := layers[0]
	minZeros := layerWithFewerZeros.HowMany(0)
	for _, layer := range layers[1:] {
		zeros := layer.HowMany(0)
		if zeros < minZeros {
			layerWithFewerZeros = layer
			minZeros = zeros
		}
	}
	answer := layerWithFewerZeros.HowMany(1) * layerWithFewerZeros.HowMany(2)
	println("Part one:", answer)

	println("Part two:")

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			for _, l := range layers {
				switch l[i][j] {
				case 0: //black
					print(" ")
				case 1: //white
					print("█")
				case 2: // transparent
					continue
				}
				break
			}
		}
		println()
	}
}

func extractLayers(input string, width int, height int) []Layer {
	var layers []Layer
	for i := 0; i < len(input); i += width * height {
		layers = append(layers, NewLayer(input[i:i+width*height], width, height))
	}
	return layers
}

type Layer [][]int

func NewLayer(input string, width, height int) Layer {
	layer := Layer(make([][]int, width))
	for i := 0; i < width; i++ {
		layer[i] = make([]int, height)
		for j := 0; j < height; j++ {
			layer[i][j] = int(input[j*width+i] - '0')
		}
	}
	return layer
}

func (l Layer) HowMany(digit int) (count int) {
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l[i]); j++ {
			if l[i][j] == digit {
				count++
			}
		}
	}
	return count
}
