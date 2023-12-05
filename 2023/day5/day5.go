package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, _ := os.Open("day5.input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	type Range struct {
		start, end, destination int
	}

	seeds := []int{}
	seed_to_soil := []Range{}
	soil_to_fertilizer := []Range{}
	fertilizer_to_water := []Range{}
	water_to_light := []Range{}
	light_to_temperature := []Range{}
	temperature_to_humidity := []Range{}
	humidity_to_location := []Range{}

	maps := [][]Range{seed_to_soil, soil_to_fertilizer, fertilizer_to_water, water_to_light, light_to_temperature, temperature_to_humidity, humidity_to_location}

	maps_idx := -1
	for scanner.Scan() {
		text := scanner.Text()
		split_text := strings.Split(text, " ")
		if text == "" {
			continue
		}
		if split_text[0] == "seeds:" {
			for _, seed := range split_text[1:] {
				seed_int, _ := strconv.Atoi(seed)
				seeds = append(seeds, seed_int)
			}
		} else {
			if len(split_text) < 3 {
				maps_idx++
				continue
			}
			destination, _ := strconv.Atoi(split_text[0])
			source, _ := strconv.Atoi(split_text[1])
			distance, _ := strconv.Atoi(split_text[2])
			maps[maps_idx] = append(maps[maps_idx], Range{source, source + distance - 1, destination})
		}
	}

	min_location := math.MaxInt64
	for _, seed := range seeds {
		next := seed
		for m_idx, m := range maps {
			for _, v := range m {
				if next >= v.start && next <= v.end {
					next = v.destination + (next - v.start)
					break
				}
			}
			if m_idx == len(maps)-1 {
				min_location = Min(next, min_location)

			}
		}
	}
	fmt.Println(min_location)

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func part2() {
	file, _ := os.Open("day5.input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	type Range struct {
		start, end, destination int
	}

	seeds := []Range{}
	seed_to_soil := []Range{}
	soil_to_fertilizer := []Range{}
	fertilizer_to_water := []Range{}
	water_to_light := []Range{}
	light_to_temperature := []Range{}
	temperature_to_humidity := []Range{}
	humidity_to_location := []Range{}

	maps := [][]Range{seed_to_soil, soil_to_fertilizer, fertilizer_to_water, water_to_light, light_to_temperature, temperature_to_humidity, humidity_to_location}

	maps_idx := -1
	for scanner.Scan() {
		text := scanner.Text()
		split_text := strings.Split(text, " ")
		if text == "" {
			continue
		}
		if split_text[0] == "seeds:" {
			for i := 1; i < len(split_text); i += 2 {
				start, _ := strconv.Atoi(split_text[i])
				end, _ := strconv.Atoi(split_text[i+1])
				seeds = append(seeds, Range{start, start + end, 0})
			}
		} else {
			if len(split_text) < 3 {
				maps_idx++
				continue
			}
			destination, _ := strconv.Atoi(split_text[0])
			source, _ := strconv.Atoi(split_text[1])
			distance, _ := strconv.Atoi(split_text[2])
			maps[maps_idx] = append(maps[maps_idx], Range{source, source + distance - 1, destination})
		}
	}

	min_location := math.MaxInt64
	for _, seed := range seeds {
		for i := seed.start; i < seed.end; i++ {
			next := i
			for m_idx, m := range maps {
				for _, v := range m {
					if next >= v.start && next <= v.end {
						next = v.destination + (next - v.start)
						break
					}
				}
				if m_idx == len(maps)-1 {
					min_location = Min(next, min_location)

				}
			}
		}
	}

	fmt.Println(min_location)
}

func main() {
	part1()
	part2()
}
