# URL Shortener

URL Shortener is a Go-based web application that allows users to shorten long URLs into more manageable and shareable links. Built with simplicity and performance in mind, it provides a fast and reliable solution for creating and managing shortened URLs.

## Features
- Shorten long URLs into easy-to-share links.
- Customize shortened URLs with user-defined aliases.
- 5 years Expiration dates for shortened URLs.
- Redis for rate-limiting
- Secure and scalable architecture using Go and Fiber framework.

## Instalation

To install and run the URL Shortener application locally, follow these steps:

1.  Clone the repository
 ```bash
git clone https://github.com/nvlhnn/url-shortener.git
```

2. Navigate to the project directory
 ```bash
cd url-shortener
```

3. Build the Docker image
 ```bash
docker-compose build
```

4. Start the application
 ```bash
docker-compose up
```

## Configuration
The URL Shortener application can be configured using environment variables. You can customize database settings, memory store configurations, rate limiting parameters, and more by modifying the .env file in the project root directory.

```bash
DATABASE_HOST=host.docker.internal
DATABASE_PORT=3306
DATABASE_USER=nvlhnn
DATABASE_PASSWORD=here
DATABASE_NAME=url_shortener


MEMSTORE_HOST=host.docker.internal
MEMSTORE_PORT=6379
MEMSTORE_PASSWORD=password
MEMSTORE_DATABASE=0


LIMITER_MAX=10
LIMITER_EXPIRATION=60
```

## Postman Documentation
Explore the API endpoints and interact with the URL Shortener application using the [Postman Documentation](https://www.postman.com/backend-jarivis/workspace/public-nvlhnn/request/24180400-11440598-749a-4222-84d2-12f6f6593418).

