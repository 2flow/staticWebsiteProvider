# Build and Deploy
Short go build which provides a static webpage from azure storage blob.
But maybe there is already a better version existing.
I will update this readme as soon as I know .
## Build
```shell
docker build -t azure-static-website-provider:latest .
```
## Deploy
```shell
docker build -t flowimmer/azure-static-website-provider:latest .
docker push flowimmer/azure-static-website-provider:latest
```
