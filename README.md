# Posts_graphQLservice

Система для добавления и чтения постов и комментариев с использованием GraphQL
Для работы с GraphQL была исползована библиоткека [gqlgen](https://gqlgen.com/)

## Запуск

### Локальный запуск

1. Клонирование репозитория

   Запуск с postgre

2. `make docker.run.db` - запуск postgres
3. `make run.srv.postgre` - запуск сервиса

   Запуск с in memory

4. `make run.srv.memory` - запуск сервиса

### Запуск в Docker

1. Клонирование репозитория
2. `make run.all.docker` - запуск сервиса и postgres в Docker

### Остановка Docker

Для остановки и удаления контейнеров docker compose down
Для удаления вольюмов run.remove.volume

## Примеры запросов

### Создание поста

```graphql
mutation CreatePost {
  CreatePost(
    post: {
      name: "test post name"
      content: "test post conten"
      author: "test aurhor"
      commentsAllowed: true
    }
  ) {
    id
    name
    author
    content
  }
}
```

### Список постов

```graphql
query GetAllPosts {
  GetAllPosts {
    id
    name
    author
    content
  }
}
```

### Получекик поста с комментариями

```graphql
query GetPostById {
  GetPostById(id: 1) {
    id
    name
    author
    content
    commentsAllowed
    comments(limit: 2, offset: 0) {
      id
      author
      content
      post
      replies {
        id
        author
        content
        post
        replyTo
      }
    }
  }
}
```

### Создание комментария

```graphql
mutation CreateComment {
  CreateComment(
    input: {
      author: "test author"
      content: "test content"
      post: "1"
      replyTo: 1
    }
  ) {
    id
    author
    content
    post
    replyTo
  }
}
```

### Подписка на комментарии поста

```graphql
subscription CommentsSubscription {
  CommentSubscription(postId: 1) {
    author
    content
    post
    replyTo
  }
}
```
