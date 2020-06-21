package model

func (o *EdgeOrder) ExistsOrder() bool {
	return o != nil && o.Key != nil
}
