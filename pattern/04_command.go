/*
Паттерн "комманда" (command) применяют:
- когда необходимо параметризовать объекты операциями
- когда необходима очередь операций
- когда необходимо поддерживать отмену операций
- когда нужно поддерживать механизм регистрации запросов и их выполнение в разное время

Плюсы:
- отделение отправителя и получателя
- поддержка отмены и повтора операций
- легкость добавления новых команд
Минусы:
- увеличивается количество типов
- усложнение кода

Примеры использования:
- команды могут использоваться для обработки событий пользовательского интерфейса, таких как нажатия кнопок, изменения состояний элементов и т.д.
- команды могут представлять транзакции в базах данных, позволяя отменять или подтверждать изменения в базе данных
*/

package pattern
import "fmt"

// Интерфейс команды
type Command interface {
	Execute()
}

// Конкретная команда
type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Получатель команды
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: Action executed")
}

// Исполнитель команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &pattern.Receiver{}
	command := &pattern.ConcreteCommand{receiver: *receiver}
	invoker := &pattern.Invoker{}
	invoker.SetCommand(command)

	invoker.ExecuteCommand()
}
