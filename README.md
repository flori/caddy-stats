# caddy-stats

## Description

This is a plugin for the Caddy HTTP server that counts the accesses to requested
URLs and logs some of their attributes using redis.

## Installation

```
go get github.com/flori/caddy-stats
```

## Usage

```
import _ "github.com/flori/caddy-stats"
```

## Configuration

The redis server can be configured via URL as the first argument to the `stats` directive in the Caddyfile:

```
https://example.com {
  stats redis://localhost:6379
}
```

Or if the argument is skipped in the Caddyfile it will be fetched from the environment variable `REDIS_URL`
instead:

```
https://example.com {
  stats
}
```

## Changes

* 2017-10-04 Initial Release

## Author

[Florian Frank](mailto:flori@ping.de)

## License

This software is licensed under the Apache 2.0 license.
