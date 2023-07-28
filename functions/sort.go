package functions

func Sort(w WordList) {
	for i := 0; i < len(w)-1; i++ {
		swapped := false
		for j := 0; j < len(w)-i-1; j++ {
			if w[j+1].Count > w[j].Count {
				w[j], w[j+1] = w[j+1], w[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
