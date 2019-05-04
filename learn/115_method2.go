package main

import "fmt"

const (
	White = iota
	Black
	Blue
	Red
	Yellow
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (b1 BoxList) BiggertColor() Color {
	v := 0.00
	k := Color(White)
	for _, b := range b1 {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (b1 BoxList) PaintItBlack() {
	for i := range b1 {
		b1[i].SetColor(Black)
	}
}

func (c Color) String() string {
	strings := []string{"White", "Black", "Blue", "Red", "Yellow"}
	return strings[c]
}

func main() {

	boxes := BoxList{
		Box{4, 4, 4, Red},
		Box{10, 10, 1, Yellow},
		Box{1, 1, 22, Black},
		Box{2, 3, 1, Blue},
		Box{23, 22, 33, White},
		Box{3, 4, 5, Yellow},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}
