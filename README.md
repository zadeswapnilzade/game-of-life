

## Option 1 - Using go
## Prerequisites
- go installed on your machine. You can download it from [here](https://go.dev/doc/install).
- git installed on your machine. You can download it from [here](https://git-scm.com/downloads).

```bash
git clone https://github.com/zadeswapnilzade/game-of-life.git

```bash
cd game-of-life

```bash
go run main.go

---------------------------------------------------------------------------------------------------------------------------

## Option 2 - Using Docker
## Prerequisites

- Docker installed on your machine. You can download it from [here](https://www.docker.com/products/docker-desktop).

## Running the Application

To run the application using Docker, follow the steps below:

### Build and run the Docker Image

```bash
docker build --pull --rm -f 'dockerfile' -t 'gameoflife:latest' '.' 

docker run gameoflife
