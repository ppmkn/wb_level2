/*
Паттерн "состояние" (state) применяют:
- когда объект может менять свое поведение в зависимости от своего внутреннего состояния
- когда операции, которые должны быть выполнены объектом, зависят от его состояния
- когда используется большой оператор switch или if-else для управления состоянием объекта

Плюсы:
- избегание больших операторов switch или if-else
- легкость добавления новых состояний
Минусы:
- величение числа типов

Примеры использования:
- объекты, представляющие персонажей или объекты в играх, могут использовать паттерн "состояние" для управления их поведением в различных ситуациях (например, состояния атаки, обороны, отдыха и т. д.)
- состояние заказа может меняться от "в обработке" до "отправлен" и "доставлен"
- состояние интерфейса пользователя (например, ввод данных, отправка формы, успешное завершение операции) может быть эффективно управляемо паттерном "состояние"
*/

package pattern
import "fmt"

// Интерфейс состояния
type State interface {
	Handle()
}

// Конкретное состояние A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Handling ConcreteStateA")
}

// Конкретное состояние B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Handling ConcreteStateB")
}

// Контекст
type Context struct {
	state State
}

// Установка состояния
func (c *Context) SetState(state State) {
	c.state = state
}

// Обработка состояния
func (c *Context) Request() {
	c.state.Handle()
}

func main() {
	context := &pattern.Context{}

	stateA := &pattern.ConcreteStateA{}
	context.SetState(stateA)
	context.Request()

	stateB := &pattern.ConcreteStateB{}
	context.SetState(stateB)
	context.Request()
}