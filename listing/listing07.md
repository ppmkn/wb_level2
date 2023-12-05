Что выведет программа? Объяснить вывод программы.Что выведет программа? Объяснить вывод программы.

```go
package main
import (
    "fmt"
    "math/rand"
    "time"
)

func asChan(vs ...int) <-chan int {
    c := make(chan int)
    go func() {
        for _, v := range vs {
            c <- v
            time.Sleep(time.Duration(rand.Intn(1000)) *
            time.Millisecond)
        }
        close(c)
    }()
    return c
}
func merge(a, b <-chan int) <-chan int {
    c := make(chan int)
    go func() {
        for {
            select {
                case v := <-a:
                    c <- v
                case v := <-b:
                    c <- v
            }
        }
        }()
    return c
}
func main() {
    a := asChan(1, 3, 5, 7)
    b := asChan(2, 4 ,6, 8)
    c := merge(a, b )
    for v := range c {
        fmt.Println(v)
    }
}
```

**Ответ:**

`Будут выведены все значения из канала С, а затем будет бесконечно выводиться цифра 0`

После получения всех значений, каналы `a` и `b` закроются, но цикл `for` в функции `merge` будет продолжать работать. 

Так как считывание из закрытого канала возвращает `0`, то `0` будет выводиться **бесконечно**. Необходимо было проверять при считывании в функции `merge` на то, закрыты ли каналы или нет, чтобы исправить эту ошибку.