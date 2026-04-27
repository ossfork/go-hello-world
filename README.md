# README

---

## Test with Docker Locally

- Build the Image
````
```sh
docker build -t go-hello-world:local .
```

- Run & Test

```sh
cd /path/to/hello-world

go fmt ./...
go test ./...
go run .
curl http://localhost:8080
```

```sh
docker run --rm -p 8080:8080 go-hello-world:local
curl http://localhost:8080
```

- Tag Image

```sh
# Docker Hub
docker build --platform linux/amd64 -t "anonid/go-hello-world:$VERSION-amd64" .
docker build --platform linux/arm64 -t "anonid/go-hello-world:$VERSION-arm64" .
docker build -t anonid/go-hello-world:latest .
docker tag anonid/go-hello-world:latest anonid/go-hello-world:v0.1.0

# GitHub
docker build -t ghcr.io/ossfork/go-hello-world:latest .
docker tag ghcr.io/ossfork/go-hello-world:latest ghcr.io/ossfork/go-hello-world:v0.1.0
```

- Push to a remote registry

```sh
# Docker Hub
docker login
docker push anonid/go-hello-world:latest
docker push anonid/go-hello-world:v0.1.0

# GitHub
echo "$GITHUB_TOKEN" | docker login ghcr.io -u ossfork --password-stdin
docker push ghcr.io/ossfork/go-hello-world:v0.1.0
docker push ghcr.io/ossfork/go-hello-world:latest
```

- Pull Image

```sh
docker pull anonid/go-hello-world:latest
docker run --rm -p 8080:8080 anonid/go-hello-world:latest
docker run --name helloworld \
  --restart unless-stopped \
  -p 8080:8080 \
  -d anonid/go-hello-world:latest
```


