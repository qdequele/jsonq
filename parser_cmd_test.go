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
		{"retrieve only", args{"{}"}, &Query{[]*Filter{}, map[string]*Query{}, []string{}, false}},
		{"retrieve only", args{"{a,b,c}"}, &Query{[]*Filter{}, map[string]*Query{}, []string{"a", "b", "c"}, false}},
		{"retrieve only", args{"{a, b, c}"}, &Query{[]*Filter{}, map[string]*Query{}, []string{"a", "b", "c"}, false}},
		{"filter only", args{"(a : 1){}"}, &Query{[]*Filter{&Filter{"a", ":", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a:1){}"}, &Query{[]*Filter{&Filter{"a", ":", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a :: 1){}"}, &Query{[]*Filter{&Filter{"a", "::", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a::1){}"}, &Query{[]*Filter{&Filter{"a", "::", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a>1){}"}, &Query{[]*Filter{&Filter{"a", ">", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a > 1){}"}, &Query{[]*Filter{&Filter{"a", ">", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a<1){}"}, &Query{[]*Filter{&Filter{"a", "<", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a < 1){}"}, &Query{[]*Filter{&Filter{"a", "<", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a=1){}"}, &Query{[]*Filter{&Filter{"a", "=", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a = 1){}"}, &Query{[]*Filter{&Filter{"a", "=", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a!=1){}"}, &Query{[]*Filter{&Filter{"a", "!=", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter only", args{"(a != 1){}"}, &Query{[]*Filter{&Filter{"a", "!=", 1}}, map[string]*Query{}, []string{}, false}},
		{"filter twice", args{"(a = 1 && b > 0){}"}, &Query{[]*Filter{&Filter{"a", "=", 1}, &Filter{"b", ">", 0}}, map[string]*Query{}, []string{}, false}},
		{"filter  and retrieve", args{"(a = 1 && b > 0){a,b,c{x,y,z}}"}, &Query{[]*Filter{&Filter{"a", "=", 1}, &Filter{"b", ">", 0}}, map[string]*Query{"c": &Query{[]*Filter{}, map[string]*Query{}, []string{"x", "y", "z"}, false}}, []string{"a", "b"}, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParser := MustParseQuery(tt.args.cmd); !gotParser.eq(*tt.wantParser) {
				t.Errorf("MustParseQuery() = %v, want %v", gotParser, tt.wantParser)
			}
		})
	}
}

func TestParseQuery(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name       string
		args       args
		wantParser *Query
		wantErr    bool
	}{
		{"retrieve only", args{"{}"}, &Query{[]*Filter{}, map[string]*Query{}, []string{}, false}, false},
		{"retrieve only", args{"{a,b,c}"}, &Query{[]*Filter{}, map[string]*Query{}, []string{"a", "b", "c"}, false}, false},
		{"retrieve only", args{"{a, b, c}"}, &Query{[]*Filter{}, map[string]*Query{}, []string{"a", "b", "c"}, false}, false},
		{"filter only", args{"(a : 1){}"}, &Query{[]*Filter{&Filter{"a", ":", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a:1){}"}, &Query{[]*Filter{&Filter{"a", ":", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a :: 1){}"}, &Query{[]*Filter{&Filter{"a", "::", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a::1){}"}, &Query{[]*Filter{&Filter{"a", "::", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a>1){}"}, &Query{[]*Filter{&Filter{"a", ">", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a > 1){}"}, &Query{[]*Filter{&Filter{"a", ">", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a<1){}"}, &Query{[]*Filter{&Filter{"a", "<", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a < 1){}"}, &Query{[]*Filter{&Filter{"a", "<", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a=1){}"}, &Query{[]*Filter{&Filter{"a", "=", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a = 1){}"}, &Query{[]*Filter{&Filter{"a", "=", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a!=1){}"}, &Query{[]*Filter{&Filter{"a", "!=", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter only", args{"(a != 1){}"}, &Query{[]*Filter{&Filter{"a", "!=", 1}}, map[string]*Query{}, []string{}, false}, false},
		{"filter twice", args{"(a = 1 && b > 0){}"}, &Query{[]*Filter{&Filter{"a", "=", 1}, &Filter{"b", ">", 0}}, map[string]*Query{}, []string{}, false}, false},
		{"filter  and retrieve", args{"(a = 1 && b > 0){a,b,c{x,y,z}}"}, &Query{[]*Filter{&Filter{"a", "=", 1}, &Filter{"b", ">", 0}}, map[string]*Query{"c": &Query{[]*Filter{}, map[string]*Query{}, []string{"x", "y", "z"}, false}}, []string{"a", "b"}, false}, false},
		{"retrieve only", args{"{"}, nil, true},
		{"retrieve only", args{"{a,b,c"}, nil, true},
		{"filter only", args{"( : 1){}"}, nil, true},
		{"filter only", args{"(a:){}"}, nil, true},
		{"filter only", args{"(a ::: 1){}"}, nil, true},
		{"filter only", args{"(a>1 || b < c){}"}, nil, true},
		{"filter only", args{"(a > 1{}"}, nil, true},
		{"filter only", args{"a<1){}"}, nil, true},
		{"filter only", args{"(a){}"}, nil, true},
		{"filter only", args{"(a === 1){}"}, nil, true},
		{"filter only", args{"(a?1){}"}, nil, true},
		{"filter only", args{"(ac 1){}"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParser, err := ParseQuery(tt.args.cmd)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseQuery() error = %v, wantErr %v : %s", err, tt.wantErr, tt.args)
				}
			} else if tt.wantParser == nil || !gotParser.eq(*tt.wantParser) {
				t.Errorf("ParseQuery() = %v, want %v", gotParser, tt.wantParser)
			}

		})
	}
}
