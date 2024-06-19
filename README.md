# Bonds app

## How to run the app

This project has been created with Golang in the backend and Next.js in the frontend.
First of all you need to clone the repository in your computer, you can do it with the following command:

```bash
git clone url_of_this_repository
```

## frontend
when you have the repository in your computer you need to go to the frontend folder and run the following command:

```bash
yarn install
```
##### IMPORTANT NOTE: you must have installed Node.js and yarn in your computer to run the previous command.

After that you can run the following command to start the frontend server:

```bash
yarn dev
```
and the frontend server will start in the following url: http://localhost:3000

## backend
To run the backend server, you must have Golang in your computer, and have the utility [golang-migrate](https://github.com/golang-migrate/migrate) installed in your computer.
Also you need to have a PostgreSQL database running in your computer.

When you have Postgres running in your computer, you need to create a database called `bonds` (or any other name you want),  You need username with enough permissions to create tables and insert data.
When you have your database created, and you database users, run the following command to set up the `DNS` of the database in the backend server:
```bash
DNS="postgres://<your_username>>:<password>@localhost/database_name?sslmode=disable"
```
And finally to run the migration you run the following command:
```bash
 migrate -path="./migrations" -database="$DSN" up
```
#### IMPORTANT NOTE: you must be in the backend folder to run the previous command.

if there are no errors, you can run the following command to start the backend server:
```bash
go run ./api/
```

if there are no errors, you can see the next message in the terminal:
```bash
time=2024-06-18T16:18:29.132-06:00 level=INFO msg="database connection pool established"
time=2024-06-18T16:18:29.132-06:00 level=INFO msg="Starting server" addr=4000
```
## options
when you are running the backend server you can pass the following flags to the server:
```bash
go run ./api/ -port=4000 -dsn="postgres://<your_username>>:<password>@localhost/database_name?sslmode=disable" --secret="your_secret"
```
where:
- port: is the port where the server will run
- dsn: is the `DSN` of the database
- secret: is the secret key that will be used to sign the tokens of the users

you can see the available flags with the following command:
```bash
go run ./api/ -h
```


## How to use the app
When you have the frontend and backend servers running, you can go to the following url: http://localhost:3000 and you will see the following screen:

## API documentation

the backend server has the following endpoints:

### `POST /api/auth/join`
### description
This endpoint is used to login the user, the user must send the following data in the body of the request:
```json
{
    "email": "example.com",
    "password": "password"
    "username": "username"
}
```
### response
if the user is created successfully, the server will return the following response:
```json
{
    "token" : "token",
    user: {
        "id": 1,
        "username": "username"
    }
}
```

### `POST /api/auth/login`
### description
This endpoint is used to login the user, the user must send the following data in the body of the request:
```json
{
    "email": "example.com",
    "password": "password"
}
```
### response
if the user is logged in successfully, the server will return the following response
```json
{
    "token" : "token",
    user: {
        "id": 1,
        "username": "username"
    }
}
```
### `GET /api/bonds`
### description
This endpoint is used to get all the bonds that have the current user logged in.
### authorization
You must send the token in the header of the request like this:
```
Authorization: Bearer <token>
```
### response
if the request is successful, the server will return the following response:
```json
{
      "bonds": [
        {
            "id": "3fe44780-b5b5-4e56-a395-22e3f004b2b1",
            "name": "test-1",
            "price": 0.0000,
            "number_bonds": 1300,
            "created_at": "2024-06-18T17:54:03.365939Z",
        }
    ],
    "pagination": {
        "current_page": 1,
        "page_size": 20,
        "last_page": 1,
        "total_records": 1
    }
}
```

### params 
you can send the following query params to the request:
- page: the page of the request
- page_size: the number of records per page

### example
```bash
curl  http://localhost:4000/api/bonds?page=1&page_size=20 -H "Authorization: Bearer <token>" 
```

### `PUT /api/bonds:id/buy`
### description
This endpoint allows the current logged user to buy a bond by this id. if the user has the ownership of the bond, the server will return an error.
### authorization
You must send the token in the header of the request like this:
```
Authorization: Bearer <token>
```
### response
if the request is successful, the server will return the following response:
```json
{
    bond:{
        "id": "3fe44780-b5b5-4e56-a395-22e3f004b2b1",
        "name": "test-1",
        "price": 0.0000,
        "number_bonds": 1300,
        "created_at": "2024-06-18T17:54:03.365939Z",
    }
}
```

### `GET /api/bonds/purchasable`
### description
This endpoint is used to get all the bonds that the current user can buy.
### authorization
You must send the token in the header of the request like this:
```
Authorization: Bearer <token>
```
### response
if the request is successful, the server will return the following response:
```json
{
      "bonds": [
        {
            "id": "3fe44780-b5b5-4e56-a395-22e3f004b2b1",
            "name": "test-1",
            "owner:"user1",
            "price": 0.0000,
            "number_bonds": 1300,
            "created_at": "2024-06-18T17:54:03.365939Z",
        }
    ],
    "pagination": {
        "current_page": 1,
        "page_size": 20,
        "last_page": 1,
        "total_records": 1
    }
}
```

### params
you can send the following query params to the request:
- page: the page of the request
- page_size: the number of records per page

### example
```bash
curl  http://localhost:4000/api/bonds/purchasable?page=1&page_size=20 -H "Authorization: Bearer <token>"
```