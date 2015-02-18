package main
//import "github.com/jbenet/go-ipfs/core"
import "github.com/andlabs/ui"

var window ui.Window


func main() {
	
    createUI()
	
}


func createUI() {
	
	
    go ui.Do(func() {
        name := ui.NewTextField()
        button := ui.NewButton("Greet")
        greeting := ui.NewLabel("")
        stack := ui.NewVerticalStack(
            ui.NewLabel("Enter your username:"),
            name,
            button,
            greeting)
        window = ui.NewWindow("Hello", 1000, 500, stack)
        button.OnClicked(func() {
            greeting.SetText("Hello, " + name.Text() + "!")
        })
        window.OnClosing(func() bool {
            ui.Stop()
            return true
        })
        window.Show()
    })
    err := ui.Go()
    if err != nil {
        panic(err)
    }

}

// func initIPFSNode() {
// 	//initialize IPFS node
// 	fmt.Printf("Hello, world.\n")
// 	builder := core.NewNodeBuilder()
// 	node, err := builder.Build(ctx)
// 	fmt.Printf("Hello, world. and again after node initialized\n")
// }