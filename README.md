# broker
Middleware that authenticates, authorizes, markets wares, and brokers transactions.

## Setup
The application depends on a Redis cache and a Postgres database. The Redis cache needs a JSON module, and the Postgres database needs a GIS extension.

### Redis
Clone the [RedisJSON](https://github.com/RedisJSON/RedisJSON#redisjson) project, compile the _Rust_ lib, and load the module into the Redis server. This document assumes Redis has been installed by
_brew_ on MacOS, and that _Rust_ has also been installed.

```bash
~/oss $ brew services stop redis
~/oss $ git clone git@github.com:RedisJSON/RedisJSON.git
~/oss $ cd RedisJSON
~/oss/RedisJSON $ cargo build --release
```

```vim
# /opt/homebrew/etc/redis.conf
loadmodule /Users/[myName]/oss/RedisJSON/target/release/librejson.dylib
```

```bash
~/ $ brew services start redis
```

### Postgres
1. The development instance of Postgres currently resides in the AWS VPC. Draft a _~/.pgpass_ file containing `host:port:dbName:dbUser:pw`.
2. Add `export DBSTRING="user=[dbUser] host=[host] port=[port] dbname=[dbName]"` to the interactive shell configuration.

The application will use _DBSTRING_ to find the relevant password written in _~/.pgpass_.

### Application Middleware
The middleware can be developed and run using the following environmental variable.

```bash
~/asymblur/broker $ ENVIRON=dev go run main.go
```
