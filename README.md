# Golang-MongoDB
This project is an example of data persistence that makes a request to an external API and records the data in MongoDB. This data can be consulted from two Get routes.

### Instalation and Testing
To install the service you will need to have DB Mongo installed and running on your machine or remotely and create the database and collections mentioned below.

If you don't have mongo installed follow these steps: https://www.mongodb.com/docs/manual/administration/install-community/

### Go Version on the project was made:
```
V1.16
```

Runing project:
```
go mod vendor
```
```
go run main.go
```
Create database in mongo with collection:
```
db = sport
collection = news_information
```

Testing Route:

```
http://localhost:8080/v1/apipopulate-test
```

Get Many News:

```
http://localhost:8080/v1/api/news
```

Get One News:

```
http://localhost:8080/v1/api/news/:id
```

### Create the .env in the project root with the key:

```
MONGODB_URI="mongodb://{host}:{port}"
```

#### Run Populate ScheduleJob:


```
 go run main.go populate-news
```

