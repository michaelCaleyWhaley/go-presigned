# go-presigned

## todo
- ~~create landing page for client to request url from~~
- ~~use aws sdk to generate signed S3 url~~
- ~~send signed url to client~~
- ~~create client for user to upload from~~
- ~~test large files~~
- ~~allow access to uploaded files~~
- set file type for upload
- limit file types
- store reference to signed url and emit event when user finishes uploading
- limit cross origin resource sharing for signed urls
- architectural drawing on draw.io

### todo maybe
- terraform runtime for go
- terraform database

## Notes
- Go templating https://gin-gonic.com/docs/examples/html-rendering/
- gin docs https://pkg.go.dev/github.com/gin-gonic/gin#Engine.Run

## Commands
- `go run presigned.go`


https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys

AWS
caleycode