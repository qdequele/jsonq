package jsonq

import (
	"testing"
)

func TestMustParseQuery(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name       string
		args       args
		wantParser *Query
	}{
		{"retrieve only", args{"{a,b,c}"}, &Query{[]*Filter{}, map[string]*Query{}, []string{"a", "b", "c"}, false}},
		{"filter only", args{"(a : 1){}"}, &Query{[]*Filter{&Filter{"a", ":", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a:1){}"}, &Query{[]*Filter{&Filter{"a", ":", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a :: 1){}"}, &Query{[]*Filter{&Filter{"a", "::", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a::1){}"}, &Query{[]*Filter{&Filter{"a", "::", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a>1){}"}, &Query{[]*Filter{&Filter{"a", ">", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a > 1){}"}, &Query{[]*Filter{&Filter{"a", ">", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a<1){}"}, &Query{[]*Filter{&Filter{"a", "<", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a < 1){}"}, &Query{[]*Filter{&Filter{"a", "<", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a==1){}"}, &Query{[]*Filter{&Filter{"a", "==", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a == 1){}"}, &Query{[]*Filter{&Filter{"a", "==", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a===1){}"}, &Query{[]*Filter{&Filter{"a", "===", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a === 1){}"}, &Query{[]*Filter{&Filter{"a", "===", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a == 1 && b > 0){}"}, &Query{[]*Filter{&Filter{"a", "==", 1}, &Filter{"b", ">", 0}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a == 1 && b > 0){a,b,c{x,y,z}}"}, &Query{[]*Filter{&Filter{"a", "==", 1}, &Filter{"b", ">", 0}}, map[string]*Query{"c": &Query{[]*Filter{}, map[string]*Query{}, []string{"x", "y", "z"}, false}}, []string{"a", "b"}, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParser := MustParseQuery(tt.args.cmd); !gotParser.eq(*tt.wantParser) {
				t.Errorf("MustParseQuery() = %v, want %v", gotParser, tt.wantParser)
			}
		})
	}
}
