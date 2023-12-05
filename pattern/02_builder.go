/*
Паттерн "строитель" (builder) применяют:
- когда процесс конструирования объекта должен обеспечивать различные представления объекта
- когда мы хотим изолировать сложный процесс создания объекта от его представления
- 

Плюсы:
- разделение конструирования и представления
- позволяет создавать различные варианты объекта
Минусы:
- увеличение сложности кода
- неудобство при работе с простыми объектами

Примеры использования:
- строитель может использоваться для создания HTML-документа, где каждый строитель отвечает за построение определенной части документа
- строитель может быть использован для удобного конфигурирования объектов в тестах, где необходимо создавать объекты с разными наборами параметров
- в игровой разработке строитель может использоваться для создания сложных объектов, таких как персонажи, с различными характеристиками и оснащением
*/

package pattern

// Продукт, который мы строим
type Product struct {
    PartA string
    PartB string
    PartC string
}

// Интерфейс строителя
type Builder interface {
    BuildPartA()
    BuildPartB()
    BuildPartC()
    GetResult() Product
}

// Конкретный строитель
type ConcreteBuilder struct {
    product Product
}

func (b *ConcreteBuilder) BuildPartA() {
    b.product.PartA = "PartA"
}

func (b *ConcreteBuilder) BuildPartB() {
    b.product.PartB = "PartB"
}

func (b *ConcreteBuilder) BuildPartC() {
    b.product.PartC = "PartC"
}

func (b *ConcreteBuilder) GetResult() Product {
    return b.product
}

// Руководитель, управляющий процессом строительства
type Director struct {
    builder Builder
}

func (d *Director) Construct() Product {
    d.builder.BuildPartA()
    d.builder.BuildPartB()
    d.builder.BuildPartC()
    return d.builder.GetResult()
}

func main() {
	builder := &pattern.ConcreteBuilder{}
	director := &pattern.Director{builder: builder}

	product := director.Construct()

	fmt.Println("Product Parts:")
	fmt.Println("PartA:", product.PartA)
	fmt.Println("PartB:", product.PartB)
	fmt.Println("PartC:", product.PartC)
}