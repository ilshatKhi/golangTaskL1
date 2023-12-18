package main

import "fmt"

// Интерфейс цели
type RussianSocket interface {
	Connect()
}

// Структура, реализующая целевой интерфейс
type SimpleRussianSocket struct{}

func (s *SimpleRussianSocket) Connect() {
	fmt.Println("Подключение к розетке RussianSocket")
}

// Интерфейс адаптера
type ChineSocket interface {
	PlugIn()
}

// Структура, реализующая адаптер
type SocketAdapter struct {
	socket RussianSocket
}

func (a *SocketAdapter) PlugIn() {
	fmt.Println("Используем адаптер для подключения китайской вилки к российской розетке")
	a.socket.Connect()
}

// Функция создания адаптера
func NewAdapter(socket RussianSocket) ChineSocket {
	return &SocketAdapter{socket: socket}
}

/*func main() {
	simpleSocket := &SimpleRussianSocket{}
	adapter := NewAdapter(simpleSocket)
	adapter.PlugIn()
}
*/
