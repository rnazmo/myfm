package myfm

import (
	"reflect"
	"testing"
)

type args struct {
	post []byte
}

type wants struct {
	firstIdx    int
	secondIdx   int
	frontmatter []byte
	content     []byte
	wantErr     bool
}

var testcases = []struct {
	name  string
	args  args
	wants wants
}{
	{
		name: "Basic",
		args: args{
			post: []byte(`+++
front_matter_version = "0.0.4"
title = "Foo Bar Baz Foo Bar Baz"
drafted = "2021-01-02-03-45"
created = "2021-06-04-15-33"
last_updated = ""
last_checked = ""
tags = ["meta:tagme", "lang::en", "golang"]
id = "AbCdEfGh"
+++

## Foo

piyo piyo piyo piyo piyo?

## Bar

piyo piyo piyo piyo piyo piyo piyo piyo.
`),
		},
		wants: wants{
			firstIdx:  4,
			secondIdx: 227,
			frontmatter: []byte(`front_matter_version = "0.0.4"
title = "Foo Bar Baz Foo Bar Baz"
drafted = "2021-01-02-03-45"
created = "2021-06-04-15-33"
last_updated = ""
last_checked = ""
tags = ["meta:tagme", "lang::en", "golang"]
id = "AbCdEfGh"
`),
			content: []byte(`
## Foo

piyo piyo piyo piyo piyo?

## Bar

piyo piyo piyo piyo piyo piyo piyo piyo.
`),
			wantErr: false,
		},
	},
	{
		name: "Empty input (zero length)",
		args: args{
			post: []byte(""),
		},
		wants: wants{
			// firstIdx:    0,
			// secondIdx:   0,
			// frontmatter: []byte(""),
			// content:     []byte(""),
			wantErr: true,
		},
	},
	{
		name: "Input is shorter than token",
		args: args{
			post: []byte("++"),
		},
		wants: wants{
			wantErr: true,
		},
	},
	{
		name: "Fontmatter must be at the top of the post",
		args: args{
			post: []byte(`some invalid text+++
front_matter_version = "0.0.4"
title = "Foo Bar Baz Foo Bar Baz"
drafted = "2021-01-02-03-45"
created = "2021-06-04-15-33"
last_updated = ""
last_checked = ""
tags = ["meta:tagme", "lang::en", "golang"]
id = "AbCdEfGh"
+++

## Foo

piyo piyo piyo piyo piyo?
`),
		},
		wants: wants{
			wantErr: true,
		},
	},
	{
		name: "First token not found",
		args: args{
			post: []byte(`
front_matter_version = "0.0.4"
title = "Foo Bar Baz Foo Bar Baz"
drafted = "2021-01-02-03-45"
created = "2021-06-04-15-33"
last_updated = ""
last_checked = ""
tags = ["meta:tagme", "lang::en", "golang"]
id = "AbCdEfGh"
+++

## Foo

piyo piyo piyo piyo piyo?
`),
		},
		wants: wants{
			wantErr: true,
		},
	},
	{
		name: "Second token not found",
		args: args{
			post: []byte(`+++
front_matter_version = "0.0.4"
title = "Foo Bar Baz Foo Bar Baz"
drafted = "2021-01-02-03-45"
created = "2021-06-04-15-33"
last_updated = ""
last_checked = ""
tags = ["meta:tagme", "lang::en", "golang"]
id = "AbCdEfGh"

## Foo

piyo piyo piyo piyo piyo?
`),
		},
		wants: wants{
			wantErr: true,
		},
	},

	// {
	// 	name: "Tokens exists in contents",
	// 	// TODO:
	// },

	// TODO: Add testcases.
}

func Test_parseIndex(t *testing.T) {
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstIdx, gotSecondIdx, err := parseIndex(tt.args.post)
			if (err != nil) != tt.wants.wantErr {
				t.Errorf("parseIndex() error = %v, wantErr %v", err, tt.wants.wantErr)
				return
			}
			if gotFirstIdx != tt.wants.firstIdx {
				t.Errorf("parseIndex() gotFirstIdx = %v, want %v", gotFirstIdx, tt.wants.firstIdx)
			}
			if gotSecondIdx != tt.wants.secondIdx {
				t.Errorf("parseIndex() gotSecondIdx = %v, want %v", gotSecondIdx, tt.wants.secondIdx)
			}
		})
	}
}

func TestParse(t *testing.T) {
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			gotFrontmatter, gotContent, err := Parse(tt.args.post)
			if (err != nil) != tt.wants.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wants.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFrontmatter, tt.wants.frontmatter) {
				t.Errorf("Parse() gotFrontmatter = %v, want %v", gotFrontmatter, tt.wants.frontmatter)
			}
			if !reflect.DeepEqual(gotContent, tt.wants.content) {
				t.Errorf("Parse() gotContent = %v, want %v", gotContent, tt.wants.content)
			}
		})
	}
}

func Test_unmarshal(t *testing.T) {
	type args struct {
		frontmatter []byte
	}
	tests := []struct {
		name    string
		args    args
		want    tomlData
		wantErr bool
	}{
		{
			name: "todo: nameme",
			args: args{
				frontmatter: []byte(`front_matter_version = "0.0.4"
title = "Foo Bar Baz Foo Bar Baz"
drafted = "2021-01-02-03-45"
created = "2021-06-04-15-33"
last_updated = ""
last_checked = ""
tags = ["meta:tagme", "lang::en", "golang"]
id = "AbCdEfGh"
`),
			},
			want: tomlData{
				FrontMatterVersion: "0.0.4",
				Title:              "Foo Bar Baz Foo Bar Baz",
				Drafted:            "2021-01-02-03-45",
				Created:            "2021-06-04-15-33",
				LastUpdated:        "",
				LastChecked:        "",
				Tags:               []string{"meta:tagme", "lang::en", "golang"},
				ID:                 "AbCdEfGh",
			},
			wantErr: false,
		},
		// TODO: Add testcases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshal(tt.args.frontmatter)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
