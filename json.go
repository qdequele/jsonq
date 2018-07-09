package jsonq

import (
	"bytes"
	"fmt"
)

func (v Value) check(filter Filter) bool {
	switch v.Type() {
	case TypeString:
		return filter.check(v.s)
	case TypeNumber:
		return filter.check(v.n)
	case TypeTrue:
		return filter.check(true)
	case TypeFalse:
		return filter.check(false)
	case TypeNull:
		return filter.check(nil)
	default:
		return false
	}
}

// Search return an Array of interface values by the given keys path
func (v Value) Search(keys ...string) ([]interface{}, error) {
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
		next := pValue.Get(keys[0])
		if next == nil {
			return nil, fmt.Errorf("key not found : %s", keys[0])
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

// Check
func (v Value) Check(request Query) error {
	switch v.Type() {
	case TypeArray:
		pValue, err := v.Array()
		if err != nil {
			return err
		}
		for _, uValue := range pValue {
			err := uValue.Check(request)
			if err == nil {
				return nil
			}
		}
		return fmt.Errorf("No element found")
	case TypeObject:
		pValue, err := v.Object()
		if err != nil {
			return err
		}
		if request.keepFilters {
			for _, filter := range request.filters {
				nValue := pValue.Get(filter.key)
				if nValue != nil && nValue.check(*filter) == false {
					return fmt.Errorf("")
				}
			}
			for name, next := range request.next {
				nValue := pValue.Get(name)
				if next != nil {
					err := nValue.Check(Query(*next))
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	case TypeString, TypeNumber, TypeFalse, TypeTrue, TypeNull:
		return nil
	default:
		return fmt.Errorf("Type not recognized")
	}
}

func (v Value) Keep(request Query) (string, error) {
	w := bytes.Buffer{}
	switch v.Type() {
	case TypeArray:
		pValue, err := v.Array()
		if err != nil {
			return "", err
		}
		w.WriteRune('[')
		for index, uValue := range pValue {
			nValue, err := uValue.Keep(request)
			if err != nil {
				return "", err
			}
			if len(nValue) > 0 {
				w.WriteString(nValue)
				if index < len(pValue)-1 {
					w.WriteRune(',')
				}
			}
		}
		w.WriteRune(']')
		return w.String(), nil
	case TypeObject:
		pValue, err := v.Object()
		if err != nil {
			return "", err
		}
		for _, filter := range request.filters {
			if nValue := pValue.Get(filter.key); nValue != nil {
				if nValue.check(*filter) == false {
					return "", nil
				}
			}
		}
		w.WriteRune('{')
		i := 0
		for _, retrieve := range request.retrieve {
			i++
			w.WriteRune('"')
			w.WriteString(retrieve)
			w.WriteRune('"')
			w.WriteRune(':')
			w.WriteString(pValue.Get(retrieve).String())
			if i < len(request.next)+len(request.retrieve) {
				w.WriteRune(',')
			}
		}
		for name, next := range request.next {
			i++
			nValue, err := pValue.Get(name).Keep(Query(*next))
			if err != nil {
				return "", err
			}
			w.WriteRune('"')
			w.WriteString(name)
			w.WriteRune('"')
			w.WriteRune(':')
			w.WriteString(nValue)
			if i < len(request.next)+len(request.retrieve) {
				w.WriteRune(',')
			}
		}
		w.WriteRune('}')
		return w.String(), nil
	case TypeString:
		return fmt.Sprintf("%q", v.s), nil
	case TypeNumber:
		if float64(int(v.n)) == v.n {
			return fmt.Sprintf("%d", int(v.n)), nil
		}
		return fmt.Sprintf("%f", v.n), nil
	case TypeFalse:
		return "false", nil
	case TypeTrue:
		return "true", nil
	case TypeNull:
		return "null", nil
	default:
		return "", fmt.Errorf("Type not recognized")
	}
}

func (v Value) Retrieve(request Query) (string, error) {
	w := bytes.Buffer{}
	switch v.Type() {
	case TypeArray:
		pValue, err := v.Array()
		if err != nil {
			return "", err
		}
		w.WriteRune('[')
		for index, uValue := range pValue {
			nValue, err := uValue.Keep(request)
			if err != nil {
				return "", err
			}
			if len(nValue) > 0 {
				w.WriteString(nValue)
				if index < len(pValue)-1 {
					w.WriteRune(',')
				}
			}
		}
		w.WriteRune(']')
		return w.String(), nil
	case TypeObject:
		pValue, err := v.Object()
		if err != nil {
			return "", err
		}
		w.WriteRune('{')
		i := 0
		for _, retrieve := range request.retrieve {
			i++
			val := pValue.Get(retrieve)
			if val != nil {
				w.WriteRune('"')
				w.WriteString(retrieve)
				w.WriteRune('"')
				w.WriteRune(':')
				w.WriteString(val.String())
				if i < len(request.next)+len(request.retrieve) {
					w.WriteRune(',')
				}
			}
		}
		for name, next := range request.next {
			i++
			nValue, err := pValue.Get(name).Keep(Query(*next))
			if err != nil {
				return "", err
			}
			w.WriteRune('"')
			w.WriteString(name)
			w.WriteRune('"')
			w.WriteRune(':')
			w.WriteString(nValue)
			if i < len(request.next)+len(request.retrieve) {
				w.WriteRune(',')
			}
		}
		w.WriteRune('}')
		return w.String(), nil
	case TypeString:
		return fmt.Sprintf("%q", v.s), nil
	case TypeNumber:
		if float64(int(v.n)) == v.n {
			return fmt.Sprintf("%d", int(v.n)), nil
		}
		return fmt.Sprintf("%f", v.n), nil
	case TypeFalse:
		return "false", nil
	case TypeTrue:
		return "true", nil
	case TypeNull:
		return "null", nil
	default:
		return "", fmt.Errorf("Type not recognized")
	}
}
