package main

import (
	"sort"
	"strconv"
	"strings"

	. "advent-of-code-2021/utils"
)

const example = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func main() {
	DownloadDayInput(2021, 8, false)
	input := SplitNewLine(ReadInputFileRelative())
	//input = SplitNewLine(example)

	count := 0
	for _, l := range input {
		for _, s := range strings.Split(strings.Split(l, "|")[1], " ") {
			if s != "" {
				switch len(s) {
				case 2, 3, 4, 7:
					count++
				}
			}
		}
	}
	println("Part 1:", count)

	count = 0
	for _, l := range input {
		count += decode(l)
	}
	println("Part 2:", count)
}

func decode(l string) int {
	strs := make([]string, 10)
	for _, p := range strings.Split(l, " ") {
		if p == "" || p == "|" {
			continue
		}
		var i int
		p = sortStr(p)
		switch len(p) {
		case 2:
			i = 1
		case 3:
			i = 7
		case 4:
			i = 4
		case 7:
			i = 8
		default:
			continue
		}
		strs[i] = p
	}
	for _, p := range strings.Split(l, " ") {
		if p == "" || p == "|" {
			continue
		}
		one := strs[1]
		four := strs[4]
		seven := strs[7]

		switch len(p) {
		case 2, 3, 4, 7:
			continue
		case 5:
			// 2,3,5
			p = sortStr(p)
			is2, is3, is5 := true, true, true
			if one != "" {
				if nbMatch(one, p) == 2 {
					strs[3] = p
					continue
				}
				if nbMatch(one, p) == 1 {
					is3 = false
				}
			}
			if four != "" {
				if nbMatch(four, p) == 2 {
					strs[2] = p
					continue
				}
				if nbMatch(four, p) == 3 {
					is2 = false
				}
			}
			if seven != "" {
				if nbMatch(seven, p) == 3 {
					strs[3] = p
					continue
				}
				if nbMatch(four, p) == 2 {
					is3 = false
				}
			}
			if four != "" && seven != "" {
				if nbMatch(merge(four, seven), p) == 3 {
					strs[2] = p
					continue
				}
				if nbMatch(merge(four, seven), p) == 4 {
					is2 = false
				}
			}
			if (is2 && is5) || (is2 && is3) || (is3 && is5) {
				panic(p)
			}
			if is2 {
				strs[2] = p
			}
			if is3 {
				strs[3] = p
			}
			if is5 {
				strs[5] = p
			}
		case 6:
			// 0,6,9
			p = sortStr(p)
			is0, is6, is9 := true, true, true
			if four != "" && seven != "" {
				if p == merge(four, seven) {
					strs[9] = p
					continue
				}
			}
			if one != "" {
				if nbMatch(one, p) == 1 {
					strs[6] = p
					continue
				}
				if nbMatch(one, p) == 2 {
					// 0 or 9
					is6 = false
				}
			}
			if four != "" {
				if nbMatch(four, p) == 4 {
					strs[9] = p
					continue
				}
				if nbMatch(four, p) == 3 {
					is9 = false
				}
			}
			if seven != "" {
				if nbMatch(seven, p) == 2 {
					strs[6] = p
					continue
				}
				if nbMatch(seven, p) == 3 {
					is6 = false
				}
			}
			if (is0 && is6) || (is0 && is9) || (is9 && is6) {
				panic(p)
			}
			if is0 {
				strs[0] = p
			}
			if is6 {
				strs[6] = p
			}
			if is9 {
				strs[9] = p
			}
		default:
			panic(p)
		}
	}

	ints := map[string]int{}
	for i, s := range strs {
		if s == "" {
			panic(s)
		}
		ints[s] = i
	}

	parts := strings.Split(l, "|")
	var decoded string
	for _, p := range strings.Split(parts[1], " ") {
		if p == "" {
			continue
		}
		i, ok := ints[sortStr(p)]
		if !ok {
			i, _ = ints[""]
		}
		decoded += strconv.Itoa(i)
	}
	i, _ := strconv.Atoi(decoded)
	return i
}

func sortStr(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func nbMatch(s, r string) int {
	count := 0
	for _, k := range s {
		if strings.ContainsRune(r, k) {
			count++
		}
	}
	return count
}

func merge(s, r string) string {
	for _, c := range r {
		if !strings.ContainsRune(s, c) {
			s += string(c)
		}
	}
	return sortStr(s)
}
