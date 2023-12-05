/*
Паттерн "стратегия" (strategy) применяют:
- когда необходимо избегать использования наследования и создания подклассов для изменения поведения
- когда есть несколько вариантов алгоритмов и они могут изменяться во время выполнения

Плюсы:
- избегание лишнего наследования
- обмен и переключение стратегий
Минусы:
- усложнение кода клиента
- увеличение числа типов

Примеры использования:
- может быть использован для реализации различных стратегий сортировки (например, сортировка по возрастанию, по убыванию, по алфавиту и т. д.)
- различные алгоритмы сжатия данных (например, zip) могут быть реализованы с использованием стратегий
- стратегия может быть использована для реализации различных алгоритмов поиска (например, поиск в глубину, поиск в ширину, двоичный поиск и т. д.)
*/

package pattern
import "fmt"

// Интерфейс стратегии
type Strategy interface {
	Execute()
}

// Конкретная стратегия A
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute() {
	fmt.Println("Executing ConcreteStrategyA")
}

// Конкретная стратегия B
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute() {
	fmt.Println("Executing ConcreteStrategyB")
}

// Контекст
type Context struct {
	strategy Strategy
}

// Установка стратегии
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// Выполнение стратегии
func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	context := &pattern.Context{}

	strategyA := &pattern.ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.ExecuteStrategy()

	strategyB := &pattern.ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.ExecuteStrategy()
}