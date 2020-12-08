package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/disintegration/imaging"
	"image/color"
	"golang.org/x/image/colornames"
)

type Color struct {
	R, G, B uint8
}

func main(){
	im, _ := imgio.Open("A07Image/flow.jpg")
	fmt.Println(im.At(300,300))
	inverted := effect.Invert(im)
	fmt.Println(inverted.At(300,300))

}

func main2(){
	//fp, err := os.Open(fmt.Sprintf("%s/%s"))
	//nrgba := image.NewRGBA(image.Rect(0,0,100,100))
	//fmt.Println()
	src, err := imaging.Open("A07Image/flow.jpg")
	if err!=nil {
		fmt.Printf("err, %v", err)
	}
	fmt.Printf("%v", src.At(300,300))
	img := src.Bounds()
	var pix []uint8
	newImg := imaging.New(img.Max.X, img.Max.Y, color.Gray{})
	for y:=img.Min.Y; y<img.Max.Y; y++{
		for x:=img.Min.X; x<img.Max.X; x++{
			c := src.At(x, y)
			r,g,b,_ := c.RGBA()
			i := newImg.PixOffset(x, y)
			s := pix[i:i+3:i+3]
			//s := newImg.Pix[i : i+3 : i+3] // Small cap improves performance, see https://golang.org/issue/27857
			s[0] = uint8(b >> 8)
			s[1] = uint8(g >> 8)
			s[2] = uint8(r >> 8)
		}
	}
	fmt.Printf("%v", newImg.At(300,300))
	imaging.Save(newImg,"f.jpg")
}
