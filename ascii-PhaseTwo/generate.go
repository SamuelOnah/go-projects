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
	for i := 'A'; i < 'Z'; i++ {
		style[i] = []string{

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
	if !ok {
		return []string{}
	}
	art := data[:]
	res = append(res, art...)
	return res
}
