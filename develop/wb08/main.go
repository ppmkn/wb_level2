package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		// Выводим приглашение для пользователя
		fmt.Print("myshell> ")

		// Считываем команду из ввода пользователя
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Проверяем наличие команды выхода
		if strings.ToLower(input) == "\\quit" {
			break
		}

		// Разбиваем введенную строку на части
		args := strings.Fields(input)

		// Обработка команды cd
		if args[0] == "cd" {
			if len(args) < 2 {
				fmt.Println("Usage: cd <directory>")
			} else {
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
			continue
		}

		// Обработка команды pwd
		if args[0] == "pwd" {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(dir)
			}
			continue
		}

		// Обработка команды echo
		if args[0] == "echo" {
			fmt.Println(strings.Join(args[1:], " "))
			continue
		}

		// Обработка команды kill
		if args[0] == "kill" {
			if len(args) < 2 {
				fmt.Println("Usage: kill <process_id>")
			} else {
				pid := args[1]
				cmd := exec.Command("kill", pid)
				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
			continue
		}

		// Обработка команды ps
		if args[0] == "ps" {
			cmd := exec.Command("ps")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		}

		// Если не совпадает ни с одной из встроенных команд, выполняем внешнюю команду
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}