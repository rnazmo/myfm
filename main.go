package myfm

import (
	"bytes"
	"fmt"

	"github.com/naoina/toml"
	"github.com/rnazmo/myfm/internal/formatter"
)

const (
	IdentifyingToken            = "+++"
	IdentifyingTokenWithNewline = IdentifyingToken + "\n"
)

type invalidatedFrontmatter frontmatter

type frontmatter struct {
	FrontMatterVersion string
	Title              string
	Drafted            string
	Created            string
	LastUpdated        string
	LastChecked        string
	Tags               []string
	ID                 string
}

// TODO: Wrap those error messages.
func NewFromPost(post []byte) (fm frontmatter, content []byte, err error) {
	frontMatterBytes, content, err := Parse(post)
	if err != nil {
		return frontmatter{}, nil, err
	}
	invalidatedFrontMatter, err := unmarshal(frontMatterBytes)
	if err != nil {
		return frontmatter{}, nil, err
	}
	frontMatter, err := validate(invalidatedFrontMatter)
	if err != nil {
		return frontmatter{}, nil, err
	}
	return frontMatter, content, nil
}

func NewFromInputs(
	frontMatterVersion string,
	title string,
	drafted string,
	created string,
	lastUpdated string,
	lastChecked string,
	tags []string,
	id string,
) (frontmatter, error) {
	return validate(
		invalidatedFrontmatter{
			FrontMatterVersion: frontMatterVersion,
			Title:              title,
			Drafted:            drafted,
			Created:            created,
			LastUpdated:        lastUpdated,
			LastChecked:        lastChecked,
			Tags:               tags,
			ID:                 id,
		},
	)
}

// TODO: Wrap those error messages.
// TODO: Should I rename this function to 'newFromInvalidatedFrontmatter'?
func validate(invalidatedFrontMatter invalidatedFrontmatter) (frontmatter, error) {
	formatedFrontMatterVersion, err := formatter.ValidateAndFormatFrontMatterVersion(invalidatedFrontMatter.FrontMatterVersion)
	if err != nil {
		return frontmatter{}, err
	}
	formatedTitle, err := formatter.ValidateAndFormatTitle(invalidatedFrontMatter.Title)
	if err != nil {
		return frontmatter{}, err
	}
	formatedDrafted, err := formatter.ValidateAndFormatDrafted(invalidatedFrontMatter.Drafted)
	if err != nil {
		return frontmatter{}, err
	}
	formatedCreated, err := formatter.ValidateAndFormatCreated(invalidatedFrontMatter.Created)
	if err != nil {
		return frontmatter{}, err
	}
	formatedLastUpdated, err := formatter.ValidateAndFormatLastUpdated(invalidatedFrontMatter.LastUpdated)
	if err != nil {
		return frontmatter{}, err
	}
	formatedLastChecked, err := formatter.ValidateAndFormatLastChecked(invalidatedFrontMatter.LastChecked)
	if err != nil {
		return frontmatter{}, err
	}
	formatedTags, err := formatter.ValidateAndFormatTags(invalidatedFrontMatter.Tags)
	if err != nil {
		return frontmatter{}, err
	}
	formatedID, err := formatter.ValidateAndFormatID(invalidatedFrontMatter.ID)
	if err != nil {
		return frontmatter{}, err
	}
	return frontmatter{
		FrontMatterVersion: formatedFrontMatterVersion,
		Title:              formatedTitle,
		Drafted:            formatedDrafted,
		Created:            formatedCreated,
		LastUpdated:        formatedLastUpdated,
		LastChecked:        formatedLastChecked,
		Tags:               formatedTags,
		ID:                 formatedID,
	}, nil
}

// parseIndex do following:
//   1. Finds first two 'IdentifyingTokenWithNewline' in 'post'.
//   2. Recognizes 'frontmatter' and 'contents'.
//   3. Returns the indexes of 'frontmatter' and 'contents'.
// It returns error if it failed to find 'IdentifyingTokenWithNewline'.
//
// Return value:
//   firstIdx:
//     Index of the starting point (first byte) of the 'frontmatter' (excluding 'IdentifyingToken').
//     NOTE: 'frontmatter' should be at the top of the 'post'. Therefore, 'firstIdx' is always 4.
//   secondIdx:
//     Index of the starting point (first byte) of the 'content'.
//
// Example:
//
//     f, s, _ := parseIndex(post)
//     frontmatter := post[f:s-4] // 4 = len(IdentifyingTokenWithNewline)
//     content := post[s:]
//
//   When the function is called like this, 'post' is equal to
//   'IdentifyingTokenWithNewline' + 'frontmatter' + 'IdentifyingTokenWithNewline' + 'content'.
//
func parseIndex(post []byte) (firstIdx, secondIdx int, err error) {
	const n = len(IdentifyingTokenWithNewline)
	// Special cases
	switch {
	case len(post) == 0:
		return 0, 0, fmt.Errorf("parseIndex(): invalid input. zero length")
	case len(post) <= n:
		return 0, 0, fmt.Errorf("parseIndex(): invalid input. shorter than token")
	}

	// The first token is always at the top of the 'post'
	const first = n // it's always 4

	// Find the second token
	i := bytes.Index(
		post[first:], // NOTE: Don't forget that 'i' is the index of 'post[first:]', no of 'post'.
		[]byte(IdentifyingTokenWithNewline),
	)
	if i == -1 {
		return 0, 0, fmt.Errorf(
			"parseIndex(): failed to find the second IdentifyingTokenWithNewline",
		)
	}
	j := i + first // NOTE: 'j' is the index of 'post', no of 'post[first:]'.

	second := j + n

	// Verify that tokens (based on the indexes) are correct
	firstToken := string(post[first-n : first])    // first-n = 0
	secondToken := string(post[second-n : second]) // post[second-n:second] = post[j:j+n] = post[first:][i:i+n]
	if firstToken != IdentifyingTokenWithNewline {
		return 0, 0, fmt.Errorf(
			"parseIndex(): failed to find the first IdentifyingTokenWithNewline",
		)
	}
	if secondToken != IdentifyingTokenWithNewline {
		return 0, 0, fmt.Errorf(
			"parseIndex(): failed to find the second IdentifyingTokenWithNewline",
		)
	}

	return first, second, nil
}

// TODO: Add document.
func Parse(post []byte) (frontmatter, content []byte, err error) {
	f, s, err := parseIndex(post)
	if err != nil {
		return nil, nil, err
	}
	return post[f : s-4], post[s:], nil // 4 = len(IdentifyingTokenWithNewline)
}

// unmarshal converts a front matter from toml to struct.
//
// Ref:
//   https://github.com/naoina/toml
//   https://pkg.go.dev/github.com/naoina/toml?utm_source=godoc#example-package-TextUnmarshalerError
//
// TODO: Use https://github.com/pelletier/go-toml instead.
//
// TODO: Add test
//
func unmarshal(frontmatter []byte) (invalidatedFrontmatter, error) {
	var td invalidatedFrontmatter
	if err := toml.Unmarshal(frontmatter, &td); err != nil {
		return invalidatedFrontmatter{}, err
	}
	return td, nil
}
