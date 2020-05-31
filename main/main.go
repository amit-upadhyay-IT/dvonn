package main

import (
	"../dvonn"
	"fmt"
)

func main() {
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

	d2.SetChildNode(dvonn.UPPER_EDGE, c2)
	d2.SetChildNode(dvonn.RIGHTUPPER_EDGE, d3)
	d2.SetChildNode(dvonn.LEFTBOTTON_EDGE, d1)
	d2.SetChildNode(dvonn.LEFTUPPER_EDGE, c1)

	d3.SetChildNode(dvonn.UPPER_EDGE, c3)
	d3.SetChildNode(dvonn.RIGHTUPPER_EDGE, d4)
	d3.SetChildNode(dvonn.LEFTBOTTON_EDGE, d2)
	d3.SetChildNode(dvonn.LEFTUPPER_EDGE, c2)

	d4.SetChildNode(dvonn.UPPER_EDGE, c4)
	d4.SetChildNode(dvonn.RIGHTUPPER_EDGE, d5)
	d4.SetChildNode(dvonn.LEFTBOTTON_EDGE, d3)
	d4.SetChildNode(dvonn.LEFTUPPER_EDGE, c3)

	d5.SetChildNode(dvonn.BOTTON_EDGE, c5)
	d5.SetChildNode(dvonn.LEFTBOTTON_EDGE, d4)
	d5.SetChildNode(dvonn.LEFTUPPER_EDGE, c4)

	res := b3.GetStraightAdjacentOnLevel(2)
	fmt.Println(res)
}
