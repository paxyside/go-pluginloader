# Plugin loader

## Тестовое задание:
Надо написать лоадер плагинов, которые реализуют какой-либо интерфейс.  
Притом чтобы загрузка и выгрузка была динамичной.  
У каждого плагина должно быть уникальное имя.  
Делаешь веб сервер, с 3 эндпоинтами:
1. Список плагинов
2. Загрузить плагин в память и вернуть имя его
3. Выгрузить плагин.

## Запуск и установка зависимостей

`git clone &&
cd --PluginsLoader &&
go mod download &&
go run cmd/main.go`

## Работа сервера
Работа сервера описана на трех эндпоинтах:\
**"/plugins/list"** для отображения всех загруженных плагинов\
**"/plugins/load"** для загрузки нового плагина (_пример: "/plugins/load?path=<path/to/plugin>"_)\
**"/plugins/unload"** для выгрузки плагинов (_пример: "/plugins/unload?name=<PluginName>"_)

## Компиляция новых плагинов для их дальнейшей загрузки/выгрузки в проект
Плагины хранятся в директории **./plugins**. Для их компиляции необходимо выполнить команду\
`go build -o <plugin_name> -buildmode=plugin path/to/plugin`
_Пример:_\
`go build -o plugin1.so -buildmode=plugin plugins/plugin1/plugin1.go`  

## Тестирование
Функции pluginloader покрыты тестами, для успешного тестирования необходимо скомпилировать тестовый плагин:\
`go build -o internal/test/test_plugin.so -buildmode=plugin internal/test/test_plugin.go`\
Для проверки:\
`go test ./internal/test`\
Для веб-сервера так же описан тестовый запрос:\
`go test ./api/test`

## Структура проекта

PluginsLoader\
├── api\
│   ├── server\
│   │   └── server.go\
│   └── test\
│       └── test_server.go\
├── cmd\
│   └── main.go\
├── internal\
│   ├── loader\
│   │   ├── loader.go\
│   │   └── plugin_loader.go\
│   │     \
│   └── test\
│       ├── loader_test.go\
│       └── test_plugin.go\
└── plugins\
    └── plugin1\
        └── plugin1.go


