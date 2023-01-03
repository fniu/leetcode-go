package problem551

func checkRecord(s string) bool {

	l := len(s)
	absents := 0
	lates := 0
	for i := 0; i < l; i++ {
		if s[i] == 'A' {
			absents++
			if absents >= 2 {
				return false
			}
		}
		if s[i] == 'L' {
			lates++
			if lates >= 3 {
				return false
			}
		} else {
			lates = 0
		}
	}

	return true
}
