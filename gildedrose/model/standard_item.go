package model

type StandardItem struct {
	*Item
}

func (s *StandardItem) Update() {
	decreaseSellIn(s.Item)

	if isExpired(s.Item) {
		decreaseQuality(s.Item, 2)
		return
	}

	decreaseQuality(s.Item, 1)

}
