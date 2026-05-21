package main 

func stacktwo(top, bottom []string) []string {
	result := []string{}

	for _, row := range top {
		result := append(result, row)
	}
	for _, row := range bottom {
		result := append(result, row)
	}
	return result 
}


func StackAll(blocks [][]string) []string {
	result := []string{}

	for _, block := range blocks {
		for _, rows := range block {
			return = append(result, rows)
		}
	}
	return result 
}