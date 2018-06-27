package jsonq

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	eq          Operation = "==="
	partialEq   Operation = "=="
	diff        Operation = "!=="
	partialDiff Operation = "!="
	sup         Operation = ">"
	supEq       Operation = ">="
	inf         Operation = "<"
	infEq       Operation = "<="
)

var cmdRegex = regexp.MustCompile(`^([a-z_]+)?(?:\(([^{\}\)\(]*)\))?{(.*)}$`)
var filterRegex = regexp.MustCompile(`(?:([a-z_]+)\s*([><!=]+)\s*((?:[^&\(\)\{}\s\")]+|(?:\"[^&\(\)\{}]*\")))\s*)+`)
var comaRegex = regexp.MustCompile(`((?:[^,]+\(.+?\){.+?})+?|(?:[^,]+{.+?})+?|(?:[^,]+))`)

type Operation string

func (o Operation) check(base, compared interface{}) bool {
	switch o {
	case eq:
		return checkEq(base, compared)
	case partialEq:
		return checkPartialEq(base, compared)
	case diff:
		return checkDiff(base, compared)
	case partialDiff:
		return checkPartialDiff(base, compared)
	case sup:
		return checkSup(base, compared)
	case supEq:
		return checkSupEq(base, compared)
	case inf:
		return checkInf(base, compared)
	case infEq:
		return checkInfEq(base, compared)
	default:
		return false
	}
	return true
}

func findOperation(line string) Operation {
	switch line {
	case "===":
		return eq
	case "==":
		return partialEq
	case "!==":
		return diff
	case "!=":
		return partialDiff
	case ">":
		return sup
	case ">=":
		return supEq
	case "<":
		return inf
	case "<=":
		return infEq
	default:
		return "error"
	}
}

func checkEq(base, compared interface{}) bool {
	switch v := base.(type) {
	case int64:
		if comp, ok := compared.(int64); ok == true && comp == v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp == float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp == v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) == v {
			return true
		}
		return false
	case string:
		if comp, ok := compared.(string); ok == true && comp == v {
			return true
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkEq(key, compared) == false {
				return false
			}
		}
		return true
	}
	return false
}

func checkPartialEq(base, compared interface{}) bool {
	switch v := base.(type) {
	case int64:
		if comp, ok := compared.(int64); ok == true && comp == v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp == float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp == v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) == v {
			return true
		}
		return false
	case string:
		if comp, ok := compared.(string); ok == true {
			return strings.Contains(comp, strings.Replace(v, "\"", "", -1))
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkEq(key, compared) == true {
				return true
			}
		}
	}
	return false
}

func checkDiff(base, compared interface{}) bool {
	switch v := base.(type) {
	case int64:
		if comp, ok := compared.(int64); ok == true && comp != v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp != float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp != v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) != v {
			return true
		}
		return false
	case string:
		if comp, ok := compared.(string); ok == true && comp != v {
			return true
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkDiff(key, compared) != false {
				return false
			}
		}
		return true
	}
	return false
}

func checkPartialDiff(base, compared interface{}) bool {
	switch v := base.(type) {
	case int64:
		if comp, ok := compared.(int64); ok == true && comp != v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp != float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp != v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) != v {
			return true
		}
		return false
	case string:
		if comp, ok := compared.(string); ok == true && comp != v {
			return true
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkDiff(key, compared) != true {
				return true
			}
		}
	}
	return false
}

func checkSup(base, compared interface{}) bool {
	switch v := base.(type) {
	case nil:
		return false
	case int64:
		if comp, ok := compared.(int64); ok == true && comp > v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp > float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp > v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) > v {
			return true
		}
		return false
	case string:
		fmt.Printf("%q -%q\n", reflect.TypeOf(base), reflect.TypeOf(compared))
		if comp, ok := compared.(string); ok == true && comp > v {
			return true
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkSup(key, compared) == false {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func checkSupEq(base, compared interface{}) bool {
	switch v := base.(type) {
	case nil:
		return false
	case int64:
		if comp, ok := compared.(int64); ok == true && comp >= v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp >= float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp >= v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) >= v {
			return true
		}
		return false
	case string:
		if comp, ok := compared.(string); ok == true && comp > v {
			return true
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkSupEq(key, compared) == true {
				return true
			}
		}
		return false
	}
	return false
}

func checkInf(base, compared interface{}) bool {
	switch v := base.(type) {
	case nil:
		return false
	case int64:
		if comp, ok := compared.(int64); ok == true && comp < v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp < float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp < v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) < v {
			return true
		}
		return false
	case string:
		if comp, ok := compared.(string); ok == true && comp < v {
			return true
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkSup(key, compared) == false {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func checkInfEq(base, compared interface{}) bool {
	switch v := base.(type) {
	case nil:
		return false
	case int64:
		if comp, ok := compared.(int64); ok == true && comp <= v {
			return true
		} else if comp, ok := compared.(float64); ok == true && comp <= float64(v) {
			return true
		}
		return false
	case float64:
		if comp, ok := compared.(float64); ok == true && comp <= v {
			return true
		} else if comp, ok := compared.(int64); ok == true && float64(comp) <= v {
			return true
		}
		return false
	case string:
		if comp, ok := compared.(string); ok == true && v < comp {
			return true
		}
		return false
	case []interface{}:
		for _, key := range v {
			if checkInfEq(key, compared) == true {
				return true
			}
		}
		return false
	}
	return false
}

type Filter struct {
	key string
	op  Operation
	val interface{}
}

func (f Filter) check(compareTo interface{}) bool {
	return f.op.check(f.val, compareTo)
}

func typed(v string) interface{} {
	switch v {
	case "true":
		return true
	case "false":
		return false
	case "null":
		return nil
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		return i
	}
	f, err := strconv.ParseFloat(v, 64)
	if err == nil {
		return f
	}
	return v
}

func newFilter(cmd string) []*Filter {
	filters := make([]*Filter, 0, len(strings.Split(cmd, "&&")))
	for _, match := range filterRegex.FindAllStringSubmatch(cmd, -1) {
		if len(match[1]) > 0 && len(match[2]) > 0 && len(match[3]) > 0 {
			filters = append(filters, &Filter{
				match[1],
				findOperation(match[2]),
				typed(match[3]),
			})
		}
	}
	return filters
}

type Level struct {
	filters  []*Filter
	next     map[string]*Level
	retrieve []string
}

func newLevel() Level {
	return Level{
		make([]*Filter, 0, 10),
		make(map[string]*Level),
		make([]string, 0, 100),
	}
}

func (l Level) print(level int) {
	fmt.Printf("%s Filters :\n", strings.Repeat("\t", level))
	for _, filter := range l.filters {
		fmt.Printf("%s - %s %s %q \n", strings.Repeat("\t", level), (*filter).key, filter.op, filter.val)
	}
	fmt.Printf("%s Retrieve :\n", strings.Repeat("\t", level))
	for _, retrieve := range l.retrieve {
		if len(retrieve) > 0 {
			fmt.Printf("%s - %s\n", strings.Repeat("\t", level), retrieve)
		}
	}
	fmt.Printf("%s Next :\n", strings.Repeat("\t", level))
	for _, next := range l.next {
		next.print(level + 1)
	}
}

func (l Level) Print() {
	l.print(0)
}

func parseCMD(cmd string) (level *Level, levelName string, err error) {
	matches := cmdRegex.FindStringSubmatch(cmd)
	lvl := newLevel()
	if len(matches[2]) > 0 {
		for _, filter := range newFilter(matches[2]) {
			if filter != nil {
				lvl.filters = append(lvl.filters, filter)
			}
		}
	}
	if len(matches[3]) > 0 {
		for _, attr := range splitComa(matches[3]) {
			if strings.ContainsAny(attr, "(){}") {
				newLevel, levelName, _ := parseCMD(attr)
				lvl.next[levelName] = newLevel
			} else {
				lvl.retrieve = append(lvl.retrieve, attr)
			}
		}
	}
	return &lvl, matches[1], nil
}

func ParseCMD(cmd string) (parser *Level, err error) {
	parser, _, err = parseCMD(cmd)
	return parser, err
}

func MustParseCMD(cmd string) (parser *Level) {
	parser, _, err := parseCMD(cmd)
	if err != nil {
		panic(err)
	}
	return parser
}

func splitComa(line string) []string {
	array := []string{}
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
				array = append(array, string(runes[firstIndex:index]))
				firstIndex = index + 1
			}
		default:
			continue
		}
	}
	if firstIndex < len(line) {
		array = append(array, string(runes[firstIndex:len(line)]))
	}
	return array
}
