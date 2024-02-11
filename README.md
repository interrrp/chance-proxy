# chance-proxy

> A tiny reverse proxy that disconnects users by chance.

## Building

To build chance-proxy, you will need:

- [Git](https://git-scm.com)
- [Go](https://go.dev) 1.22 or later

After ensuring you have installed all of the prerequisites, run the following commands to build chance-proxy:

```sh
git clone https://github.com/interrrp/chance-proxy
cd chance-proxy
go build -ldflags "-s -w" -o chance-proxy .
```

The executable will now be in the `chance-proxy` directory. Good job! 😀👍

## Configuration

All configuration happens through environment variables or `.env`.

Copy `example.env` into `.env` and change its contents as needed.

### Examples

Listen on `localhost:25566` and proxy to `localhost:25565`

```sh
chance-proxy -address localhost:25566 -target localhost:25565
```

Fail 90% of the time:

```sh
chance-proxy -chance 90
```

## License

chance-proxy is licensed under the [GNU General Public License 3.0](./LICENSE) license.
