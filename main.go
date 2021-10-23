package main

// import fyne
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"os/exec"

)
var btn fyne.Widget;
func main() {
    // New app
    a := app.New()
    // New title
    w := a.NewWindow("file Handling - Open Images")
    //resize
    w.Resize(fyne.NewSize(400, 400))
    btn = widget.NewButton("Open .jpg & .Png", func() {
        // dialog for opening files
        // 2 arguments
        fileDialog := dialog.NewFileOpen(
            // _ to ignore error
            func(uc fyne.URIReadCloser, _ error) {
                // reader to read data
                data, _ := ioutil.ReadAll(uc)
                // static resource
                // 2 arguments
                // first is file name (string)
                // second is data from reader
                // res := fyne.NewStaticResource(uc.URI().Name(), data)
                // Now image widget to display our image
                // img := canvas.NewImageFromResource(res)
                // setup new window for image and set content
                // w := fyne.CurrentApp().NewWindow(uc.URI().Name())
                // w.SetContent()
				type App struct {
					Name string
					Version int
					Script string
				  }
			  
				  // json data
				  var obj App
			  
				  // unmarshall it
				  err := json.Unmarshal(data, &obj)
				  if err != nil {
					  fmt.Println("error:", err)
				  }
			  
				  // can access using struct now
				fmt.Println(obj.Name);
				// if(obj.Name == "calculator"){
				// 	w.SetContent(container.NewVBox(btn,))
				// }
				var btn1 fyne.Widget = widget.NewButton(obj.Name,func() {
					cmd:=exec.Command("ls")
					err:=cmd.Run()
					if err!=nil{
						fmt.Println("Error",err)
					}
				})
                // resize window
				w.SetContent(container.NewVBox(btn,btn1));
                w.Resize(fyne.NewSize(400, 400))
                w.Show() // display our image
            }, w)
        // filtering files
        fileDialog.SetFilter(
            // filter jpg and png
            // ignore rest of the files
            storage.NewExtensionFileFilter([]string{".json",}))
        fileDialog.Show()
        // we are done 🙂
    })
    // display button in parent window
    w.SetContent(container.NewVBox(btn))
    w.ShowAndRun()
}