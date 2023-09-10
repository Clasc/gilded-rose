package gildedrose

const SULFURAS = "Sulfuras, Hand of Ragnaros"
const BACKSTAGE = "Backstage passes to a TAFKAL80ETC concert"
const AGED_BRIE = "Aged Brie"

type Item struct {
	Name            string
	SellIn, Quality int
}

type Calculatable interface {
	getNewQuality() int
	updateSellin()
}

type Sulfuras struct {
	Item
}

type AgedBrie struct {
	Item
}

type Backstage struct {
	Item
}

// func calculatable(item *Item) Calculatable {
// 	switch item.Name {
// 	case BACKSTAGE:
// 		{
// 			return &Backstage{Item: *item}
// 		}
// 	case SULFURAS:
// 		{
// 			return &Sulfuras{Item: *item}
// 		}
// 	case AGED_BRIE:
// 		{
// 			return &AgedBrie{Item: *item}
// 		}
// 	default:
// 		return item
// 	}
// }

func UpdateItems(items []*Item) {
	for i := 0; i < len(items); i++ {
		items[i].process()
	}
}

func (item *Item) process() {
	item.updateQuality()
	item.updateSellin()
}

func (item *Item) updateQuality() {
	if item.Name != AGED_BRIE && item.Name != BACKSTAGE {
		if item.Quality > 0 {
			if item.Name != SULFURAS {
				item.Quality = item.Quality - 1
			}
		}
		return
	}
	increaseQuality(item)

}
func (item *Item) updateSellin() {
	if item.Name != SULFURAS {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != AGED_BRIE {
			if item.Name != BACKSTAGE {
				if item.Quality > 0 {
					if item.Name != SULFURAS {
						item.Quality -= 1
					}
				}
			} else {
				item.Quality -= item.Quality
			}
		} else {
			if item.Quality < 50 {
				item.Quality += 1
			}
		}
	}
}

func increaseQuality(item *Item) {
	if item.Quality >= 50 {
		return
	}

	item.Quality += +1
	if item.Name == BACKSTAGE {
		if item.SellIn < 11 {
			if item.Quality < 50 {
				item.Quality += +1
			}
		}
		if item.SellIn < 6 {
			if item.Quality < 50 {
				item.Quality += +1
			}
		}
	}
}
