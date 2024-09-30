# BISTRO FOOD 🍽️

## Описание

**BISTRO FOOD** — это мобильное приложение для онлайн-заказа еды в ресторане. Оно позволяет клиентам просматривать меню, добавлять блюда в корзину и оформлять заказы. В основе приложения лежит архитектура микросервисов с использованием **gRPC**, **API Gateway** и аутентификацией через **JWT-токены** для безопасности.

#### [Требования к проекту](docs/requirments.md)

## Основные возможности

- 📜 **Меню ресторана**: Просмотр актуального меню с фотографиями, описаниями и ценами.
- 🛒 **Корзина**: Добавление и удаление блюд в корзину, изменение количества.
- 📦 **Оформление заказа**: Клиенты могут оформлять заказы.
- 📊 **История заказов**: Пользователи могут просматривать историю своих заказов.
- 🔐 **Аутентификация через JWT**: Безопасный доступ к API с использованием JWT-токенов.
- 💬 **Отзывы и оценки**: Пользователи могут оставлять отзывы и оценки блюд.


## Технологии

- **Backend**:
    - **Go (Golang)**
    - **gRPC**
    - **JWT-токены**
    - **API Gateway**
    - **PostgreSQL, MongoDB**
    - **Docker и Docker Compose**
- **Frontend (Мобильное приложение)**:
    - **Android Studio** с **Java**
    - **Retrofit**
