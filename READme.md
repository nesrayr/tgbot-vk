# Телеграм бот - профильное задание для отбора на стажировку ВК. Ссылка на бота: https://t.me/passwordSavingBot

## Бот выполняет команды /set, /get и /del по условию задачи

## Бот написан на Golang. Была использована база данных PostgreSQL для хранения паролей и логинов для каждого пользователя. Для более удобного взаимодействия с БД использовал GORM. Пароли в базе хранятся в закодированном формате(base64) не более 3 суток. По истечении этого времени значения для этого сервиса удаляются

## Собрал приложение при помощи Docker и развернул в облаке
