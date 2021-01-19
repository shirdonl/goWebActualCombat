//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
)

func main() {
	//需要加水印的图片
	imgfile, _ := os.Open("./pic.jpg")
	defer imgfile.Close()

	jpgimg, _ := jpeg.Decode(imgfile)

	img := image.NewNRGBA(jpgimg.Bounds())

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, jpgimg.At(x, y))
		}
	}
	//拷贝一个字体文件到运行目录
	fontBytes, err := ioutil.ReadFile("simsun.ttc")
	if err != nil {
		log.Println(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}

	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(font)
	f.SetFontSize(12)
	f.SetClip(jpgimg.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255}))

	pt := freetype.Pt(img.Bounds().Dx()-200, img.Bounds().Dy()-12)
	_, err = f.DrawString("Shirdon", pt)

	//draw.Draw(img,jpgimg.Bounds(),jpgimg,image.ZP,draw.Over)

	//保存到新文件中
	newfile, _ := os.Create("./aaa.jpg")
	defer newfile.Close()

	err = jpeg.Encode(newfile, img, &jpeg.Options{100})
	if err != nil {
		fmt.Println(err)
	}
}