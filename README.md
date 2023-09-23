**структура** 

`migrations` - скрипты для создания таблиц 

`pkg` - logger и методы для подключения к бд

`internal/entity` - модели

`config` - инициализация структуры config 

`internal/transcation` -  вся логрика для системы транзакций

`cmd/server` - файл main.go

Необходимо задать переменные среды:
DB_NAME  
DB_HOST 
DB_PORT 
DB_USER 
DB_PASSWORD 