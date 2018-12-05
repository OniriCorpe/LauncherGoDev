package main

import (
	"github.com/veandco/go-sdl2/ttf"

  "path/filepath"
//	"github.com/cuu/gogame/surface"
	"github.com/cuu/gogame/event"
	"github.com/cuu/gogame/rect"
	"github.com/cuu/gogame/color"
	"github.com/cuu/gogame/draw"
	
	"github.com/cuu/LauncherGoDev/sysgo/UI"
  
  "github.com/cuu/LauncherGoDev/Menu/GameShell/10_Settings/LauncherPy"
  
	
)

type SettingsPageSelector struct {
	UI.PageSelector
	BackgroundColor *color.Color
	
}

func NewSettingsPageSelector() *SettingsPageSelector{
	s := &SettingsPageSelector{}
	s.BackgroundColor = &color.Color{131,199,219,255}

	s.Width = UI.Width
	
	return s
}

func (self *SettingsPageSelector) Draw() {
	idx := self.Parent.GetPsIndex()
	mylist := self.Parent.GetMyList()
	if idx < len( mylist) {
		_,y_ := mylist[idx].Coord()
		_,h_  := mylist[idx].Size()
		
		x := 2
		y := y_+1
		h := h_-3
		self.PosX = x
		self.PosY = y
		self.Height = h

		rect_ := rect.Rect(x,y,self.Width-4,h)
		canvas_ := self.Parent.GetCanvasHWND()
		draw.AARoundRect(canvas_, &rect_,self.BackgroundColor,4,0,self.BackgroundColor)
		
	}
}


type SettingPlugin struct{
  Type int
  SoFile string
  FolderName string
  LabelText  string
  EmbInterface  UI.PluginInterface
}
//##############################################//

type SettingsPage struct {
	UI.Page
	AList map[string]map[string]string
	ListFontObj  *ttf.Font
	Scrolled int
	BGwidth int
	BGheight int
	DrawOnce bool
	Scroller *UI.ListScroller
	Icons map[string]UI.IconItemInterface

	MyPath string
	
}

func NewSettingsPage() *SettingsPage {
	p := &SettingsPage{}
	p.FootMsg = [5]string{"Nav","","","Back","Enter"}
	p.ListFontObj = UI.Fonts["varela15"]

	p.MyPath = "Menu/GameShell/10_Settings"
	
	return p
}

func (self *SettingsPage) GenList() []*SettingPlugin {
  alist := []*SettingPlugin{
    
    &SettingPlugin{0,"wifi.so",  "Wifi",      "Wi-Fi",nil},
    
    &SettingPlugin{0,"sound.so",      "Sound",      "Sound Volume" ,         nil},
    &SettingPlugin{0,"brightness.so", "Brightness", "BackLight Brightness", nil},
    &SettingPlugin{0,"update.so",     "Update",     "Update",               nil},
    &SettingPlugin{0,"about.so",      "About",      "About",                nil},
    
    &SettingPlugin{1,"",         "LauncherPy","Switch to Launcher",&LauncherPy.APIOBJ},
    
  }
  
  return alist
}

func (self *SettingsPage) Init() {
	if self.Screen != nil {
		
		self.PosX = self.Index * self.Screen.Width
		self.Width = self.Screen.Width
		self.Height = self.Screen.Height
		self.CanvasHWND = self.Screen.CanvasHWND


		ps := NewSettingsPageSelector()
		ps.Parent = self
		self.Ps = ps
		self.PsIndex = 0
		
    
		start_x := 0
		start_y := 0
    
    alist := self.GenList()
    
		for i,v := range alist{
			li := UI.NewListItem()
			li.Parent = self
			li.PosX   = start_x
			li.PosY   = start_y + i*li.Height
			li.Width  = UI.Width

			li.Fonts["normal"] = self.ListFontObj

			if v.LabelText != "" {
				li.Init(v.LabelText)
			}else{
				li.Init(v.FolderName)
			}
			
			if v.SoFile!= "" && UI.FileExists( filepath.Join(self.MyPath,v.FolderName,v.SoFile )) {
				pi,err := UI.LoadPlugin(filepath.Join(self.MyPath,v.FolderName,v.SoFile ))
				UI.Assert(err)
				li.LinkObj  = UI.InitPlugin(pi,self.Screen)
				self.MyList = append(self.MyList,li)
				
			}else {
        if v.EmbInterface != nil {
          v.EmbInterface.Init(self.Screen)
          li.LinkObj = v.EmbInterface
          self.MyList = append(self.MyList,li)
        }
      }
		}

		self.Scroller = UI.NewListScroller()
		self.Scroller.Parent = self
		self.Scroller.PosX  = self.Width - 10
		self.Scroller.PosY  = 2
		self.Scroller.Init()

	}
}

func (self *SettingsPage) ScrollUp() {
	if len(self.MyList) == 0 {
    return
  }
  
  self.PsIndex -= 1
  
  if self.PsIndex < 0 {
    self.PsIndex = 0
  }
  cur_li := self.MyList[self.PsIndex]
  x,y := cur_li.Coord()
  if x < 0 {
    for i:=0;i<len(self.MyList);i++ {
      _,h := self.MyList[i].Size()
      x,y  = self.MyList[i].Coord()
      self.MyList[i].NewCoord(x, y+h)
    }
  }
}


func (self *SettingsPage) ScrollDown() {
  if len(self.MyList) == 0 {
    return
  }
  
  self.PsIndex += 1
  if self.PsIndex >= len(self.MyList) {
    self.PsIndex = len(self.MyList) - 1
  }
  
  cur_li := self.MyList[self.PsIndex]
  x,y  := cur_li.Coord()
  _,h  := cur_li.Size()
  
  if y + h > self.Height {
    for i:=0;i<len(self.MyList);i++ {
      _,h = self.MyList[i].Size()
      x,y = self.MyList[i].Coord()
      self.MyList[i].NewCoord(x, y - h)
    }
  }
  
}

func (self *SettingsPage) Click() {
	if len(self.MyList) == 0 {
		return
	}
	
	cur_li := self.MyList[self.PsIndex]

	lk_obj := cur_li.GetLinkObj()

	if lk_obj != nil {
		lk_obj.Run(self.Screen)
	}

}

func (self *SettingsPage) KeyDown( ev *event.Event) {
	
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

	if ev.Data["Key"] == UI.CurKeys["Enter"] {
		self.Click()
	}
}


func (self *SettingsPage) Draw() {
	self.ClearCanvas()
	
	if len(self.MyList) == 0 {
		return
	}

	_,h_ := self.MyList[0].Size()

	if len(self.MyList) * h_ > self.Height {
		_,ph_ := self.Ps.Size()
		self.Ps.NewSize(self.Width - 11, ph_)
		self.Ps.Draw()

		for _,v := range self.MyList {
			v.Draw()
		}

		self.Scroller.UpdateSize(len(self.MyList)*h_,self.PsIndex*h_)
		self.Scroller.Draw()
		
	}else {
		_,ph_ := self.Ps.Size()
		self.Ps.NewSize(self.Width,ph_)
		self.Ps.Draw()
		for _,v := range self.MyList {
			v.Draw()
		}
		
	}
}


/******************************************************************************/
type SettingsPlugin struct {
	UI.Plugin
	Page UI.PageInterface
}


func (self *SettingsPlugin) Init( main_screen *UI.MainScreen ) {
	self.Page = NewSettingsPage()
	self.Page.SetScreen( main_screen)
	self.Page.SetName("Settings")
	self.Page.Init()
}

func (self *SettingsPlugin) Run( main_screen *UI.MainScreen ) {
	if main_screen != nil {
		main_screen.PushPage(self.Page)
		main_screen.Draw()
		main_screen.SwapAndShow()
	}
}

var APIOBJ SettingsPlugin
