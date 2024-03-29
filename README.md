# MailHedgehog

Email previewer. Inspired by [MailHog](https://github.com/mailhog/MailHog). 
Allows to create multiple rooms with different credentials

![mailhedgehog.gif](docs%2Fimg%2Fmailhedgehog.gif)

## Configuration and usage

```shell
MailHedgehog init
```

## Usage

```shell
MailHedgehog serve [.mh-config.yml]
```

## Development

```shell
go mod tidy
go mod verify
```

Example configuration:
```yaml
#.mh-config.yml
smtp:
  port: 1026
http:
  port: 8026
  # path: box
  allow_origins: "http://localhost:5173"
  # assets_root: "./public"
storage:
  use: directory
  config:
    per_room_limit: 100
  directory:
    path: "/home/me/work/go/mailhog/_storage"
#  use: mongodb
#  #  per_room_limit: 100
#  mongodb:
#    connection: mongodb
#    collection: emails_dev
ui:
  file: ui.json
authentication:
  admin: testuser
#  use: mongodb
#  mongodb:
#    connection: mongodb
#    collection: users_dev
  use: file
  file:
    path: ".mh-authfile"
sharing:
    use: csv
    csv:
        path: ".mh-sharing.csv"
# dbConnectionMongo:
#  connections:
#    mongodb:
#      uri: 127.0.0.1:27017
#      db_name: mailhedgehog
#      db_user: root
#      db_pass: secret
```

## Credits

- [![Think Studio](https://yaroslawww.github.io/images/sponsors/packages/logo-think-studio.png)](https://think.studio/)
