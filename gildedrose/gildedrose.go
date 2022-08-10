package gildedrose

import (
	"strings"

	"github.com/dosanma1/go-gildedrose-refactoring-kata/gildedrose/model"
)

const (
	agedBrie  = "Aged Brie"
	sulfuras  = "Sulfuras, Hand of Ragnaros"
	backstage = "Backstage passes to a TAFKAL80ETC concert"
	conjured  = "Conjured"
)

func UpdateQuality(items []*model.Item) {
	var uItem model.UpdatableItem
	for _, item := range items {

		switch item.Name {
		case sulfuras:
			uItem = &model.Sulfuras{Item: item}
		case agedBrie:
			uItem = &model.AgedBrie{Item: item}
		case backstage:
			uItem = &model.Backstage{Item: item}
		default:
			if strings.HasPrefix(item.Name, conjured) {
				uItem = &model.Conjured{Item: item}
			} else {
				uItem = &model.StandardItem{Item: item}
			}
		}

		uItem.Update()
	}

}
