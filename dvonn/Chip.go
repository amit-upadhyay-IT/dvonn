package dvonn

type Chip struct {
	color ChipColor
}

func GetChips(color ChipColor) Chip {
	return Chip{color}
}

func (c *Chip) GetColor() ChipColor {
	return c.color
}
