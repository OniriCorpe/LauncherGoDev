package UI

import (
	
	"github.com/veandco/go-sdl2/ttf"
	
//	"github.com/cuu/gogame/surface"
  "github.com/cuu/gogame/event"
  "github.com/cuu/gogame/rect"
	"github.com/cuu/gogame/color"
//	"github.com/cuu/gogame/font"
	"github.com/cuu/gogame/draw"
	
)

type ListPageSelector struct {
  PageSelector
  BackgroundColor *color.Color
}

func NewListPageSelector() *ListPageSelector {
  p := &ListPageSelector{}
  p.Width = Width
  p.BackgroundColor = &color.Color{131,199,219,255}
  return p
}

func (self *ListPageSelector) Draw() {
  idx := self.Parent.GetPsIndex()
  mylist := self.Parent.GetMyList()
  if idx > (len(mylist) -1) {
    idx = len(mylist)
    if idx > 0 {
      idx -= 1
    }else if idx == 0 {
      return
    }
  }
  
  x,y  := mylist[idx].Coord()
  _,h  := mylist[idx].Size()
  
  self.PosX = x
  self.PosY = y
  self.Height = h -3
  
  canvas_ := self.Parent.GetCanvasHWND()
  rect_ := rect.Rect(self.PosX, self.PosY, self.Width-4, self.Height)
  draw.AARoundRect(canvas_,&rect_,self.BackgroundColor,4,0,self.BackgroundColor)
  
  
}

type ConfirmPage struct {
  Page
  ListFont  *ttf.Font
  FileName string
  TrashDir string
  ConfirmText string
  BGPosX    int
  BGPosY    int
  BGWidth   int
  BGHeight  int
  Icons map[string]IconItemInterface
  
  MyList []LabelInterface
  
}


func NewConfirmPage() *ConfirmPage  {
  p := &ConfirmPage{}
  p.ListFont = Fonts["veramono20"]
  p.FootMsg = [5]string{"Nav","","","Cancel","Yes"}
  p.ConfirmText ="Confirm?"
  
  return p
}


func (self *ConfirmPage) GetMyList() []LabelInterface {
	return self.MyList
}

func (self *ConfirmPage) Reset() {
  self.MyList[0].SetText(self.ConfirmText)
  
  self.MyList[0].PosX = (self.Width - self.MyList[0].Width)/2
  self.MyList[0].PosY = (self.Width - self.MyList[0].Height)/2
  
  self.BGPosX = self.MyList[0].PosX - 10
  self.BGPosY = self.MyList[0].PosY - 10
  
  self.BGWidth = self.MyList[0].Width + 20
  self.BGHeight = self.MyList[0].Height +20  
}

func (self *ConfirmPage) SnapMsg(msg string) {
  self.MyList[0].SetText(msg)
  self.MyList[0].PosX = (self.Width - self.MyList[0].Width)/2
  self.MyList[0].PosY = (self.Width - self.MyList[0].Height)/2
  
  self.BGPosX = self.MyList[0].PosX - 10
  self.BGPosY = self.MyList[0].PosY - 10
  
  self.BGWidth = self.MyList[0].Width + 20
  self.BGHeight = self.MyList[0].Height +20
  
}

func (self *ConfirmPage) Init() {
  if self.Screen != nil {
  
    self.PosX = self.Index * self.Screen.Width
		self.Width = self.Screen.Width
		self.Height = self.Screen.Height
		self.CanvasHWND = self.Screen.CanvasHWND
    
    ps := NewListPageSelector()
    ps.Parent = self
    self.Ps = ps
    self.PsIndex = 0
    
    li := NewLabel()
    li.SetCanvasHWND(self.CanvasHWND)
    li.Init(self.ConfirmText,self.ListFont)
    
    li.PosX = (self.Width - li.Width)/2
    li.PosY = (self.Height - li.Height)/2
    
    self.BGPosX = li.PosX - 10
    self.BGPosY = li.PosY - 10
    self.BGWidth = li.Width + 20
    self.BGHeight = li.Height + 20
    
    self.MyList = append(self.MyList,li)
    
  }
}


func (self *ConfirmPage) KeyDown( ev *event.Event ) {

	if ev.Data["Key"] == UI.CurKeys["A"] || ev.Data["Key"] == UI.CurKeys["Menu"] {
		self.ReturnToUpLevelPage()
		self.Screen.Draw()
		self.Screen.SwapAndShow()
	}  
}

func (self *ConfirmPage) DrawBG() {
  rect_ := rect.Rect(self.BGPosX,self.BGPosY,self.BGWidth,self.BGHeight)
  
  draw.Rect(self.CanvasHWND,&color.Color{255,255,255,255}, &rect_, 0) // SkinManager().GiveColor('White')
  draw.Rect(self.CanvasHWND,&color.Color{83,83,83,255}, &rect_, 1)//SkinManager().GiveColor('Text')
}


func (self *ConfirmPage) Draw() {
  self.DrawBG()
  for _,v := range self.MyList{
    v.Draw()
  }  
  self.Reset()
}
