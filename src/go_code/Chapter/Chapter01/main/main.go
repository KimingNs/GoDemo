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
type Usb2 interface {
	Start()
	Stop()
}
type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机starting...")
}

func (p Phone) Stop() {
	fmt.Println("手机stopping...")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机starting...")
}

func (c Camera) Stop() {
	fmt.Println("相机stopping...")
}

type Computer struct {
}

func (pc Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
func main() {
	pc := Computer{}
	phone := Phone{}
	camera := Camera{}

	pc.Working(phone)
	pc.Working(camera)
}
