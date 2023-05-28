
# Task Manager Backend GoLang

The task manager allows you to perform basic operations such as creating, reading, updating, deleting, and changing the status of tasks. The goal of this exercise is to improve your understanding and proficiency in GoLang backend development by implementing these functionalities.



## API Reference

#### Create Task

```http
  POST /create
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `title` | `string` | **Required**. Your Task Title |
| `description` | `string` | **Required**. Your Task Description |

#### Get All Tasks

```http
  GET /get
```

#### Get Task By Id

```http
  GET /get/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of task to fetch |

#### Delete All Tasks

```http
  DELETE /delete
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of task to fetch |

#### Delete Task By Id

```http
  DELETE /delete/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of task to delete |


#### Mark Task As Completed

```http
  PUT /complete/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of task to mark as completed |

#### Update Task By Id

```http
  PUT /update/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of task to update |

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. New Title |
| `description`      | `string` | **Required**. New Description |

## Authors

- [@prasoonsoni](https://www.github.com/prasoonsoni)

