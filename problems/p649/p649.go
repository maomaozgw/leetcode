package p649

func predictPartyVictory(senate string) string {
	var rq, dq = make([]int, 0), make([]int, 0)
	for i := range senate {
		if senate[i] == 'R' {
			rq = append(rq, i)
		} else {
			dq = append(dq, i)
		}
	}
	for len(rq) > 0 && len(dq) > 0 {
		if rq[0] < dq[0] {
			rq = append(rq, rq[0]+len(senate))
		} else {
			dq = append(dq, dq[0]+len(senate))
		}
		rq = rq[1:]
		dq = dq[1:]
	}
	if len(rq) > 0 {
		return "Radiant"
	}
	return "Dire"
}
