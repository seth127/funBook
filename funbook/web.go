package funbook

import (
	"strings"
	"regexp"
)

var parseDict = map[string]string {
	"&rsquo;" : "'",
	"&mdash;" : "--",
	"&ldquo;" : `"`,
	"&rdquo;" : `"`,
}

func ParseHtml(s string) (bool, string) {

	// trim whitespace
	ms := regexp.MustCompile(` +`)
	s = ms.ReplaceAllString(s, " ")

	// parse from the dict
	for old, new := range parseDict {
		s = strings.ReplaceAll(s, old, new)
	}

	// replace html tags
	mp := regexp.MustCompile(`<.*?>`)
	s = mp.ReplaceAllString(s, "")

	// if there are no letters return false
	ml := regexp.MustCompile("[A-Za-z]")
	b :=  ml.MatchString(s)
	return b, s
}