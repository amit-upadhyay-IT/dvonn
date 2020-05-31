package testcases

import (
	"../dvonn"
	"testing"
)

func TestHexNodeLevelWiseAdjacentNodes_01(t *testing.T) {

	a1 := dvonn.GetHexNode("a1")
	a2 := dvonn.GetHexNode("a2")
	a3 := dvonn.GetHexNode("a3")
	b1 := dvonn.GetHexNode("b1")
	b2 := dvonn.GetHexNode("b2")
	b3 := dvonn.GetHexNode("b3")
	b4 := dvonn.GetHexNode("b4")
	c1 := dvonn.GetHexNode("c1")
	c2 := dvonn.GetHexNode("c2")
	c3 := dvonn.GetHexNode("c3")
	c4 := dvonn.GetHexNode("c4")
	c5 := dvonn.GetHexNode("c5")
	d1 := dvonn.GetHexNode("d1")
	d2 := dvonn.GetHexNode("d2")
	d3 := dvonn.GetHexNode("d3")
	d4 := dvonn.GetHexNode("d4")
	d5 := dvonn.GetHexNode("d5")
	e1 := dvonn.GetHexNode("e1")
	e2 := dvonn.GetHexNode("e2")
	e3 := dvonn.GetHexNode("e3")
	e4 := dvonn.GetHexNode("e4")
	e5 := dvonn.GetHexNode("e5")
	f1 := dvonn.GetHexNode("f1")
	f2 := dvonn.GetHexNode("f2")
	f3 := dvonn.GetHexNode("f3")
	f4 := dvonn.GetHexNode("f4")
	f5 := dvonn.GetHexNode("f5")
	g1 := dvonn.GetHexNode("g1")
	g2 := dvonn.GetHexNode("g2")
	g3 := dvonn.GetHexNode("g3")
	g4 := dvonn.GetHexNode("g4")
	g5 := dvonn.GetHexNode("g5")
	h1 := dvonn.GetHexNode("h1")
	h2 := dvonn.GetHexNode("h2")
	h3 := dvonn.GetHexNode("h3")
	h4 := dvonn.GetHexNode("h4")
	h5 := dvonn.GetHexNode("h5")
	i1 := dvonn.GetHexNode("i1")
	i2 := dvonn.GetHexNode("i2")
	i3 := dvonn.GetHexNode("i3")
	i4 := dvonn.GetHexNode("i4")
	i5 := dvonn.GetHexNode("i5")
	j2 := dvonn.GetHexNode("j2")
	j3 := dvonn.GetHexNode("j3")
	j4 := dvonn.GetHexNode("j4")
	j5 := dvonn.GetHexNode("j5")
	k3 := dvonn.GetHexNode("k3")
	k4 := dvonn.GetHexNode("k4")
	k5 := dvonn.GetHexNode("k5")


	a1.SetChildNode(dvonn.BOTTON_EDGE, b1)
	a1.SetChildNode(dvonn.RIGHTUPPER_EDGE, a2)
	a1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, b2)

	a2.SetChildNode(dvonn.LEFTBOTTON_EDGE, a1)
	a2.SetChildNode(dvonn.BOTTON_EDGE, b2)
	a2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, b3)
	a2.SetChildNode(dvonn.RIGHTUPPER_EDGE, a3)

	a3.SetChildNode(dvonn.LEFTBOTTON_EDGE, a2)
	a3.SetChildNode(dvonn.BOTTON_EDGE, b3)
	a3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, b4)

	b1.SetChildNode(dvonn.UPPER_EDGE, a1)
	b1.SetChildNode(dvonn.RIGHTUPPER_EDGE, b2)
	b1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, c2)
	b1.SetChildNode(dvonn.BOTTON_EDGE, c1)

	b2.SetChildNode(dvonn.UPPER_EDGE, a2)
	b2.SetChildNode(dvonn.RIGHTUPPER_EDGE, b3)
	b2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, c3)
	b2.SetChildNode(dvonn.BOTTON_EDGE, c2)
	b2.SetChildNode(dvonn.LEFTBOTTON_EDGE, b1)
	b2.SetChildNode(dvonn.LEFTUPPER_EDGE, a1)

	b3.SetChildNode(dvonn.UPPER_EDGE, a3)
	b3.SetChildNode(dvonn.RIGHTUPPER_EDGE, b4)
	b3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, c4)
	b3.SetChildNode(dvonn.BOTTON_EDGE, c3)
	b3.SetChildNode(dvonn.LEFTBOTTON_EDGE, b2)
	b3.SetChildNode(dvonn.LEFTUPPER_EDGE, a2)

	b4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, c5)
	b4.SetChildNode(dvonn.BOTTON_EDGE, c3)
	b4.SetChildNode(dvonn.LEFTBOTTON_EDGE, b3)
	b4.SetChildNode(dvonn.LEFTUPPER_EDGE, a3)

	c1.SetChildNode(dvonn.UPPER_EDGE, b1)
	c1.SetChildNode(dvonn.RIGHTUPPER_EDGE, c2)
	c1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, d2)
	c1.SetChildNode(dvonn.BOTTON_EDGE, d1)

	c2.SetChildNode(dvonn.UPPER_EDGE, b2)
	c2.SetChildNode(dvonn.RIGHTUPPER_EDGE, c3)
	c2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, d3)
	c2.SetChildNode(dvonn.BOTTON_EDGE, d2)
	c2.SetChildNode(dvonn.LEFTBOTTON_EDGE, c1)
	c2.SetChildNode(dvonn.LEFTUPPER_EDGE, b1)

	c3.SetChildNode(dvonn.UPPER_EDGE, c3)
	c3.SetChildNode(dvonn.RIGHTUPPER_EDGE, c4)
	c3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, d4)
	c3.SetChildNode(dvonn.BOTTON_EDGE, d3)
	c3.SetChildNode(dvonn.LEFTBOTTON_EDGE, c2)
	c3.SetChildNode(dvonn.LEFTUPPER_EDGE, b2)

	c4.SetChildNode(dvonn.UPPER_EDGE, b4)
	c4.SetChildNode(dvonn.RIGHTUPPER_EDGE, c5)
	c4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, d5)
	c4.SetChildNode(dvonn.BOTTON_EDGE, d4)
	c4.SetChildNode(dvonn.LEFTBOTTON_EDGE, c3)
	c4.SetChildNode(dvonn.LEFTUPPER_EDGE, b3)

	c5.SetChildNode(dvonn.BOTTON_EDGE, d5)
	c5.SetChildNode(dvonn.LEFTBOTTON_EDGE, c4)
	c5.SetChildNode(dvonn.LEFTUPPER_EDGE, b4)


	d1.SetChildNode(dvonn.UPPER_EDGE, c1)
	d1.SetChildNode(dvonn.RIGHTUPPER_EDGE, d2)
	d1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, e2)
	d1.SetChildNode(dvonn.BOTTON_EDGE, e1)

	d2.SetChildNode(dvonn.UPPER_EDGE, c2)
	d2.SetChildNode(dvonn.RIGHTUPPER_EDGE, d3)
	d2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, e3)
	d2.SetChildNode(dvonn.BOTTON_EDGE, e2)
	d2.SetChildNode(dvonn.LEFTBOTTON_EDGE, d1)
	d2.SetChildNode(dvonn.LEFTUPPER_EDGE, c1)

	d3.SetChildNode(dvonn.UPPER_EDGE, c3)
	d3.SetChildNode(dvonn.RIGHTUPPER_EDGE, d4)
	d3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, e4)
	d3.SetChildNode(dvonn.BOTTON_EDGE, e3)
	d3.SetChildNode(dvonn.LEFTBOTTON_EDGE, d2)
	d3.SetChildNode(dvonn.LEFTUPPER_EDGE, c2)

	d4.SetChildNode(dvonn.UPPER_EDGE, c4)
	d4.SetChildNode(dvonn.RIGHTUPPER_EDGE, d5)
	d4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, e5)
	d4.SetChildNode(dvonn.BOTTON_EDGE, e4)
	d4.SetChildNode(dvonn.LEFTBOTTON_EDGE, d3)
	d4.SetChildNode(dvonn.LEFTUPPER_EDGE, c3)


	e1.SetChildNode(dvonn.UPPER_EDGE, d1)
	e1.SetChildNode(dvonn.RIGHTUPPER_EDGE, e2)
	e1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, f2)
	e1.SetChildNode(dvonn.BOTTON_EDGE, f1)

	e2.SetChildNode(dvonn.UPPER_EDGE, d2)
	e2.SetChildNode(dvonn.RIGHTUPPER_EDGE, e3)
	e2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, f3)
	e2.SetChildNode(dvonn.BOTTON_EDGE, f2)
	e2.SetChildNode(dvonn.LEFTBOTTON_EDGE, e1)
	e2.SetChildNode(dvonn.LEFTUPPER_EDGE, d1)

	e3.SetChildNode(dvonn.UPPER_EDGE, d3)
	e3.SetChildNode(dvonn.RIGHTUPPER_EDGE, e4)
	e3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, f4)
	e3.SetChildNode(dvonn.BOTTON_EDGE, f3)
	e3.SetChildNode(dvonn.LEFTBOTTON_EDGE, e2)
	e3.SetChildNode(dvonn.LEFTUPPER_EDGE, d2)

	e4.SetChildNode(dvonn.UPPER_EDGE, d4)
	e4.SetChildNode(dvonn.RIGHTUPPER_EDGE, e5)
	e4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, f5)
	e4.SetChildNode(dvonn.BOTTON_EDGE, f4)
	e4.SetChildNode(dvonn.LEFTBOTTON_EDGE, e3)
	e4.SetChildNode(dvonn.LEFTUPPER_EDGE, d3)

	e5.SetChildNode(dvonn.UPPER_EDGE, d5)
	e5.SetChildNode(dvonn.BOTTON_EDGE, f5)
	e5.SetChildNode(dvonn.LEFTBOTTON_EDGE, e4)
	e5.SetChildNode(dvonn.LEFTUPPER_EDGE, d4)


	f1.SetChildNode(dvonn.UPPER_EDGE, e1)
	f1.SetChildNode(dvonn.RIGHTUPPER_EDGE, f2)
	f1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, g2)
	f1.SetChildNode(dvonn.BOTTON_EDGE, g1)

	f2.SetChildNode(dvonn.UPPER_EDGE, e2)
	f2.SetChildNode(dvonn.RIGHTUPPER_EDGE, f3)
	f2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, g3)
	f2.SetChildNode(dvonn.BOTTON_EDGE, g2)
	f2.SetChildNode(dvonn.LEFTBOTTON_EDGE, f1)
	f2.SetChildNode(dvonn.LEFTUPPER_EDGE, e1)

	f3.SetChildNode(dvonn.UPPER_EDGE, e3)
	f3.SetChildNode(dvonn.RIGHTUPPER_EDGE, f4)
	f3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, g4)
	f3.SetChildNode(dvonn.BOTTON_EDGE, g3)
	f3.SetChildNode(dvonn.LEFTBOTTON_EDGE, f2)
	f3.SetChildNode(dvonn.LEFTUPPER_EDGE, e2)

	f4.SetChildNode(dvonn.UPPER_EDGE, e4)
	f4.SetChildNode(dvonn.RIGHTUPPER_EDGE, f5)
	f4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, g5)
	f4.SetChildNode(dvonn.BOTTON_EDGE, g4)
	f4.SetChildNode(dvonn.LEFTBOTTON_EDGE, f3)
	f4.SetChildNode(dvonn.LEFTUPPER_EDGE, e3)

	f5.SetChildNode(dvonn.UPPER_EDGE, e5)
	f5.SetChildNode(dvonn.BOTTON_EDGE, g5)
	f5.SetChildNode(dvonn.LEFTBOTTON_EDGE, f4)
	f5.SetChildNode(dvonn.LEFTUPPER_EDGE, e4)

	g1.SetChildNode(dvonn.UPPER_EDGE, f1)
	g1.SetChildNode(dvonn.RIGHTUPPER_EDGE, g2)
	g1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, h2)
	g1.SetChildNode(dvonn.BOTTON_EDGE, h1)

	g2.SetChildNode(dvonn.UPPER_EDGE, f2)
	g2.SetChildNode(dvonn.RIGHTUPPER_EDGE, g3)
	g2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, h3)
	g2.SetChildNode(dvonn.BOTTON_EDGE, h2)
	g2.SetChildNode(dvonn.LEFTBOTTON_EDGE, g1)
	g2.SetChildNode(dvonn.LEFTUPPER_EDGE, f1)

	g3.SetChildNode(dvonn.UPPER_EDGE, f3)
	g3.SetChildNode(dvonn.RIGHTUPPER_EDGE, g4)
	g3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, h4)
	g3.SetChildNode(dvonn.BOTTON_EDGE, h3)
	g3.SetChildNode(dvonn.LEFTBOTTON_EDGE, g2)
	g3.SetChildNode(dvonn.LEFTUPPER_EDGE, f2)

	g4.SetChildNode(dvonn.UPPER_EDGE, f4)
	g4.SetChildNode(dvonn.RIGHTUPPER_EDGE, g5)
	g4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, h5)
	g4.SetChildNode(dvonn.BOTTON_EDGE, h4)
	g4.SetChildNode(dvonn.LEFTBOTTON_EDGE, g3)
	g4.SetChildNode(dvonn.LEFTUPPER_EDGE, f3)

	g5.SetChildNode(dvonn.UPPER_EDGE, f5)
	g5.SetChildNode(dvonn.BOTTON_EDGE, h5)
	g5.SetChildNode(dvonn.LEFTBOTTON_EDGE, g4)
	g5.SetChildNode(dvonn.LEFTUPPER_EDGE, f4)

	h1.SetChildNode(dvonn.UPPER_EDGE, g1)
	h1.SetChildNode(dvonn.RIGHTUPPER_EDGE, h2)
	h1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, i2)
	h1.SetChildNode(dvonn.BOTTON_EDGE, i1)

	h2.SetChildNode(dvonn.UPPER_EDGE, g2)
	h2.SetChildNode(dvonn.RIGHTUPPER_EDGE, h3)
	h2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, i3)
	h2.SetChildNode(dvonn.BOTTON_EDGE, i2)
	h2.SetChildNode(dvonn.LEFTBOTTON_EDGE, h1)
	h2.SetChildNode(dvonn.LEFTUPPER_EDGE, g1)

	h3.SetChildNode(dvonn.UPPER_EDGE, g3)
	h3.SetChildNode(dvonn.RIGHTUPPER_EDGE, h4)
	h3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, i4)
	h3.SetChildNode(dvonn.BOTTON_EDGE, i3)
	h3.SetChildNode(dvonn.LEFTBOTTON_EDGE, h2)
	h3.SetChildNode(dvonn.LEFTUPPER_EDGE, g2)

	h4.SetChildNode(dvonn.UPPER_EDGE, g4)
	h4.SetChildNode(dvonn.RIGHTUPPER_EDGE, h5)
	h4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, i5)
	h4.SetChildNode(dvonn.BOTTON_EDGE, i4)
	h4.SetChildNode(dvonn.LEFTBOTTON_EDGE, h3)
	h4.SetChildNode(dvonn.LEFTUPPER_EDGE, g3)

	h5.SetChildNode(dvonn.UPPER_EDGE, g5)
	h5.SetChildNode(dvonn.BOTTON_EDGE, i5)
	h5.SetChildNode(dvonn.LEFTBOTTON_EDGE, h4)
	h5.SetChildNode(dvonn.LEFTUPPER_EDGE, g4)

	i1.SetChildNode(dvonn.UPPER_EDGE, h1)
	i1.SetChildNode(dvonn.RIGHTUPPER_EDGE, i2)
	i1.SetChildNode(dvonn.RIGHTBOTTON_EDGE, j2)

	i2.SetChildNode(dvonn.UPPER_EDGE, h2)
	i2.SetChildNode(dvonn.RIGHTUPPER_EDGE, i3)
	i2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, j3)
	i2.SetChildNode(dvonn.BOTTON_EDGE, j2)
	i2.SetChildNode(dvonn.LEFTBOTTON_EDGE, i1)
	i2.SetChildNode(dvonn.LEFTUPPER_EDGE, h1)

	i3.SetChildNode(dvonn.UPPER_EDGE, h3)
	i3.SetChildNode(dvonn.RIGHTUPPER_EDGE, i4)
	i3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, j4)
	i3.SetChildNode(dvonn.BOTTON_EDGE, j3)
	i3.SetChildNode(dvonn.LEFTBOTTON_EDGE, i2)
	i3.SetChildNode(dvonn.LEFTUPPER_EDGE, h2)

	i4.SetChildNode(dvonn.UPPER_EDGE, h4)
	i4.SetChildNode(dvonn.RIGHTUPPER_EDGE, i5)
	i4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, j5)
	i4.SetChildNode(dvonn.BOTTON_EDGE, j4)
	i4.SetChildNode(dvonn.LEFTBOTTON_EDGE, i3)
	i4.SetChildNode(dvonn.LEFTUPPER_EDGE, h3)

	i5.SetChildNode(dvonn.UPPER_EDGE, h5)
	i5.SetChildNode(dvonn.BOTTON_EDGE, j5)
	i5.SetChildNode(dvonn.LEFTBOTTON_EDGE, i4)
	i5.SetChildNode(dvonn.LEFTUPPER_EDGE, h4)

	j2.SetChildNode(dvonn.UPPER_EDGE, i2)
	j2.SetChildNode(dvonn.RIGHTUPPER_EDGE, j3)
	j2.SetChildNode(dvonn.RIGHTBOTTON_EDGE, k3)
	j2.SetChildNode(dvonn.LEFTUPPER_EDGE, i1)

	j3.SetChildNode(dvonn.UPPER_EDGE, i3)
	j3.SetChildNode(dvonn.RIGHTUPPER_EDGE, j4)
	j3.SetChildNode(dvonn.RIGHTBOTTON_EDGE, k4)
	j3.SetChildNode(dvonn.BOTTON_EDGE, k3)
	j3.SetChildNode(dvonn.LEFTBOTTON_EDGE, j2)
	j3.SetChildNode(dvonn.LEFTUPPER_EDGE, i2)

	j4.SetChildNode(dvonn.UPPER_EDGE, i4)
	j4.SetChildNode(dvonn.RIGHTUPPER_EDGE, j5)
	j4.SetChildNode(dvonn.RIGHTBOTTON_EDGE, k5)
	j4.SetChildNode(dvonn.BOTTON_EDGE, k4)
	j4.SetChildNode(dvonn.LEFTBOTTON_EDGE, j3)
	j4.SetChildNode(dvonn.LEFTUPPER_EDGE, i3)

	j5.SetChildNode(dvonn.UPPER_EDGE, i5)
	j5.SetChildNode(dvonn.BOTTON_EDGE, k5)
	j5.SetChildNode(dvonn.LEFTBOTTON_EDGE, j4)
	j5.SetChildNode(dvonn.LEFTUPPER_EDGE, i4)

	k3.SetChildNode(dvonn.UPPER_EDGE, j3)
	k3.SetChildNode(dvonn.RIGHTUPPER_EDGE, k4)
	k3.SetChildNode(dvonn.LEFTUPPER_EDGE, j2)

	k4.SetChildNode(dvonn.UPPER_EDGE, j4)
	k4.SetChildNode(dvonn.RIGHTUPPER_EDGE, k5)
	k4.SetChildNode(dvonn.LEFTBOTTON_EDGE, k3)
	k4.SetChildNode(dvonn.LEFTUPPER_EDGE, j3)

	k5.SetChildNode(dvonn.UPPER_EDGE, j5)
	k5.SetChildNode(dvonn.LEFTBOTTON_EDGE, k4)
	k5.SetChildNode(dvonn.LEFTUPPER_EDGE, j4)



	foundRes := b3.GetStraightAdjacentOnLevel(1)
	expectedRes := []string{"a3", "b4", "c4", "c3", "b2", "a2"}
	if len(expectedRes) != len(foundRes) {
		t.Error("Adjacent nodes length mismatch")
	}
	for _, item := range foundRes {
		if !contains(expectedRes, item.GetIdentifier()) {
			t.Error("For", item)
		}
	}
	foundRes = nil

	foundRes = b3.GetStraightAdjacentOnLevel(2)
	expectedRes = []string{"b1", "d3", "d5"}
	if len(expectedRes) != len(foundRes) {
		t.Error("Adjacent nodes length mismatch")
	}
	for _, item := range foundRes {
		if !contains(expectedRes, item.GetIdentifier()) {
			t.Error("For", item)
		}
	}
	foundRes = nil

	foundRes = b3.GetStraightAdjacentOnLevel(3)
	expectedRes = []string{"e3"}
	if len(expectedRes) != len(foundRes) {
		t.Error("Adjacent nodes length mismatch")
	}
	for _, item := range foundRes {
		if !contains(expectedRes, item.GetIdentifier()) {
			t.Error("For", item)
		}
	}
	foundRes = nil
}

func contains(master []string, key string) bool {
	for _, item := range master {
		if item == key {
			return true
		}
	}
	return false
}
