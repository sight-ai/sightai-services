# Go Echo API Server for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0.0
- Build date: 2024-06-26T16:35:53.602608+08:00[Asia/Singapore]

### Running the server

To run the server, follow these simple steps:

```
go mod download
go build -o app
```

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once the image is built, just run
```
docker run --rm -it openapi
```

### Known Issue

TBA