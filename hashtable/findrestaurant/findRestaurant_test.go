package findrestaurant

import (
	"reflect"
	"testing"
)

type Tables struct {
	list1  []string
	list2  []string
	expect []string
}

func TestFindRestaurant(t *testing.T) {
	var result []string
	tables := []Tables{
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			[]string{"Shogun"}},
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"KFC", "Shogun", "Burger King"},
			[]string{"Shogun"}},
	}

	for _, v := range tables {
		if result = findRestaurant(v.list1, v.list2); reflect.DeepEqual(result, v.expect) != true {
			t.Errorf("list1=%v, list2=%v, get %v, expect %v", v.list1, v.list2, result, v.expect)
		}
	}
}
