package jsonQuerry

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Object represents JSON object.
//
// Object cannot be used from concurrent goroutines.
// Use per-goroutine parsers or ParserPool instead.
type Object struct {
	kvs           []kv
	keysUnescaped bool
}

func (o *Object) reset() {
	o.kvs = o.kvs[:0]
	o.keysUnescaped = false
}

// String returns string representation for the o.
//
// This function is for debugging purposes only. It isn't optimized for speed.
func (o *Object) String() string {
	o.unescapeKeys()

	// Use bytes.Buffer instead of strings.Builder,
	// so it works on go 1.9 and below.
	var bb bytes.Buffer
	bb.WriteString("{")
	for i, kv := range o.kvs {
		fmt.Fprintf(&bb, "%q:%s", kv.k, kv.v)
		if i != len(o.kvs)-1 {
			bb.WriteString(",")
		}
	}
	bb.WriteString("}")
	return bb.String()
}

func (o *Object) getKV() *kv {
	if cap(o.kvs) > len(o.kvs) {
		o.kvs = o.kvs[:len(o.kvs)+1]
	} else {
		o.kvs = append(o.kvs, kv{})
	}
	return &o.kvs[len(o.kvs)-1]
}

func (o *Object) unescapeKeys() {
	if o.keysUnescaped {
		return
	}
	for i := range o.kvs {
		kv := &o.kvs[i]
		kv.k = unescapeStringBestEffort(kv.k)
	}
	o.keysUnescaped = true
}

// Get returns the value for the given key in the o.
//
// Returns nil if the value for the given key isn't found.
//
// The returned value is valid until Parse is called on the Parser returned o.
func (o *Object) Get(key string) *Value {
	o.unescapeKeys()

	for _, kv := range o.kvs {
		if kv.k == key {
			return kv.v
		}
	}
	return nil
}

// Visit calls f for each item in the o.
//
// f cannot hold key and/or v after returning.
func (o *Object) Visit(f func(key []byte, v *Value)) {
	if o == nil {
		return
	}

	o.unescapeKeys()

	for _, kv := range o.kvs {
		f(s2b(kv.k), kv.v)
	}
}

// Value represents any JSON value.
//
// Call Type in order to determine the actual type of the JSON value.
//
// Value cannot be used from concurrent goroutines.
// Use per-goroutine parsers or ParserPool instead.
type Value struct {
	o Object
	a []*Value
	s string
	n float64
	t Type
}

func (v *Value) reset() {
	v.o.reset()
	v.a = v.a[:0]
	v.s = ""
	v.n = 0
	v.t = TypeNull
}

// String returns string representation of the v.
//
// The function is for debugging purposes only. It isn't optimized for speed.
//
// Don't confuse this function with StringBytes, which must be called
// for obtaining the underlying JSON string for the v.
func (v *Value) String() string {
	switch v.Type() {
	case TypeObject:
		return v.o.String()
	case TypeArray:
		// Use bytes.Buffer instead of strings.Builder,
		// so it works on go 1.9 and below.
		var bb bytes.Buffer
		bb.WriteString("[")
		for i, vv := range v.a {
			fmt.Fprintf(&bb, "%s", vv)
			if i != len(v.a)-1 {
				bb.WriteString(",")
			}
		}
		bb.WriteString("]")
		return bb.String()
	case TypeString:
		return fmt.Sprintf("%q", v.s)
	case TypeNumber:
		if float64(int(v.n)) == v.n {
			return fmt.Sprintf("%d", int(v.n))
		}
		return fmt.Sprintf("%f", v.n)
	case TypeTrue:
		return "true"
	case TypeFalse:
		return "false"
	case TypeNull:
		return "null"
	default:
		panic(fmt.Errorf("BUG: unknown Value type: %d", v.Type()))
	}
}

// Type returns the type of the v.
func (v *Value) Type() Type {
	switch v.t {
	case typeRawString:
		v.s = unescapeStringBestEffort(v.s)
		v.t = TypeString
	case typeRawNumber:
		f, err := strconv.ParseFloat(v.s, 64)
		if err != nil {
			f = 0
		}
		v.n = f
		v.t = TypeNumber
	}
	return v.t
}

// Exists returns true if the field exists for the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
// func (v *Value) Exists(keys ...string) bool {
// 	v = v.Get(keys...)
// 	return v != nil
// }

// Get returns value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// nil is returned for non-existing keys path.
//
// The returned value is valid until Parse is called on the Parser returned v.
func (v *Value) Get(keys ...string) *Value {
	if v == nil {
		return nil
	}
	for _, key := range keys {
		switch v.t {
		case TypeObject:
			v = v.o.Get(key)
			if v == nil {
				return nil
			}
		case TypeArray:
			n, err := strconv.Atoi(key)
			if err != nil || n < 0 || n >= len(v.a) {
				return nil
			}
			v = v.a[n]
		default:
			return nil
		}
	}
	return v
}

// Object returns the underlying JSON object for the v.
//
// The returned object is valid until Parse is called on the Parser returned v.
//
// Use GetObject if you don't need error handling.
func (v *Value) Object() (*Object, error) {
	if v.Type() != TypeObject {
		return nil, fmt.Errorf("value doesn't contain object; it contains %s", v.Type())
	}
	return &v.o, nil
}

// Array returns the underlying JSON array for the v.
//
// The returned array is valid until Parse is called on the Parser returned v.
//
// Use GetArray if you don't need error handling.
func (v *Value) Array() ([]*Value, error) {
	if v.Type() != TypeArray {
		return nil, fmt.Errorf("value doesn't contain array; it contains %s", v.Type())
	}
	return v.a, nil
}

// StringBytes returns the underlying JSON string for the v.
//
// The returned string is valid until Parse is called on the Parser returned v.
//
// Use GetStringBytes if you don't need error handling.
func (v *Value) StringBytes() ([]byte, error) {
	if v.Type() != TypeString {
		return nil, fmt.Errorf("value doesn't contain string; it contains %s", v.Type())
	}
	return s2b(v.s), nil
}

// Float64 returns the underlying JSON number for the v.
//
// Use GetFloat64 if you don't need error handling.
func (v *Value) Float64() (float64, error) {
	if v.Type() != TypeNumber {
		return 0, fmt.Errorf("value doesn't contain number; it contains %s", v.Type())
	}
	return v.n, nil
}

// Int returns the underlying JSON int for the v.
//
// Use GetInt if you don't need error handling.
func (v *Value) Int() (int, error) {
	f, err := v.Float64()
	return int(f), err
}

// Bool returns the underlying JSON bool for the v.
//
// Use GetBool if you don't need error handling.
func (v *Value) Bool() (bool, error) {
	switch v.Type() {
	case TypeTrue:
		return true, nil
	case TypeFalse:
		return false, nil
	default:
		return false, fmt.Errorf("value doesn't contain bool; it contains %s", v.Type())
	}
}

// Search return an Array of interface values by the given keys path
func (v *Value) Search(keys ...string) ([]interface{}, error) {
	var rValues []interface{}
	switch v.Type() {
	case TypeArray:
		pValue, err := v.Array()
		if err != nil {
			return nil, err
		}
		for _, uValue := range pValue {
			nValue, err := uValue.Search(keys...)
			if err != nil {
				return nil, err
			}
			rValues = append(rValues, nValue...)
		}
	case TypeObject:
		pValue, err := v.Object()
		if err != nil {
			return nil, err
		}
		nValue, err := pValue.Get(keys[0]).Search(keys[1:]...)
		if err != nil {
			return nil, err
		}
		rValues = append(rValues, nValue...)
	case TypeString:
		rValues = append(rValues, string(v.String()))
	case TypeNumber:
		pValue, err := v.Float64()
		if err != nil {
			return nil, err
		}
		rValues = append(rValues, float64(pValue))
	case TypeFalse:
		rValues = append(rValues, false)
	case TypeTrue:
		rValues = append(rValues, true)
	default:
		return nil, fmt.Errorf("Type not recognized")
	}
	return rValues, nil
}

// {description, produit:{truc,machin,}}

func (v *Value) Keep(request string) (interface{}, error) {
	strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, request)
	switch v.Type() {
	case TypeArray:
		pValue, err := v.Array()
		if err != nil {
			return nil, err
		}
		rValues := []interface{}{}
		for _, uValue := range pValue {
			nValue, err := uValue.Keep(request)
			if err != nil {
				return nil, err
			}
			rValues = append(rValues, nValue)
		}
		return rValues, nil
	case TypeObject:
		pValue, err := v.Object()
		if err != nil {
			return nil, err
		}
		rValues := map[string]interface{}{}
		stays, conts := getKeys(request)
		for _, stay := range stays {
			rValues[stay], err = pValue.Get(stay).Keep("")
			if err != nil {
				return nil, err
			}
		}
		for _, cont := range conts {
			key := strings.Split(cont, ":")[0]
			val := splitBraces(cont)
			rValues[key], err = pValue.Get(key).Keep(val[0])
			if err != nil {
				return nil, err
			}
		}
		return rValues, nil
	case TypeString:
		return v.String(), nil
	case TypeNumber:
		return v.Float64()
	case TypeFalse:
		return false, nil
	case TypeTrue:
		return true, nil
	default:
		return nil, fmt.Errorf("Type not recognized")
	}
	return nil, nil
}

func splitBraces(line string) []string {
	array := []string{}
	runes := []rune(line)
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
				array = append(array, string(runes[firstIndex:index]))
			}
		default:
			continue
		}
	}
	return array
}

func splitComa(line string) []string {
	array := []string{}
	runes := []rune(line)
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

func getKeys(cmd string) (stay []string, cont []string) {
	strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, cmd)
	stay = []string{}
	cont = []string{}
	for _, sb := range splitBraces(cmd) {
		for _, sc := range splitComa(sb) {
			if strings.Contains(sc, ":") {
				cont = append(cont, sc)
			} else {
				stay = append(stay, sc)
			}
		}
	}
	return stay, cont
}

var (
	valueTrue   = &Value{t: TypeTrue}
	valueFalse  = &Value{t: TypeFalse}
	valueNull   = &Value{t: TypeNull}
	emptyObject = &Value{t: TypeObject}
	emptyArray  = &Value{t: TypeArray}
)
