# Notes REST API with Yandex Speller integration (KODETestCase)
<h1>Запуск:</h1>
<ol>
  <li>Клонировать репозиторий
    
  ```console
    git clone https://github.com/Serveresty/Zametki.git
  ```
  </li>
  <li>Создать Docker контейнер с указанием файла переменных окружения

  ```console
    docker-compose --env-file configs/.env up --build
  ```
  </li>
  <li>Зарегистрироваться

   ```console
    curl http://localhost:8000/sign-up -H "Content-Type: application/json" -d '{"first_name": "YourFirstName", "last_name": "YourLastName", "email": "your.email@mail.ru", "password": "qwerty"}'
   ```
  </li>
  <li>Авторизоваться(После успешной авторизации скопировать JWT Bearer token)
    
   ```console
    curl http://localhost:8000/sign-in -H "Content-Type: application/json" -d '{"email": "your.email@mail.ru", "password": "qwerty"}'
   ```
  </li>
  <li>Создать заметку (приложив JWT токен, который был получен после авторизации)

   ```console
    curl http://localhost:8000/create-notes -H "Content-Type: application/json" -H "Authorization: <PASTE YOUR TOKEN HERE>" -d '{"title": "Test title", "content": "I went to the libary to borow some books. Эта кошка живёт у миня дома."}'
   ```
  </li>
  <li>Получить список заметок (приложив JWT токен, который был получен после авторизации)

   ```console
    curl http://localhost:8000/notes -H "Content-Type: application/json" -H "Authorization: <PASTE YOUR TOKEN HERE>"
   ```
  </li>
</ol>
