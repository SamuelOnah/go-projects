package generate

func GeneratePattern(c rune) []string {
	style := map[rune][]string{
		
		'Z': {
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		},
	}

	for x :='A'; x < 'Z'; x++{
		style[x] = []string{
			
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
	
		}
	}
	var res []string
	data, ok := style[c]

	if !ok{
		return []string{}
	}
	art := make([]string, 8)
	for range 8 {
		art = data[:]
	}
	res = append(res, art...)

	return res
}