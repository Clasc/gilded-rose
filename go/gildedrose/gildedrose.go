package gildedrose

const SULFURAS = "Sulfuras, Hand of Ragnaros"
const BACKSTAGE = "Backstage passes to a TAFKAL80ETC concert"
const AGED_BRIE = "Aged Brie"

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		process(items[i])
	}
}

func process(item *Item) {
	if item.Name != AGED_BRIE && item.Name != BACKSTAGE {
		if item.Quality > 0 {
			if item.Name != SULFURAS {
				item.Quality = item.Quality - 1
			}
		}
	} else {
		increaseQuality(item)
	}

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

	item.Quality = item.Quality + 1
	if item.Name == BACKSTAGE {
		if item.SellIn < 11 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
		if item.SellIn < 6 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
}
