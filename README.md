# Go Image Server

This image will use `/images/` directory to store images.

When run in Docker, please make sure to use volumes to avoid data loss.

## Endpoints

#### GET `/ping`

This endpoint will return a pong message. You can use this to check if the server is on.

#### POST `/upload`

Image upload, the body should be the base64 encoded of the image, not include the `data:image....`.

Request body

```json
{
    "base64": "base64 encoded image",
    "ext": "extesion of the image"
}
```

Supported extensions are `jpg`, `png`, and `jpeg`

Server will hash the image with SHA256,
the hash result will be the name of the image.
The extension will be kept as is.

You can retrieve the image with `<sha256 hash>.[png, jpeg, jpg]`

#### GET `/images/*hash.ext`

This endpoint will retrieve the image that you have uploaded.
From the hash calculated when uploading the image and the extension,
you can access the image from server.
