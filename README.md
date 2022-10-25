<div align="center">
  <h1>Jomon</h1>
  <p>
    <strong>An accounting support system</strong>
  </p>
  <p>
    <a href="https://apis.trap.jp/?urls.primaryName=Jomon%20v2%20API">Jomon API</a>&emsp;<a href="https://github.com/traPtitech/Jomon-UI">Jomon-UI</a>
  </p>
  <p>
    <a href="https://github.com/traPtitech/Jomon/actions/workflows/image-v2.yml"><img src="https://github.com/traPtitech/Jomon/actions/workflows/image-v2.yml/badge.svg"></a>
    <a href="https://github.com/traPtitech/Jomon/actions/workflows/go.yml"><img src="https://github.com/traPtitech/Jomon/actions/workflows/go.yml/badge.svg"></a>
    <a href="https://codecov.io/gh/traPtitech/Jomon"><img src="https://codecov.io/gh/traPtitech/Jomon/branch/v2/graph/badge.svg"></a>
  </p>
</div>

## Environment

### Testing

1. Make the server running.

2. Run the following command in the project root.
```shell script
make test
```

### Running

1. Run the following command in the project root.

```sh
make up
```

Now, you can send http requests to `localhost:3000`.

2. Run the following command in the project root when making the server down.

```sh
make down
```

## Staging

1.Enter the server for Jomon staging server and run the following command in the project root.

```sh
sudo docker pull ghcr.io/traptitech/jomon-v2:latest
sudo docker run -d -p 1323:1323 --env-file .env ghcr.io/traptitech/jomon-v2
```

(At first, you need to set .env file)
