package main

import (
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows" //Windows specific system clas
)

func main() {
	/*
		Create a windows message box using winAPI
		hWnd represents the handle to the parent window
		Using 0 means no parent window(message box is a standalone)
	*/
	hWnd := uintptr(0) //a HANDLE in C++

	//Call the windows Message box function
	windows.MessageBox(
		windows.HWND(hWnd), //Handle to our messagebox
		windows.StringToUTF16Ptr("Used windows package"), //main content displayed to the box
		windows.StringToUTF16Ptr("MessageBox 1/2"),       //displayed in the title bar of the message box
		windows.MB_OK) //Creates a simple OK button

	/*
		THe above used  a Go Friendly approach
		===Syscall Package Below:
	*/
	//Load the user32.dll library dynamically at runtime
	//User32.dll Conatains the MessageBoxW functions we need
	user32dll := syscall.NewLazyDLL("User32.dll")

	//Get a reference to the specific MessagBoxW function from the DLL
	//The "W" suffix indicates the Unicode/wide-chars version
	procMsgBox := user32dll.NewProc("MessageBoxW")

	//we set hWnd to 0, no parent window for the second message box
	hWnd = uintptr(0)

	//Convert Go strings to UTF-16 format for the message box
	IpText, err := syscall.UTF16PtrFromString("Used Syscall Package")

	if err != nil {
		log.Fatalln("IpText UTF16PtrFromString Failed") //Exit if conversion fails

	}

	//Convert Go string to windows format for title
	IpCaption, err := syscall.UTF16PtrFromString("MessageBox 2/2")
	if err != nil {
		log.Fatalln("IpCaption UTF16PtrFromString Failed") //exit if conversion fails
	}
	//uType defines the buttons and icons to display
	//0= MB_OK(OK button only, no icon)
	uType := uint(0)

	//Manually call the messageBoxW function using syscall
	//unsafe.Pointer is needed because we're passing to C-sty;e strings
	procMsgBox.Call(
		hWnd,                               //parent window
		uintptr(unsafe.Pointer(IpText)),    //Message text converted to pointer
		uintptr(unsafe.Pointer(IpCaption)), //Title text
		uintptr(uType))                     //Button/style flags
}
