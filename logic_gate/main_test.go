package main

import "testing"

func Test_Not(t *testing.T)  {
	if Not(true) != false {
		t.Error("error")
	}
}

func Test_And(t *testing.T)  {
	if And(true, true) != true {
		t.Error("error")
	} 
	if And(true, false) != false {
		t.Error("error")
	} 
	if And(false, true) != false {
		t.Error("error")
	} 
	if And(false, false) != false {
		t.Error("error")
	}
}

func Test_Or(t *testing.T)  {
	if Or(true, true) != true {
		t.Error("error")
	} 
	if Or(true, false) != true {
		t.Error("error")
	} 
	if Or(false, true) != true {
		t.Error("error")
	} 
	if Or(false, false) != false {
		t.Error("error")
	}
}

func Test_XOr(t *testing.T)  {
	if XOr(true, true) != false {
		t.Error("error")
	} 
	if XOr(true, false) != true {
		t.Error("error")
	} 
	if XOr(false, true) != true {
		t.Error("error")
	} 
	if XOr(false, false) != false {
		t.Error("error")
	}
}
