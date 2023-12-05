/*
Паттерн "фабричный метод" (factory method) применяют:
- когда создание объекта связано с какой-то конкретной логикой или контекстом

Плюсы:
- гибкость
- расширяемость
Минусы:
- сложность
- сложность в понимании

Примеры использования:
- фабричный метод используются для создания элементов интерфейса, таких как кнопки, текстовые поля и т.д.
- фабричный метод может использоваться для создания объектов доступа к данным, соответствующих конкретным СУБД
- 
*/

package pattern
import "fmt"

// Интерфейс продукта
type Product interface {
	Use()
}

// Конкретный продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProductA")
}

// Конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProductB")
}

// Интерфейс фабрики
type Creator interface {
	CreateProduct() Product
}

// Конкретная фабрика A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// Конкретная фабрика B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	creatorA := &pattern.ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	productA.Use()

	creatorB := &pattern.ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	productB.Use()
}

