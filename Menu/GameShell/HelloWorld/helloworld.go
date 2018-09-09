package main

import (
	"github.com/veandco/go-sdl2/ttf"

	"github.com/cuu/gogame/surface"
	"github.com/cuu/gogame/event"
	"github.com/cuu/gogame/rect"
	
	"github.com/cuu/LauncherGo/sysgo/UI"
	
)

type InfoPageListItem struct{
	PosX   int
	PosY   int
	Width  int
	Height int
	Labels map[string]LabelInterface
	Icons  map[string]IconItemInterface
	Fonts  map[string]*ttf.Font

	Parent UI.PageInterface

	Flag   string
}

func NewInfoPageListItem() *InfoPageListItem {
	i := &InfoPageListItem{}
	i.Labels = make(map[string]LabelInterface)
	i.Icons = make( map[string]IconItemInterface)
	i.Fonts = make(map[string]*ttf.Font)

	i.Height = 20
	i.Width  = 0
	
	return i
}

func (self *InfoPageListItem) Init(text string) {
	l := NewLabel()
	l.PosX = 10
	l.SetCanvasHWND(self.Parent.GetCanvasHWND())
	l.Init(text,self.Fonts["normal"],nil)

	self.Labels["Text"] = l
	
}

func (self *InfoPageListItem) SetSmallText( text string) {
	l := NewMultiLabel()
	l.SetCanvasHWND(self.Parent.GetCanvasHWND())
	l.Init(text,self.Fonts["small"],nil)

	self.Labels["Small"] = l
	
	
}

func (self *InfoPageListItem) Draw() {
	x_,_ := self.Labels["Text"].Coord()
	self.Labels["Text"].NewCoord(x_,self.PosY)
	self.Labels["Text"].Draw()

	if val, ok := self.Labels["Small"]; ok {
		w_,_ := self.Labels["Text"].Size()
		self.Labels["Small"].NewCoord(w_+16,self.PosY)
		self.Labels["Small"].Draw()
	}
	
}


type HelloWorldPage struct {
	UI.Page
	AList map[string]map[string]string
	ListFontObj  *ttf.Font
	Scrolled int
	BGwidth int
	BGheight int
	DrawOnce bool
	Scroller *UI.ListScroller

	MyList []*InfoPageListItem
}

func NewHelloWorldPage() *HelloWorldPage {
	p := &HelloWorldPage{}
	
	p.FootMsg = [5]string{"Nav.","","","Back",""}

	p.AList = make(map[string]map[string]string)

	p.BGwidth = 320
	p.BGheight = 240-24-20
	p.DrawOnce = false

	p.MyList = make([]*InfoPageListItem,0)

	p.ListFontObj = UI.Fonts["varela13"]
	
	return p
}

func (self *HelloWorldPage) Init() {
	if self.Screen != nil {
		if self.Screen.CanvasHWND != nil && self.CanvasHWND == nil {
			self.HWND = self.Screen.CanvasHWND
			self.CanvasHWND = surface.Surface(self.Screen.Width,self.BGeight)
		}

		self.PosX = self.Index * self.Screen.Width
		self.Width = self.Screen.Width
		self.Height = self.Screen.Height



		self.HelloWorld()
		self.GenList()

		self.Scroller = UI.NewListScroller()
		
		self.Scroller.Parent = self
		self.Scroller.PosX = self.Width - 10
		self.Scroller.PosY = 2
		self.Scroller.Init()
		self.Scroller.SetCanvasHWND(self.HWND)
		
	}

}

func (self *HelloWorldPage) HelloWorld() {
	hello := make(map[string]string)
	hello["key"] = "helloworld"
	hello["label"] = "HelloWorld "
	hello["value"] = "GameShell"

	p.AList["hello"] = hello
	
}

func (self *HelloWorldPage) GenList() {
	self.MyList = nil
	self.MyList = make([]*InfoPageListItem,0)

	start_x := 0
	start_y := 0

	last_height := 0


	for i,u := range []string{"hello"} {
		if val,ok := self.AList[u];ok {

			li := NewInfoPageListItem()
			li.Parent = self
			li.PosX = start_x
			li.PosY = start_y + last_height
			li.Width = UI.Width
			li.Fonts["normal"] = self.ListFontObj
			li.Fonts["small"] = UI.Fonts["varela12"]

			if val["label"] != "" {
				li.Init(val["label"])
			}else{
				li.Init(val["key"])
			}

			li.Flag = val["key"]
			li.SetSmallText(val["value"])

			last_height += li.Height

			self.MyList = append(self.MyList,li)
			
		}else{
			continue
		}
	}	
}

func (self *HelloWorldPage) ScrollDown() {
	dis := 10
	if UI.abs(self.Scrolled) < ( self.BGheight - self.Height)/2 + 0 {
		self.PosY -= dis
		self.Scrolled -= dis
	}
}

func (self *HelloWorldPage) ScrollUp() {
	dis := 10
	if self.PosY < 0 {
		self.PosY += dis
		self.Scrolled += dis
	}
}

func (self *HelloWorldPage) OnLoadCb() {
	self.Scrolled = 0
	self.PosY     = 0
	self.DrawOnce = false
}

func (self *HelloWorldPage) OnReturnBackCb() {
	self.ReturnToUpLevelPage()
	self.Screen.Draw()
	self.Screen.SwapAndShow()
}


func (self *HelloWorldPage) KeyDown( ev *event.Event) {
	if ev.Data["Key"] == UI.CurKeys["A"] || ev.Data["Key"] == UI.CurKeys["Menu"] {
		self.ReturnToUpLevelPage()
		self.Screen.Draw()
		self.Screen.SwapAndShow()
	}

	if ev.Data["Key"] == UI.CurKeys["Up"] {
		self.ScrollUp()
		self.Screen.Draw()
		self.Screen.SwapAndShow()
	}

	if ev.Data["Key"] == UI.CurKeys["Down"] {
		self.ScrollDown()
		self.Screen.Draw()
		self.Screen.SwapAndShow()
	}	
	
}


func (self *HelloWorldPage) Draw() {
	if self.DrawOnce == false {

		self.ClearCanvas()

		for _,v := range self.MyList {
			v.Draw()
		}
		
		self.DrawOnce = true
	}

	if self.HWND != nil {
		surface.Fill(self.HWND, &color.Color{255,255,255,255})

		surface.Blit(self.HWND,self.CanvasHWND,&rect.Rect(self.PosX,self.PosY,self.Width,self.Height), nil)
		self.Scroller.UpdateSize(self.BGheight,UI.abs(self.Scrolled)*3)
		self.Scroller.Draw()
		
	}
}

/******************************************************************************/
type HelloWorldPlugin struct {
	UI.Plugin
	Page UI.PageInterface
}


func (self *HelloWorldPlugin) Init( main_screen *UI.MainScreen ) {
	self.Page = NewHelloWorldPage()
	self.Page.SetScreen( main_screen)
	self.Page.SetName("HelloWorld")
	self.Page.Init()
}

func (self *HelloWorldPlugin) Run( main_screen *UI.MainScreen ) {
	if main_screen != nil {
		main_screen.PushPage(self.Page)
		main_screen.Draw()
		main_screen.SwapAndShow()
	}
}

var APIOBJ HelloWorldPlugin





