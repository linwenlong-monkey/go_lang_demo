package main

//接口
type Usbering interface {
    start()
    stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	println(p.Name,"Phone start")
}
func (p Phone) stop() {
	println(p.Name,"Phone stop")
}

func main() {
	phone := Phone{"HuaWei"}
	phone.start()
	phone.stop()
	
	var usb Usbering = phone
	usb.start()
	usb.stop()
}
