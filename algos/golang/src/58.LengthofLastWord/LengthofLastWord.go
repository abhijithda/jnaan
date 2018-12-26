package lengthoflastword

func lengthOfLastWord(s string) int {
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if cnt != 0 {
				return cnt
			}
			continue
		}
		cnt++
	}

	return cnt
}

func lengthOfLastWordOld(s string) int {
	pcnt := 0
	cnt := 0

	for c := range s {
		if string(s[c]) == " " {
			if cnt != 0 {
				pcnt = cnt
			}
			cnt = 0
			continue
		}
		cnt++
	}

	if cnt == 0 {
		return pcnt
	}
	return cnt
}
