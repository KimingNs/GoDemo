package main

import (
	"fmt"
)

func init() {
	fmt.Println("---------begin---------")
}

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("phone starting...")
}
func (p Phone) Stop() {
	fmt.Println("phone stopping...")
}
func (p Phone) Call() {
	fmt.Println("phone calling...")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("camera starting...")
}

func (c Camera) Stop() {
	fmt.Println("camera stopping...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	//类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}
func main() {
	var usbArr [4]Usb
	usbArr[0] = Phone{Name: "phone1"}
	usbArr[1] = Camera{Name: "camera1"}
	usbArr[2] = Camera{Name: "camera2"}
	usbArr[3] = Phone{Name: "phone2"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
