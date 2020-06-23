package main

// Not is a function implements boolean `NOT` operations.
func Not(situation bool) bool {
	return !situation
}

// And is a function implements boolean `AND` operations.
func And(situation1, stituation2 bool) bool {
	if situation1 != true {
		return false
	} else if stituation2 != true {
		return false
	} else {
		return true
	}
}

// Or is a function implements boolean `OR` operations.
func Or(situation1, situation2 bool) bool {
	if situation1 == true {
		return true
	} else if situation2 == true{
		return true
	} else {
		return false
	}
}

// XOr is a function implements boolean `XOR` operations.
func XOr(situation1, situation2 bool) bool {
	return And(Not(And(situation1, situation2)), Or(situation1, situation2))
}

