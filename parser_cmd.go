package jsonq

import (
	"fmt"
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
	contain     Operation = ":"
	notContain  Operation = "!:"
	like        Operation = "::"
	notLike     Operation = "!::"
)

var cmdRegex = regexp.MustCompile(`^([a-z_]+)?(?:\(([^{\}\)\(]*)\))?{(.*)}$`)
var filterRegex = regexp.MustCompile(`(?:([a-zA-Z_-]+)\s*([><!:=]+)\s*((?:[^&\(\)\{}\s\")]+|(?:\"[^&\(\)\{}]*\")))\s*)+`)

// Operation is common possible operations in filters (==, ===, !=, !==, >, <, >=, <=, :).
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
	case contain:
		return checkContain(base, compared)
	case notContain:
		return checkNotContain(base, compared)
	case like:
		return checkLike(base, compared)
	case notLike:
		return checkNotLike(base, compared)
	default:
		return false
	}
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
	case ":":
		return contain
	case "!:":
		return notContain
	case "::":
		return like
	case "!::":
		return notLike
	default:
		return "error"
	}
}

func checkEq(base, compared interface{}) bool {
	switch v := base.(type) {
	case bool:
		if comp, ok := compared.(bool); ok == true {
			return comp == v
		}
		return false
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
	case bool:
		if comp, ok := compared.(bool); ok == true {
			return comp == v
		}
		return false
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
	case bool:
		if comp, ok := compared.(bool); ok == true {
			return comp != v
		}
		return false
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
	case bool:
		if comp, ok := compared.(bool); ok == true {
			return comp != v
		}
		return false
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
	case bool, nil:
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
	case bool, nil:
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
	case bool, nil:
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
	case bool, nil:
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

// In this case we check if the compare string is contained int the base string
func checkContain(base, compared interface{}) bool {
	if b, ok := base.(string); ok == true {
		if c, ok := compared.(string); ok == true {
			b = strings.ToLower(b)
			b = strings.TrimLeft(b, `"`)
			b = strings.TrimRight(b, `"`)
			c = strings.ToLower(c)
			return strings.Contains(c, b)
		}
	}
	return false
}

func checkNotContain(base, compared interface{}) bool {
	if b, ok := base.(string); ok == true {
		if c, ok := compared.(string); ok == true {
			b = strings.ToLower(b)
			b = strings.TrimLeft(b, `"`)
			b = strings.TrimRight(b, `"`)
			c = strings.ToLower(c)
			return !strings.Contains(c, b)
		}
	}
	return false
}

// In this function base was the json value, compared the string used for the regex. Both should be strings
func checkLike(base, compared interface{}) bool {
	if b, ok := base.(string); ok == true {
		if c, ok := compared.(string); ok == true {
			b = strings.ToLower(b)
			b = strings.TrimLeft(b, `"`)
			b = strings.TrimRight(b, `"`)
			c = strings.ToLower(c)
			if ok, err := regexp.MatchString(b, c); err == nil {
				return ok
			}
		}
	}
	return false
}

func checkNotLike(base, compared interface{}) bool {
	if b, ok := base.(string); ok == true {
		if c, ok := compared.(string); ok == true {
			b = strings.ToLower(b)
			b = strings.TrimLeft(b, `"`)
			b = strings.TrimRight(b, `"`)
			c = strings.ToLower(c)
			if ok, err := regexp.MatchString(b, c); err == nil {
				return !ok
			}
		}
	}
	return false
}

//Filter is the type used for describe a operation of filtering
type Filter struct {
	key string
	op  Operation
	val interface{}
}

func (f Filter) eq(other Filter) bool {
	bkey := f.key == other.key
	bop := f.op == other.op
	bval := fmt.Sprintln(f.val) == fmt.Sprintln(other.val)
	return bkey && bop && bval
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

// Query is a description of a Query in a graphql like request
type Query struct {
	filters     []*Filter
	next        map[string]*Query
	retrieve    []string
	keepFilters bool
}

func (q Query) eq(other Query) bool {
	for index, filter := range q.filters {
		if other.filters[index] == nil || !filter.eq(*other.filters[index]) {
			return false
		}
	}
	for key, query := range q.next {
		if other.next[key] == nil || !query.eq(*other.next[key]) {
			return false
		}
	}
	for index, retrieve := range q.retrieve {
		if retrieve != other.retrieve[index] {
			return false
		}
	}
	return true
}

func newQuery() Query {
	return Query{
		[]*Filter{},
		map[string]*Query{},
		[]string{},
		false,
	}
}

func (l Query) print(Query int) {
	fmt.Printf("%s Filters :\n", strings.Repeat("\t", Query))
	for _, filter := range l.filters {
		fmt.Printf("%s - %s %s %q \n", strings.Repeat("\t", Query), (*filter).key, filter.op, filter.val)
	}
	fmt.Printf("%s Retrieve :\n", strings.Repeat("\t", Query))
	for _, retrieve := range l.retrieve {
		if len(retrieve) > 0 {
			fmt.Printf("%s - %s\n", strings.Repeat("\t", Query), retrieve)
		}
	}
	fmt.Printf("%s Next :\n", strings.Repeat("\t", Query))
	for _, next := range l.next {
		next.print(Query + 1)
	}
}

// Print will recursively show the content of Querys.
func (l Query) Print() {
	l.print(0)
}

func parseQuery(cmd string) (Query *Query, QueryName string, err error) {
	matches := cmdRegex.FindStringSubmatch(cmd)
	lvl := newQuery()
	if len(matches) == 0 {
		return nil, "", fmt.Errorf("mal formated")
	}
	if len(matches) > 2 && len(matches[2]) > 0 {
		for _, filter := range newFilter(matches[2]) {
			if filter != nil {
				lvl.filters = append(lvl.filters, filter)
			}
		}
		if len(lvl.filters) > 0 {
			lvl.keepFilters = true
		}
	}
	if len(matches) > 3 && len(matches[3]) > 0 {
		for _, attr := range splitComa(matches[3]) {
			if strings.ContainsAny(attr, "(){}") {
				newQuery, QueryName, _ := parseQuery(attr)
				if newQuery.keepFilters == true {
					lvl.keepFilters = true
				}
				lvl.next[QueryName] = newQuery
			} else {
				lvl.retrieve = append(lvl.retrieve, attr)
			}
		}
	}
	return &lvl, matches[1], nil
}

// ParseQuery create a easy traversable structure from a graphql like query.
func ParseQuery(cmd string) (parser *Query, err error) {
	parser, _, err = parseQuery(cmd)
	return parser, err
}

// MustParseQuery is parseQuery without error return. You should be sure of your query syntax !
func MustParseQuery(cmd string) (parser *Query) {
	parser, _, err := parseQuery(cmd)
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
		case ' ', '\n', '\t':
			if count == 0 {
				firstIndex++
			}
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
