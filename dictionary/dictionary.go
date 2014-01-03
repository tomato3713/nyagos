package Dictionary

import "strings"
import "sort"

type KeyValuePair struct {
	Key       string
	Lower_key string
	Value     string
}

type Dictionary map[string]*KeyValuePair

func NewDictionary() Dictionary {
	return make(Dictionary)
}

func (this *Dictionary) Get(key string) string {
	lower_key := strings.ToLower(key)

	node, ok := (*this)[lower_key]
	if ok {
		return node.Value
	} else {
		return ""
	}
}

func (this *Dictionary) Remove(key string) {
	delete(*this,strings.ToLower(key))
}

func (this *Dictionary) Set(key string, value string) {
	if value == "" {
		this.Remove(key)
		return
	}
	tmp := new(KeyValuePair)
	tmp.Key = key
	tmp.Lower_key = strings.ToLower(key)
	tmp.Value = value
	(*this)[tmp.Lower_key] = tmp
}

func (this *Dictionary) Iter() chan *KeyValuePair {
	ch := make(chan *KeyValuePair, 0)
	go func(this *Dictionary) {
		for _, pair := range *this {
			ch <- pair
		}
		close(ch)
		return
	}(this)
	return ch
}

type KeyValueList struct {
	list []KeyValuePair
}

func (this *KeyValueList) Len() int {
	return len(this.list)
}

func (this *KeyValueList) Less(i, j int) bool {
	return this.list[i].Key < this.list[j].Key
}

func (this *KeyValueList) Swap(i, j int) {
	this.list[i], this.list[j] = this.list[j], this.list[i]
}

func (this *Dictionary) ToArray() []KeyValuePair {
	list := []KeyValuePair{}
	for _, pair := range *this {
		list = append(list, *pair)
	}
	return list
}

func (this *Dictionary) SortIter() chan KeyValuePair {
	array := new(KeyValueList)
	array.list = this.ToArray()
	sort.Sort(array)
	ch := make(chan KeyValuePair)
	go func() {
		for i := 0; i < len(array.list); i++ {
			ch <- array.list[i]
		}
		close(ch)
	}()
	return ch
}
