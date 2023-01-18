# broker
Middleware that authenticates, authorizes, markets wares, and brokers transactions.

## Setup
The application depends on a Redis cache and a Postgres database. The Redis cache needs a JSON module, and the Postgres database needs a GIS extension.

### Redis
Clone the [RedisJSON](https://github.com/RedisJSON/RedisJSON#redisjson) project, compile the _Rust_ lib, and load the module into the Redis server. This
document assumes Redis has been installed by _brew_ on MacOS, and that _Rust_ has also been installed.

```bash
~/oss $ git clone git@github.com:RedisJSON/RedisJSON.git
~/oss $ cd RedisJSON
~/oss/RedisJSON $ cargo build --release
```

```vim
# /opt/homebrew/etc/redis.conf
loadmodule /Users/[myName]/oss/RedisJSON/target/release/librejson.dylib
```

```bash
~/ $ brew services restart redis
```

### Postgres
Local development lacks a password. Prior local development needed a _~/.pgpass_, but recent re-configuration of _Broker_ abandoned the use of that file. The
future deployment of _Broker_ as a container will use environ variables, and a call to AWS _Secrets_ for a database password.

#### Postgres environmental variables
| VAR | TYPE | EXAMPLE |
| --- | --- | --- |
| BROKER_DBUSER | string | "clerk" |
| BROKER_DBHOST | string | "localhost" |
| BROKER_DBPORT | string | "60000" |
| BROKER_DBNAME | string | "omni" |

The application will form a string from the environmental variables and use the string to connect to the database. It will find the relevant password written in _~/.pgpass_.

### Application Middleware
The middleware can be developed and run using the following command.

```bash
~/asymblur/broker $ go run main.go
```

#### AWS Session
The AWS session prioritizes information found in the _~/.aws/config_ file. Simply provide the environ variable _AWS_PROFILE_.
**This may change for deployment to EC2.**
```go
// security/ticket.go
func CreateAwsSession() *awsSession.Session {
	configSec := readConfig()

	options := awsSession.Options{
		Profile:           configSec.AwsProfile,
		SharedConfigState: awsSession.SharedConfigEnable,
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}

	return awsSession.Must(awsSession.NewSessionWithOptions(options))
}
```

#### AWS Secrets Manager environmental variables
The Golang webserver uses the retired [Gorilla SecureCookie](https://github.com/gorilla/securecookie) for signing and encrypting cookies. The cookie uses two
keys: A _hash key_ that needs AES-256, and a _block key_ that needs AES-128. Simply provide the names of the AWS Secrets as environ variables to the Golang
webserver. A session ID is written into the cookie.

| VAR | TYPE | EXAMPLE |
| --- | --- | --- |
| SECRET_COOKIE_AUTH | string | "alexander" |
| SECRET_COOKIE_ENCRYPT | string | "localhost" |

#### Router environmental variables
| VAR | VALUE |
| --- | --- |
| APP_PROXIES | "127.0.0.1/32,XXX.XX.XX.XX/24,XXX.X.X.X/24"
