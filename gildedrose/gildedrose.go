package gildedrose

import (
	"strings"

	"github.com/dosanma1/go-gildedrose-refactoring-kata/gildedrose/model"
)

const (
	AGED_BRIE = "Aged Brie"
	SULFURAS  = "Sulfuras, Hand of Ragnaros"
	BACKSTAGE = "Backstage passes to a TAFKAL80ETC concert"
	CONJURED  = "Conjured"
)

func UpdateQuality(items []*model.Item) {
	var item *model.Item
	var uItem model.UpdatableItem
	for i := 0; i < len(items); i++ {

		item = items[i]

		switch item.Name {
		case SULFURAS:
			uItem = &model.Sulfuras{}
		case AGED_BRIE:
			uItem = &model.AgedBrie{}
		case BACKSTAGE:
			uItem = &model.Backstage{}
		default:
			if strings.HasPrefix(item.Name, CONJURED) {
				uItem = &model.Conjured{}
			} else {
				uItem = &model.StandardItem{}
			}
		}

		uItem.Update(item)
	}

}
