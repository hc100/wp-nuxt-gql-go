package model

import "math"

func (c *PageCondition) ExistsPaging() bool {
	if c == nil {
		return false
	}
	return c.Backward != nil || c.Forward != nil
}

func (c *PageCondition) IsInitialPageView() bool {
	if c == nil {
		return true
	}
	return c.Backward == nil && c.Forward == nil
}

func (c *PageCondition) HasInitialLimit() bool {
	if c == nil {
		return false
	}
	return c.InitialLimit != nil && *c.InitialLimit > 0
}

func (c *PageCondition) TotalPage(totalCount int) int {
	if c == nil {
		return 0
	}
	targetCount := 0
	if c.Backward == nil && c.Forward == nil {
		if c.InitialLimit == nil {
			return 0
		} else {
			targetCount = *c.InitialLimit
		}
	} else {
		if c.Backward != nil {
			targetCount = c.Backward.Last
		}
		if c.Forward != nil {
			targetCount = c.Forward.First
		}
	}
	return int(math.Ceil(float64(totalCount) / float64(targetCount)))
}

func (c *PageCondition) MoveToPageNo() int {
	if c == nil {
		return 1
	}
	if c.Backward == nil && c.Forward == nil {
		return c.NowPageNo
	}
	if c.Backward != nil {
		if c.NowPageNo <= 2 {
			return 1
		}
		return c.NowPageNo - 1
	}
	if c.Forward != nil {
		return c.NowPageNo + 1
	}
	return 1
}
