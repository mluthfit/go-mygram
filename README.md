# About

MyGram is photo storage application and every registered user can comment on other users' photos

## Getting Started

This project developed using `go 1.19.2` and `mysql`.

To start running this project locally, you must follow these steps:

First, clone these repository to the your folder.

```
> https://github.com/mluthfit/go-mygram.git
```

Then, open the folder and **install** all packages.

```
> go mod tidy
```

Then, copy `.env.example` to `.env`

```
cp .env.example .env
```

Then fill the environment variables.

Last, start the server.

```
> go run main.go
```
