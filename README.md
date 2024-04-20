
# Depublic

The project entails the development of a backend service designed to facilitate
the purchase of event tickets. Primarily built using Golang, with the Echo
framework, and leveraging the GORM Object-Relational Mapping (ORM) for
database interactions, the system aims to provide a seamless experience for
users in procuring tickets for various events, And still development.


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

- JWT-based Authentication: Users can securely log in and register through the system using JSON Web Tokens (JWT). This ensures authentication and authorization for accessing various functionalities.
- Event Ticket Booking: The platform enables users to browse, select, and purchase tickets for a wide array of events. Through intuitive interfaces, users can seamlessly navigate the available options and secure their desired tickets.
- Scalable and Efficient Backend: Leveraging Golang's robustness and efficiency, combined with the Echo framework's lightweight yet powerful features, the backend service ensures scalability and responsiveness even under high loads.
- ORM with GORM: Database interactions are streamlined and simplified with the utilization of GORM ORM, offering an intuitive and efficient way to map Go structs to database tables and perform CRUD operations.
- PostgreSQL Database: The project employs PostgreSQL as the backend database, ensuring data integrity, scalability, and flexibility for storing various aspects of event details, user information, and transaction records.

## Feedback

It is currently under development for many other features. If you have any feedback, please reach out to us at instagram @syahmalik_ or linkedin Malik Yus Fajriansyah

