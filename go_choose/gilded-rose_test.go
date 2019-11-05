package main

import (
	"reflect"
	"testing"
)

func Test_Foo(t *testing.T) {
	var items = []*Item{
		&Item{"+5 Dexterity Vest", 2, 1},
		&Item{"+5 Dexterity Vest", 1, 1},
		&Item{"+5 Dexterity Vest", 1, 0},
		&Item{"+5 Dexterity Vest", 0, 2},
		&Item{"+5 Dexterity Vest", 0, 1},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", 1, 80},
		&Item{"Aged Brie", 1, 49},
		&Item{"Aged Brie", 1, 50},
		&Item{"Aged Brie", 0, 48},
		&Item{"Aged Brie", 0, 49},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 11, 49},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 11, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 47},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 48},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 1, 48},
	}
	var items_expect = []*Item{
		&Item{"+5 Dexterity Vest", 1, 0},
		&Item{"+5 Dexterity Vest", 0, 0},
		&Item{"+5 Dexterity Vest", 0, 0},
		&Item{"+5 Dexterity Vest", -1, 0},
		&Item{"+5 Dexterity Vest", -1, 0},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", 1, 80},
		&Item{"Aged Brie", 0, 50},
		&Item{"Aged Brie", 0, 50},
		&Item{"Aged Brie", -1, 50},
		&Item{"Aged Brie", -1, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
	}

	UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		if reflect.DeepEqual(*items[i], *items_expect[i]) {
		} else {
			t.Errorf("Name: %s Expected %d but got %d,day: %d", items[i].name, items_expect[i].quality, items[i].quality, items[i].sellIn)
		}
	}
}
