# User and Entity Behavior Analytics


## 🚀 Доступные ручки

1. **Вывод данных по id**
    
    - Метод: GET
    - Путь: `localhost:3003/get-items`
    - Описание: принимает идентификаторы записей ``id``, возвращает массив json’ов c данными по этим идентификаторам
      
## Примеры ответов
### 200 
![](https://github.com/faringet/UEBA/blob/master/screenshots/200.jpeg)

### 400 
![](https://github.com/faringet/UEBA/blob/master/screenshots/400.jpeg)

### 404 
![](https://github.com/faringet/UEBA/blob/master/screenshots/404.jpeg)



   

## ⚙️ Использованные технологии
Проект разработан с использованием следующих технологий:

- [**Gin**](https://github.com/gin-gonic/gin) - роутер
- [**zap**](https://github.com/uber-go/zap) - инструмент для эффективного логирования
- [**viper**](https://github.com/spf13/viper) - библиотека для конфигурирования приложения
