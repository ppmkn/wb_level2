Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
    msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
    {
        // do something
    }
    return nil
}

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

**Ответ:**

`error`

Мы реализовали интерфейс `error` типом `customError`.

Функция `test()` возвращает тип `*customError`. Поэтому, при сравнении переменной `err` сравниваться будут следующие пары: `[*customError, nil]` и `[*customError, nil]`, что есть `true`, поэтому и выведется `error`.

Если же в функции `test()` возвращался бы тип `error`, то вывод был бы `ok`.