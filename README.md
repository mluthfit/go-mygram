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

Then, adjust the database configuration in `utils/db.go` file.

Last, start the server.

```
> go run main.go
```
