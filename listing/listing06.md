Что выведет программа? Объяснить вывод программы. Рассказать про
внутреннее устройство слайсов и что происходит при передаче их в качестве
аргументов функции.

```go
package main
import (
    "fmt"
)

func main() {
    var s = []string{"1", "2", "3"}
    modifySlice(s)
    fmt.Println(s)
}

func modifySlice(i []string) {
    i[0] = "3"
    i = append(i, "4")
    i[1] = "5"
    i = append(i, "6")
}
```

**Ответ:**

`[3 2 3]`

**Слайс** - тип данных, который хранит в себе указатель на массив. Передав слайс в функцию, мы будем работать с его копией и изменения останутся лишь в копии слайса внутри данной функции.

Присвоение новых чисел затронуло оригинальный слайс, так как мы изменяем значение массива по **указателю**.