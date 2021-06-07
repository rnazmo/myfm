package formatter

import (
	"reflect"
	"testing"
)

func Test_formatTags(t *testing.T) {
	type args struct {
		validatedTags []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Basic",
			args: args{
				[]string{"meta:tagme", "lang::en", "golang", "github"},
			},
			want: []string{"github", "golang", "lang::en", "meta:tagme"},
		},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTags(tt.args.validatedTags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
