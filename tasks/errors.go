package tasks

import "errors"

var emptyTitle = errors.New("Ошибка: Пустое название задачи!")
var notFound = errors.New("Ошибка: Задачи с таким ID не существует!")
var failedOutput = errors.New("Ошибка: Произошла ошибка при получении списка задач!")
var alreadyDone = errors.New("Ошибка: Задача с таким ID уже выполнена!")
var failedToSave = errors.New("Ошибка: не удалось сохранить задачи в файл.")
var failedToLoad = errors.New("Ошибка: не удалось загрузить задачи из файла.")
