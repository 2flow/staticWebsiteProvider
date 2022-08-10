# Build and Deploy
## Build
```shell
docker build -t azure-static-website-provider:latest .
```
## Deploy
```shell
docker build -t flowimmer/azure-static-website-provider:latest .
docker push flowimmer/azure-static-website-provider:latest
```
