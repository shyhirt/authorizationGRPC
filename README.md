# Сервис авторизации go-zero GRPC 
### Запуск без использования докера на хост машине Linux (Debian/Ubuntu)

## Установка зависимостей
#####  Mysql server, redis, etcd
<pre>
    sudo apt install golang mysql-server redis etcd git
</pre>
Создаем пользователя и БД, а также даем права на созданную БД Mysql
<pre>
    sudo mysql -u root
</pre>
Вводим пароль, после ввода пароля у нас появится консоль mysql>
<pre>
    CREATE DATABASE test; // Для тестов
    CREATE DATABASE dev; // Для разработки
    CREATE USER 'test'@'localhost' IDENTIFIED BY 'testPassword';
    GRANT ALL PRIVILEGES ON test . * TO 'test'@'localhost';
    GRANT ALL PRIVILEGES ON dev . * TO 'test'@'localhost';
</pre>
Клонируем репозиторий
<pre>
    git clone https://github.com/shyhirt/authorizationGRPC.git
    cd authorizationGRPC
</pre>

После чего развернем бекапы БД - для тестов структура+данные, для dev просто структура.
<pre>
    mysql -u test -p test < test/userTest.sql
    mysql -u test -p dev  < user/model/user.sql
</pre>
Генерируем JWT приватный и публичный ключи 
<pre>
    cd etc
    mkdir keys
    cd keys
    ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key
    openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
</pre>

Разберем файл конфига

<pre>
Name: user.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: user.rpc
DataSource: test:passwordtest@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true
Cache:
  - Host: 127.0.0.1:6379
    Pass:
    Type: node
jwt:
  AccessSecret: "etc/keys/jwtRS256.key"
  AccessPublic: "etc/keys/jwtRS256.key.pub"
  AccessExpire: 3600
  RefreshExpire: 345600
mail:
  Login: "smtpLogin"
  Identity: ""
  From: "test@gmail.com"
  Password: "secretPasswordSmtpService"
  SmtpHost: "smtp.elasticemail.com"
  SmtpPort: "2525"

</pre>
 - <b>Etcd</b> Настройки Etcd сервера
 - <b>DataSource:</b> DSN нашего mysql сервера
 - <b>Cache</b> Адрес Redis
 - <b>jwt</b> AccessSecret AccessPublic Пути к ключам для валидации и подписи JWT
 - <b>mail</b> Параметры SMTP сервера меняем на свои.
 - В репозитории оставлены сгенерированые RSA ключи и файлы конфигурации, но в реальной жизни так делать нельзя)))
 - <b>Использовать эти ключи уже где либо катерогически не рекомендуется.</b>

## Сборка
Вернемся в наш каталог с исходным кодом 
<pre>
    cd ../..
    go mod install
    go mod tidy
    go build 
    ./authorizationGRPC -f etc/user.yaml
</pre>

## Тестирование
<pre>
    ./authorizationGRPC -f etc/user-test.yaml &>/dev/null & cd test | go test ./test/...
</pre>

Убьем процесс
<pre>
    killall ./authorizationGRPC
</pre>

## Описание API GRPC


|            Функция            | Описание                                                                                                                                                                                                                                                                          |  Параметры  |                        Возвращает                    |
|:-----------------------------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------:|:----------------------------------------------------:|
|         Registration          | Функция выполняет регистрацию пользователя, перед этим осуществляя проверку email и username на наличие дубликатов. После успешной регистрации в отдельном потоке отправляется  email с кодом подтверждения на указанный почтовый адрес. <br> В случае неудачи возвращает ошибку. | <pre>message RegReq{<br>  string email = 1;<br>  string username = 2;<br>  string firstName = 3;<br>  string lastName = 4;<br>  string password = 5;<br>}<br></pre> | <pre>message RegResp{<br>  int64 id = 1;<br>}</pre>  |
|Login| Аутентификация пользователя по логину (email) и паролю, служит для получения access токена и refresh токена для дальнейшей авторизации пользователя. Возвращает ошибки в случае неуспешной аутентификации.                                                                        |<pre>message LoginReq{<br>  string login = 1;<br>  string password = 2;<br>}</pre>|<pre>message LoginResp{<br>  string accessToken = 1;<br>  string refresh = 2;<br>}<br></pre>|
|UserInfo| Возвращает текущего пользователя по переданному  access токену. В случае неудачи вернет ошибку.| <pre>message UserInfoReq{<br>  string accessToken = 1;<br>}</pre>|<pre>message UserInfoResp{<br>  int64 id = 1;<br>  string email = 2;<br>  string username = 3;<br>  string firstName = 4;<br>  string lastName = 5;<br>}</pre>|
|CheckVerificationCode| Проверка кода пользователя который мы прислали при регистрации на эмейл, в реальном не тестовом примере рационально еще добавить идентификатор пользователя. В данном случае этого было не указано, просто проверяется код в БД.| <pre>message VerificationCodeReq{<br>  int64 code = 1;<br>}<br></pre>| <pre><br>message VerificationCodeResp{<br>  bool result = 1;<br>}</pre>|

## Что дальше:
 - Makefile
 - Доработать Dockerfile