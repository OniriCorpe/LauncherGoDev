package UI

import (
	
	"github.com/veandco/go-sdl2/ttf"
	
	"github.com/cuu/gogame/surface"
	"github.com/cuu/gogame/color"
	"github.com/cuu/gogame/font"
	"github.com/cuu/gogame/draw"
	
)

type TextItem struct {
	IconItem
	Str string
	Color *color.Color
	FontObj *ttf.Font
	Bold bool
}

func NewTextItem() *TextItem {
	p := &TextItem{}
	p.Color = &color.Color{83,83,83,255}
	p.MyType = ICON_TYPES["LETTER"]
	p.Bold  = false
	
	return p
}

func (self *TextItem) Draw() {
	font.SetBold(self.FontObj,self.Bold)

	my_text := font.Render(self.FontObj,self.Str,true,self.Color,nil)

	if surface.GetWidth(my_text) != self.Width {
		self.Width = surface.GetWidth(my_text)
	}
	
	if surface.GetHeight(my_text) != self.Height {
		self.Height = surface.GetHeight(my_text)
	}


	rect_ := draw.MidRect(self.PosX,self.PosY,self.Width,self.Height,Width,Height)
	surface.Blit(self.Parent.GetCanvasHWND(),my_text,rect_,nil)
	
}