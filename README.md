
# Depublic

Depublic is an online ticketing service. This repo contains a web service or backend that will be used by the front end.This is a project intended as a submission and also to practice my skills. In development I used Echo framework, GORM, Postgresql.


## API Reference
### Authentication
#### Register
```http
  POST /register
```
#### Login
```http
  use middleware.BasicAuth from echo framework
  POST /depublic/login
```

### Event
#### Get all event

```http
  GET /events
```

#### Get event by category

```http
  GET /events/category/${category}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `category`      | `string` | **Required**. category of item to fetch |

#### Get event by location

```http
  GET /events/location/${location}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `location`      | `string` | **Required**. location of item to fetch |

#### Search event

```http
  GET /events/search?event_name=
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `event_name`      | `string` | **Required**. event_name of item to fetch |


## Features

- Login and Register user (done)
- Search and Filter (done)
- Implement JWT for authentication and authorization (On Progress)
- User Profile (To Do)
- Transaction History (To Do)
- App Notification (To Do)


## Feedback

It is currently under development for many other features. If you have any feedback, please reach out to us at instagram @syahmalik_ or linkedin Malik Yus Fajriansyah

