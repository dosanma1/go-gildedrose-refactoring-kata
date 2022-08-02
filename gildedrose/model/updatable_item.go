package model

type UpdatableItem interface {
	Update(item *Item)
}
