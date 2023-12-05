/*
Паттерн "цепочка вызовов" (chain of resp) применяют:
- когда у есть несколько объектов, которые могут обработать определенный запрос, и порядок обработки не важен
- когда надо передать запрос нескольким объектам, но мы не знаем заранее, какие именно объекты могут это сделать
- когда набор объектов и их порядок в цепочке могут изменяться динамически во время выполнения программы

Плюсы:
- избавление от привязки отправителя запроса к получателю
- гибкость и расширяемость
Минусы:
- не гарантирует обработку запроса
- 

Примеры использования:
- каждый элемент интерфейса может быть обработчиком событий, и они могут быть организованы в цепочку для обработки событий в различных частях интерфейса
- веб-фреймворки часто используют паттерн "цепочка вызовов" для обработки запросов от клиентов
- различные обработчики могут добавлять логирование или аудит к запросам
*/

package pattern
import "fmt"

// Интерфейс обработчика
type Handler interface {
	HandleRequest(request int)
	SetSuccessor(successor Handler)
}

// Конкретный обработчик A
type ConcreteHandlerA struct {
	successor Handler
}

func (h *ConcreteHandlerA) HandleRequest(request int) {
	if request <= 10 {
		fmt.Println("ConcreteHandlerA handled the request")
	} else if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}

func (h *ConcreteHandlerA) SetSuccessor(successor Handler) {
	h.successor = successor
}

// Конкретный обработчик B
type ConcreteHandlerB struct {
	successor Handler
}

func (h *ConcreteHandlerB) HandleRequest(request int) {
	if request > 10 && request <= 20 {
		fmt.Println("ConcreteHandlerB handled the request")
	} else if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}

func (h *ConcreteHandlerB) SetSuccessor(successor Handler) {
	h.successor = successor
}

func main() {
	handlerA := &pattern.ConcreteHandlerA{}
	handlerB := &pattern.ConcreteHandlerB{}

	handlerA.SetSuccessor(handlerB)

	requests := []int{5, 15, 25}

	for _, request := range requests {
		handlerA.HandleRequest(request)
	}
}
