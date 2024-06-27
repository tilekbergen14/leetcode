package main

func findCenter(edges [][]int) int {
	first := edges[0][0]
	inO := true
	for index, arr := range edges {
		if index != 0 {
			in := false
			for _, current := range arr {
				if current == first {
					in = true
				}
			}
			if !in {
				inO = false
				break
			}
		}
	}
	second := edges[0][1]
	inO1 := true
	for index, arr := range edges {
		if index != 0 {
			in := false
			for _, current := range arr {
				if current == second {
					in = true
				}
			}
			if !in {
				inO1 = false
				break
			}
		}
	}
	if inO {
		return first
	}
	if inO1 {
		return second
	}
	return 0
}
