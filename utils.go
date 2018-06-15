package jsonq

import "strings"

func splitBraces(line KeepRequest) []KeepRequest {
	array := []KeepRequest{}
	runes := []rune(string(line))
	count := 0
	firstIndex := 0
	for index, char := range line {
		switch char {
		case '{':
			if count == 0 {
				firstIndex = index + 1
			}
			count++
		case '}':
			count--
			if count == 0 {
				array = append(array, KeepRequest(runes[firstIndex:index]))
			}
		default:
			continue
		}
	}
	if len(array) == 0 {
		array = append(array, line)
	}
	return array
}

func splitComa(line KeepRequest) []KeepRequest {
	array := []KeepRequest{}
	runes := []rune(string(line))
	count := 0
	firstIndex := 0
	for index, char := range line {
		switch char {
		case '{':
			count++
		case '}':
			count--
		case ',':
			if count == 0 {
				array = append(array, KeepRequest(runes[firstIndex:index]))
				firstIndex = index + 1
			}
		default:
			continue
		}
	}
	if firstIndex < len(line) {
		array = append(array, KeepRequest(runes[firstIndex:len(line)]))
	}
	return array
}

func GetKeys(cmd KeepRequest) (stay []KeepRequest, cont []KeepRequest) {
	stay = []KeepRequest{}
	cont = []KeepRequest{}
	if strings.Index(string(cmd), ":") != -1 && strings.Index(string(cmd), ":") < strings.Index(string(cmd), "{") {
		for _, sc := range splitComa(cmd) {
			if strings.Contains(string(sc), ":") {
				cont = append(cont, sc)
			} else {
				stay = append(stay, sc)
			}
		}
	} else {
		for _, sb := range splitBraces(cmd) {
			for _, sc := range splitComa(sb) {
				if strings.Contains(string(sc), ":") {
					cont = append(cont, sc)
				} else {
					stay = append(stay, sc)
				}
			}
		}
	}
	return stay, cont
}
