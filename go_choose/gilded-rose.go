package main

type Item struct {
	name            string
	sellIn, quality int
}

func qualityNoLimit(item *Item) (quality int) {
	// quality:
	// 1) Aged Brie: `Quality`会随着时间推移而提高
	if item.name == "Aged Brie" {
		if item.sellIn < 0 {
			item.quality = item.quality + 2
			return item.quality
		}
		item.quality = item.quality + 1
		return item.quality
	}

	// 2) Backstage passes: sellIn<11时
	if item.name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.sellIn < 0 {
			item.quality = item.quality - item.quality
			return item.quality
		}
		if item.sellIn < 5 {
			item.quality = item.quality + 3
			return item.quality
		}
		if item.sellIn < 10 {
			item.quality = item.quality + 2
			return item.quality
		}
		item.quality = item.quality + 1
		return item.quality
	}

	// 3) 其他: sellIn>0则-1; sellIn<0则-2
	if item.sellIn < 0 {
		item.quality = item.quality - 2
		return item.quality
	}
	item.quality = item.quality - 1

	return item.quality
}

// 品质`Quality`上限下限
func qualityLimit(quality int) int {
	if quality < 0 {
		quality = 0
	}
	if quality > 50 {
		quality = 50
	}
	return quality
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		// Sulfuras: 不会降低品质`Quality`,永不到期
		if items[i].name != "Sulfuras, Hand of Ragnaros" {

			items[i].sellIn = items[i].sellIn - 1

			items[i].quality = qualityNoLimit(items[i])
			items[i].quality = qualityLimit(items[i].quality)

		}
	}
}
