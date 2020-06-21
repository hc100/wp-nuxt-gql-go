package model

func (c *TextFilterCondition) ExistsFilter() bool {
	if c == nil {
		return false
	}
	if c.FilterWord == "" {
		return false
	}
	if c.MatchingPattern == nil {
		return false
	}
	return true
}

func (c *TextFilterCondition) MatchString() string {
	if c == nil {
		return ""
	}
	matchStr := "%" + c.FilterWord + "%"
	if c.MatchingPattern == nil {
		return matchStr
	}
	if *c.MatchingPattern == MatchingPatternExactMatch {
		matchStr = c.FilterWord
	}
	return matchStr
}
