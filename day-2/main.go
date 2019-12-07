package main

import (
	"log"
)

func main() {
	part1()
	part2()
}

func part2() {
	intcode := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 5, 23, 1, 23, 6, 27, 2, 9, 27, 31, 1, 5, 31, 35, 1, 35, 10, 39, 1, 39, 10, 43, 2, 43, 9, 47, 1, 6, 47, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 59, 10, 63, 1, 9, 63, 67, 1, 9, 67, 71, 2, 71, 6, 75, 1, 5, 75, 79, 1, 5, 79, 83, 1, 9, 83, 87, 2, 87, 10, 91, 2, 10, 91, 95, 1, 95, 9, 99, 2, 99, 9, 103, 2, 10, 103, 107, 2, 9, 107, 111, 1, 111, 5, 115, 1, 115, 2, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}
	mem := make([]int, len(intcode))

	// Reset
	intcode[1], intcode[2] = 12, 2

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {

			copy(mem, intcode)
			mem[1], mem[2] = noun, verb

			position := 0

			for {
				opcode := mem[position]
				if opcode == 99 {
					break
				}

				switch opcode {
				case 1:
					mem[mem[position+3]] = mem[mem[position+1]] + mem[mem[position+2]]
				case 2:
					mem[mem[position+3]] = mem[mem[position+1]] * mem[mem[position+2]]
				default:
					log.Panicln("opcode not found", opcode)
				}

				position += 4
			}

			// Check result
			if mem[0] == 19690720 {
				log.Printf("Result 2: %v", 100*noun+verb)
				return
			}
		}
	}

	log.Printf("Result 2: %v", intcode[0])
}

func part1() {
	var intcode = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 5, 23, 1, 23, 6, 27, 2, 9, 27, 31, 1, 5, 31, 35, 1, 35, 10, 39, 1, 39, 10, 43, 2, 43, 9, 47, 1, 6, 47, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 59, 10, 63, 1, 9, 63, 67, 1, 9, 67, 71, 2, 71, 6, 75, 1, 5, 75, 79, 1, 5, 79, 83, 1, 9, 83, 87, 2, 87, 10, 91, 2, 10, 91, 95, 1, 95, 9, 99, 2, 99, 9, 103, 2, 10, 103, 107, 2, 9, 107, 111, 1, 111, 5, 115, 1, 115, 2, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}

	// Reset
	intcode[1], intcode[2] = 12, 2

	position := 0
	for {
		opcode := intcode[position]
		if opcode == 99 {
			break
		}

		switch opcode {
		case 1:
			intcode[intcode[position+3]] = intcode[intcode[position+1]] + intcode[intcode[position+2]]
		case 2:
			intcode[intcode[position+3]] = intcode[intcode[position+1]] * intcode[intcode[position+2]]
		default:
			log.Panicln("opcode not found", opcode)
		}

		position += 4
	}

	log.Printf("Result 1: %v", intcode[0])
	return
}
