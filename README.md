# traefik-plugin-brotli

WIP - trying to get brotli compression working as traefik middleware plugin

# Getting Started

create `.env` file with `TRAEFIK_PILOT_TOKEN`:

```conf
TRAEFIK_PILOT_TOKEN=secret-pilot-token
```

```sh
docker-compose up
```

```sh
curl --verbose  --header 'Host:file-traefik-plugin-brotli.docker.localhost' --header 'Accept-Encodin
g: br' 'http://localhost:80/traefik'
```
