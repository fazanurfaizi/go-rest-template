## Project Overview
Clean architecture template with gin gonic as web framework, go-fx as dependency container, gorm as orm for database related operations, with air as hot-reload.

## Running the project for local environment

- Declare config in config/config-local.yml
- Run `air --build.cmd "go build -o bin/api cmd/app/main.go" --build.bin "./bin/api app:serve"`
- Go to `localhost:5000` to verify if the server works.

#### Environment Variables

<details>
    <summary>Variables Defined in the project </summary>

| Key                            | Value                    | Desc                                        |
| ------------------------------ | ------------------------ | ------------------------------------------- |
| `server.AppName`               | `go_rest_template`       | Application name                            |
| `server.Port`                  | `5000`                   | Port at which app runs                      |
| `server.Mode`                  | `development,production` | App running Environment                     |
| `server.JwtSecretKey`          | `secret`                 | JWT Token Secret key                        |
| `LOG_OUTPUT`                   | `./server.log`           | Output Directory to save logs               |
| `logger.Level`                 | `info`                   | Level for logging (check lib/logger.go:172) |
| `postgres.PostgresqlUser`      | `username`               | Database Username                           |
| `postgres.PostgresqlPassword`  | `password`               | Database Password                           |
| `postgres.PostgresqlHost`      | `0.0.0.0`                | Database Host                               |
| `postgres.PostgresqlPort`      | `3306`                   | Database Port                               |
| `postgres.PostgresqlDbname`    | `test`                   | Database Name                               |

</details>

## Implemented Features

- Dependency Injection (go-fx)
- Routing (gin web framework)
- Environment Files
- Logging (file saving on `production`) [zap](https://github.com/uber-go/zap)
- Middlewares (cors)
- Database Setup (postgresql)
- Models Setup and Automigrate (gorm)
- Repositories
- Live code refresh
- Cobra Commander CLI Support. try: `go run . --help`
