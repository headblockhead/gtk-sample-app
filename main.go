package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
    // Initialize GTK without parsing any command line arguments.
    gtk.Init(nil)

    // Create a new toplevel window, set its title, and connect it to the
    // "destroy" signal to exit the GTK main loop when it is destroyed.
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unable to create window:", err)
    }
    win.SetTitle("Simple Example")
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })
    s, err :=gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
    if err != nil {
	    log.Fatal("no stack work :(", err);
    }

    // Create a new label widget to show in the window.
    l, err := gtk.LabelNew("Hello!")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }
    l_alt, err := gtk.LabelNew("asjndkjashjdk")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }
    tb, err := gtk.ButtonNewWithLabel("testbutton")
    if err != nil {
	    log.Fatal("no button work :(", err);
    }
    tb.Connect("enter", func() {
	    log.Println("hover detected!")
	    s.Add(l_alt)
	    l.Destroy()
	    win.ShowAll()
    })
    // Add the label to the window.
    s.Add(l)
    s.Add(tb)
    win.Add(s)

    // Set the default window size.
    win.SetDefaultSize(800, 600)

    // Recursively show all widgets contained in this window.
    win.ShowAll()

    // Begin executing the GTK main loop.  This blocks until
    // gtk.MainQuit() is run.
    gtk.Main()
}

