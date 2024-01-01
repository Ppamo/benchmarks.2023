package striphtml

import (
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/microcosm-cc/bluemonday"
)

func StripHtmlTagsWithRegexpP1(text string) string {
	const pattern = `<.*?>`
	var (
		r *regexp.Regexp
	)
	r = regexp.MustCompile(pattern)
	return r.ReplaceAllString(text, "")
}

func StripHtmlTagsWithRegexpP2(text string) string {
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	var (
		r      *regexp.Regexp
		groups []string
	)
	r = regexp.MustCompile(pattern)
	groups = r.FindAllString(text, -1)
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			text = strings.ReplaceAll(text, group, "")
		}
	}
	return text
}

func StripHtmlTagsWithStringBuilderV1(text string) string {
	const (
		htmlTagStart = 60
		htmlTagEnd   = 62
	)
	var (
		in      bool = false
		start   int  = 0
		end     int  = 0
		builder strings.Builder
	)
	builder.Grow(len(text) + utf8.UTFMax)

	for index, chr := range text {
		if (index+1) == len(text) && end >= start {
			if !in {
				builder.WriteString(text[end:])
			}
			break
		}
		if chr != htmlTagStart && chr != htmlTagEnd {
			continue
		}
		if chr == htmlTagStart {
			if !in {
				start = index
			}
			in = true
			builder.WriteString(text[end:start])
			continue
		}
		in = false
		end = index + 1
	}
	return builder.String()
}

func StripHtmlTagsWithStringBuilderV2(text string) string {
	const (
		htmlTagStart = 60
		htmlTagEnd   = 62
	)
	var (
		in      bool   = false
		start   int    = 0
		end     int    = 0
		d       []rune = []rune(text)
		builder strings.Builder
	)
	builder.Grow(len(d) + utf8.UTFMax)

	for index, chr := range d {
		if (index+1) == len(d) && end >= start {
			if !in {
				builder.WriteString(string(d[end:]))
			}
			break
		}
		if chr != htmlTagStart && chr != htmlTagEnd {
			continue
		}
		if chr == htmlTagStart {
			if !in {
				start = index
			}
			in = true
			builder.WriteString(string(d[end:start]))
			continue
		}
		in = false
		end = index + 1
	}
	return builder.String()
}

func StripHtmlTagsWithBlueMonday(text string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(text)
}
