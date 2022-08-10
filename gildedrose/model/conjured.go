package model

type Conjured struct {
	*Item
}

func (c *Conjured) Update() {
	decreaseSellIn(c.Item)
	decreaseQuality(c.Item, 2)
}
