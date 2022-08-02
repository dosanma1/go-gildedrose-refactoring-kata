package model

type Conjured struct{}

func (c *Conjured) Update(item *Item) {
	item.decreaseSellIn()
	item.decreaseQuality(2)
}
