Что выведет программа? Объяснить вывод программы. Объяснить как работают
defer’ы и порядок их вызовов.

```go
package main
import (
  "fmt"
)

func test() (x int) {
  defer func() {
  x++
  }()
  x = 1
  return
}

func anotherTest() int {
  var x int
  defer func() {
  x++
  }()
  x = 1
  return x
}

func main() {
  fmt.Println(test())
  fmt.Println(anotherTest())
}
```

**Ответ:**

`2, 1`

В функции `test`, `defer` откладывает увеличение `x` после возвращения значения.
В функции `anotherTest`, `defer` тоже откладывает увеличение `x`, но после возвращения текущего значения.
Таким образом, `test` возвращает увеличенное значение `2`, а `anotherTest` возвращает исходное `1`.