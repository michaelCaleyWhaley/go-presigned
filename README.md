# go-presigned

## todo
- ~~create landing page for client to request url from~~
- ~~use aws sdk to generate signed S3 url~~
- ~~send signed url to client~~
- ~~create client for user to upload from~~
- ~~test large files~~
- ~~allow access to uploaded files~~
- ~~limit file size~~
- add login
- scope uploads to user
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

Integrate sign up 
https://go-presigned.auth.eu-west-2.amazoncognito.com/login?client_id=1g2l1ifmfcano6vm58quvb9ogj&response_type=code&scope=email+openid+phone&redirect_uri=http%3A%2F%2Flocalhost%3A3000

https://stackoverflow.com/questions/45785898/how-to-use-the-code-returned-from-cognito-to-get-aws-credentials