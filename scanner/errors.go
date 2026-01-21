package scanner

import "errors"

var failedToScan error = errors.New("Ошибка программы! Повторите еще раз.")
var emptyInput error = errors.New("Ошибка: вы ничего не ввели!")
var incorretInput error = errors.New("Ошибка: некорректный ввод! Введите целое число, больше нуля.")
var IncorrectCMD error = errors.New("Ошибка: вы ввели неверную команду! | Чтобы узнать все команды введите help")
var exitCMD error = errors.New("exit command")
