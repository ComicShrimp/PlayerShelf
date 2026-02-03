# PlayerShelf

Backend PlayerShelf where you can track your games.

## Core Goals

These are the goals that I'll keep in mind when create new features or taking decisions. PS: The goals can change, as always.

- Be able to track my Games (Played, to play, playing and any onther relevant status)
- Overview over my library (How many games I have, in wich platform and etc.)
- Wich platform I have the game (Is any subscription ?, or I bought the game ?)
- Create custom list (Nice to have some base list like games by platform, for example)

## Bonus Goals

Some things that are nice to have, if possible.

- Played hours (In each platform)
- Price tracking/alerts (Be able to set the price that I paid for the game, in each platform, to calculate the total)
- Release calendar

## Running the project

### Requirements for development

- [air](https://github.com/air-verse/air)
- GO

### Download Project Deps

```shell
go mod download
```

### Run Backend

For better development, we use `air` to have a hot reload backend, if you want, you don't need to use it, but is highly recommended

```sh
air
```

### Run local db

The local db is set by the docker compose, it will run an empty DB that you can use for dev and/or tests.

```sh
docker compose up -d
```

### Clean unused deps

After the development that change the deps, run the command below to remove unused deps from `go.mod`.

```sh
go mod tidy
```

### Update deps

```sh
go get -u ./...
```
