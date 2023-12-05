/*
Паттерн "посетитель" (visitor) применяют:
- когда есть сложная структура объектов, и надо выполнить различные операции над этой структурой без изменения самих объектов
- когда типы объектов содержат стабильное интерфейсное поведение, но могут иметь изменяющееся во времени поведение
- 

Плюсы:
- отделение алгоритмов от структуры объектов
- легкость добавления новых операций
Минусы:
- сложность добавления новых классов объектов
- нарушение инкапсуляции

Примеры использования:
- если у есть иерархия элементов документа (например, текст, изображение, таблица), можно использовать паттерн "посетитель" для реализации операций экспорта в различные форматы (PDF, HTML, JSON)
- в компиляторах или интерпретаторах можно использовать посетителей для выполнения различных операций (например, оптимизации, генерации кода) над узлами AST без изменения структуры дерева
- 
*/

package pattern

// Интерфейс посетителя
type Visitor interface {
    VisitElementA(element ElementA)
    VisitElementB(element ElementB)
}

// Интерфейс элемента
type Element interface {
    Accept(visitor Visitor)
}

// Конкретный элемент A
type ElementA struct{}

func (e ElementA) Accept(visitor Visitor) {
    visitor.VisitElementA(e)
}

// Конкретный элемент B
type ElementB struct{}

func (e ElementB) Accept(visitor Visitor) {
    visitor.VisitElementB(e)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (v ConcreteVisitor) VisitElementA(element ElementA) {
    // Реализация операции для ElementA
}

func (v ConcreteVisitor) VisitElementB(element ElementB) {
    // Реализация операции для ElementB
}

func main() {
    // Создаем элементы
    elementA := &pattern.ElementA{}
    elementB := &pattern.ElementB{}

    // Создаем посетителя
    visitor := &pattern.ConcreteVisitor{}

    // Применяем посетителя к элементам
    elementA.Accept(visitor)
    elementB.Accept(visitor)
}
