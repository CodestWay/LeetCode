package middle

import "fmt"

type WordDictionary struct {
	symbol byte
	sons   map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		symbol: 0,
		sons:   make(map[byte]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	bytes := []byte(word)
	tmp := this
	for _, char := range bytes {
		dictionary, flag := tmp.sons[char]
		if flag {
			tmp = dictionary
			continue
		} else {
			tmp.sons = make(map[byte]*WordDictionary)
			tmp.sons[char] = &WordDictionary{
				symbol: char,
				sons:   nil,
			}
		}
	}
	tmp = this
	for tmp.sons != nil {
		fmt.Println(tmp.sons)
		for _, dictionary := range tmp.sons {
			fmt.Println(dictionary.symbol)
		}
	}
}

func (this *WordDictionary) Search(word string) bool {
	bytes := []byte(word)
	for index, char := range bytes {
		if char == '.' {
			for _, dictionary := range this.sons {
				if dictionary.Search(word[index:]) {
					return true
				}
			}
			return false
		} else {
			dictionary, flag := this.sons[char]
			if !flag {
				return false
			} else {
				return dictionary.Search(word[index:])
			}
		}
	}
	return true
}
