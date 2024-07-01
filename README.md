
## Installation

### Clone the Repo

```bash
  git clone https://github.com/j0shiJ/leaderboard-service.git
  cd leaderboard-service

```

### 1. Setup Postgres

1. Install PostgreSQL and create a database named leaderboard.
2. Update the database connection URL in main.go and repository/repository.go.


### 2. Build and run the service locally:

```bash
go build -o leaderboard-service .
./leaderboard-service

```
The service will run on http://localhost:8080 by default

### 3. Run using Docker:

```bash
docker build -t leaderboard-service .
docker run -p 8080:8080 leaderboard-service
```


## API Documentation (Swagger)

Swagger documentation is generated for the API endpoints. Access the documentation by running the service locally or using Docker and visiting http://localhost:8080/swagger/index.html in your browser.

Submit Score:

- URL: /submit  
- Method: POST
- Request Body:
```json
{
    "user_name": "User 1",
    "country": "India",
    "state": "Karnataka",
    "score": 100
}
```

Get Rank:

- URL: /get_rank
- Method: GET
- Query Parameters: userName, category (state, country, globally)


List Top N Scores:

- URL: /list_top_n
- Method: GET
- Query Parameters: n, category (state, country, globally)

## Screen showing working project

## Screenshots

![Screenshot 1](Screenshots/Screenshot 2024-07-01 at 11.22.59 PM.png)
*Swagger Documentation*

![Screenshot 2](Screenshots/Screenshot 2024-07-01 at 3.06.27 PM.png)
*PostgreSQL db*

![Screenshot 3](Screenshots/Screenshot 2024-07-01 at 3.02.31 PM.png)
*Postman*






