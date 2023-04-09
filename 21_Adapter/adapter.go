// 21. Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Mac USB - квадратый, Windows - круглый
// У клиента есть только квадратный - соответсвенно можно вставить его только в мак
// Нам нужен windows-адаптер
// У адаптера мы реализуем метод с квадратным USB
// Таким образом клиент сможет использовать InsertSquareUSBInComputer и для мака, и для винды

type Computer interface {
	InsertSquarePort()
}

type Mac struct {
}

func (m *Mac) InsertSquarePort() {
	fmt.Println("Inserting square port into computer...")
}

type Windows struct {
}

func (w *Windows) InsertCirclePort() {
	fmt.Println("Inserting circle port into computer...")
}

type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (w *WindowsAdapter) InsertSquarePort() {
	w.WindowsMachine.InsertCirclePort()
}

type Client struct {
}

func (c *Client) InsertSquareUSBInComputer(com Computer) {
	com.InsertSquarePort()
}

func main() {
	client := &Client{}
	mac := &Mac{}
	client.InsertSquareUSBInComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{ // создаём адаптер с машиной виндоус
		WindowsMachine: windowsMachine,
	}
	client.InsertSquareUSBInComputer(windowsMachineAdapter) // теперь через него мы можем использовать квадратый usb

}
