
## Prerequisites

- Docker installed on your machine. You can download it from [here](https://www.docker.com/products/docker-desktop).

## Running the Application

To run the application using Docker, follow the steps below:

### Build and run the Docker Image

```bash
docker build --pull --rm -f 'dockerfile' -t 'gameoflife:latest' '.' 

docker run gameoflife