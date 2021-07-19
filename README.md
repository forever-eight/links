# links

Привет!
Нужно запустить docker-compose.yml, после чего поддтянется базка, создаем базу links, после чего выполняем миграции (вставляем в терминал все из migration.sql). Чтобы добавить что-то в бд и, соотвественно, получить короткую ссулку, нужно сделать такой curl --location --request POST 'http://localhost:1323/links' \
--header 'Content-Type: application/json' \
--data-raw '{
"link" : "https://www.google.com/search?client=safari&rls=en&q=hump+%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B4&ie=UTF-8&oe=UTF-8"
}'

