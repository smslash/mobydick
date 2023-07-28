package functions

func WordCount(s [][]byte) WordList {
	wordList := make(WordList, 0)
	for i := 0; i < len(s); i++ {
		index := checkIndex(wordList, s[i])
		if index != -1 {
			wordList[index].Count++
		} else {
			wordList = append(wordList, Word{Word: s[i], Count: 1})
		}
	}
	return wordList
}

func checkIndex(w WordList, word []byte) int {
	for i, wl := range w {
		if isExist(wl.Word, word) {
			return i
		}
	}
	return -1
}

func isExist(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
