package main

func intToRoman(num int) string {
	roman := ""
	divider := 1000
	for divider != 0 {
		left := num / divider
		if left >= 1 && left <= 3 {
			for i := 0; i < left; i++ {
				if divider == 1000 {
					roman += "M"
				}
				if divider == 100 {
					roman += "C"
				}
				if divider == 10 {
					roman += "X"
				}
				if divider == 1 {
					roman += "I"
				}
			}
		}
		if left == 4 {
			if divider == 100 {
				roman += "CD"
			}
			if divider == 10 {
				roman += "XL"
			}
			if divider == 1 {
				roman += "IV"
			}
		}
		if left == 5 {
			if divider == 100 {
				roman += "D"
			}
			if divider == 10 {
				roman += "L"
			}
			if divider == 1 {
				roman += "V"
			}
		}
		if left == 6 {
			if divider == 100 {
				roman += "DC"
			}
			if divider == 10 {
				roman += "LX"
			}
			if divider == 1 {
				roman += "VI"
			}
		}
		if left == 7 {
			if divider == 100 {
				roman += "DCC"
			}
			if divider == 10 {
				roman += "LXX"
			}
			if divider == 1 {
				roman += "VII"
			}
		}
		if left == 8 {
			if divider == 100 {
				roman += "DCCC"
			}
			if divider == 10 {
				roman += "LXXX"
			}
			if divider == 1 {
				roman += "VIII"
			}
		}
		if left == 9 {
			if divider == 100 {
				roman += "CM"
			}
			if divider == 10 {
				roman += "XC"
			}
			if divider == 1 {
				roman += "IX"
			}
		}
		num = num - left*divider
		divider /= 10
	}

	return roman
}
