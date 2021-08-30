package providergen

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want FieldAttributes
	}{
		{"none", args{``}, FieldAttributes{false, false, false, false}},
		{"key", args{`hotcereal:"key"`}, FieldAttributes{true, false, false, false}},
		{"searchable key", args{`hotcereal:"key,searchable"`}, FieldAttributes{true, true, false, false}},
		{"lookup", args{`hotcereal:"lookup"`}, FieldAttributes{false, false, true, false}},
		{"searchable lookup", args{`hotcereal:"lookup,searchable"`}, FieldAttributes{false, true, true, false}},
		{"searchable", args{`hotcereal:"searchable"`}, FieldAttributes{false, true, false, false}},
		{"lazy", args{`hotcereal:"lazy"`}, FieldAttributes{false, false, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTag(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
