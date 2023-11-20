package matcher

import (
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
)

// NewKeyword creates Fragment orFragmentFold matcher with non embedded option
func NewKeyword(value string, options ...Option) parsly.Matcher {
	caseOpt := &option.Case{}
	embed := false
	matchedOptions := AssignOption(options, &caseOpt)
	caseOpt.Embed = &embed
	if matchedOptions && !caseOpt.Sensitive {
		return &FragmentFold{
			value: []byte(value),
			size:  len(value),
			embed: caseOpt.Embed,
		}
	}
	return &Fragment{
		value: []byte(value),
		size:  len(value),
		embed: caseOpt.Embed,
	}
}
