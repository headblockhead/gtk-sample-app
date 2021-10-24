package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Add buffers and symbol variables
	var buffer0 = []int{}
	var buffer_dots = []int{}
	buffer_dots = append(buffer_dots, 0)
	if buffer_dots[0] == 1 {
		log.Println("what.")
	}
	//    var buffer1 []int
	//    var operation = ""

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Calculator")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	// Create a vertical stack.
	sv, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		log.Fatal("Unable to create vertical stack:", err)
	}
	// Put horisontal stacks in the vertical stack.
	sh_0, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	if err != nil {
		log.Fatal("Unable to create horiontal stack 0:", err)
	}
	sh_1, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	if err != nil {
		log.Fatal("Unable to create horiontal stack 1:", err)
	}
	sh_2, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	if err != nil {
		log.Fatal("Unable to create horiontal stack 2:", err)
	}
	sh_3, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	if err != nil {
		log.Fatal("Unable to create horiontal stack 3:", err)
	}
	// Create a new label widget to show in the window.
	buffer_label, err := gtk.LabelNew(list_to_string(buffer0))
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	// Create the buttons
	button_0, err := gtk.ButtonNewWithLabel("0")
	if err != nil {
		log.Fatal("Unable to create button 0:", err)
	}
	button_0.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 0)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_1, err := gtk.ButtonNewWithLabel("1")
	if err != nil {
		log.Fatal("Unable to create button 1:", err)
	}
	button_1.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 1)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_2, err := gtk.ButtonNewWithLabel("2")
	if err != nil {
		log.Fatal("Unable to create button 2:", err)
	}
	button_2.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 2)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_3, err := gtk.ButtonNewWithLabel("3")
	if err != nil {
		log.Fatal("Unable to create button 3:", err)
	}
	button_3.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 3)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_4, err := gtk.ButtonNewWithLabel("4")
	if err != nil {
		log.Fatal("Unable to create button 4:", err)
	}
	button_4.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 4)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})

	button_5, err := gtk.ButtonNewWithLabel("5")
	if err != nil {
		log.Fatal("Unable to create button 5:", err)
	}
	button_5.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 5)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_6, err := gtk.ButtonNewWithLabel("6")
	if err != nil {
		log.Fatal("Unable to create button 6:", err)
	}
	button_6.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 6)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_7, err := gtk.ButtonNewWithLabel("7")
	if err != nil {
		log.Fatal("Unable to create button 7:", err)
	}
	button_7.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 7)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_8, err := gtk.ButtonNewWithLabel("8")
	if err != nil {
		log.Fatal("Unable to create button 8:", err)
	}
	button_8.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 8)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_9, err := gtk.ButtonNewWithLabel("9")
	if err != nil {
		log.Fatal("Unable to create button 9:", err)
	}
	button_9.Connect("clicked", func() {
		// Append the selected number
		buffer0 = append(buffer0, 9)
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	button_dot, err := gtk.ButtonNewWithLabel(".")
	if err != nil {
		log.Fatal("Unable to create button dot:", err)
	}
	button_dot.Connect("clicked", func() {
		// Append the selected number
		buffer_dots = append(buffer_dots, len(buffer0))
		// Set the label's contents
		buffer_label.SetLabel(list_to_string(buffer0))
		win.ShowAll()
	})
	// Add the widgets to the window.
	sv.Add(sh_0)
	sv.Add(sh_1)
	sv.Add(sh_2)
	sv.Add(sh_3)
	sv.Add(buffer_label)

	sh_0.Add(button_1)
	sh_0.Add(button_2)
	sh_0.Add(button_3)
	sh_1.Add(button_4)
	sh_1.Add(button_5)
	sh_1.Add(button_6)
	sh_2.Add(button_7)
	sh_2.Add(button_8)
	sh_2.Add(button_9)
	sh_3.Add(button_dot)
	sh_3.Add(button_0)
	//sh_3.Add(button_equals)
	win.Add(sv)

	// Set the default window size.
	win.SetDefaultSize(300, 200)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func list_to_string(values []int) string {
	valuesText := []string{}
	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	// Join our string slice.
	result := strings.Join(valuesText, "")
	return result
}
