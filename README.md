# MAIL-SERVICE

## For development

This is project use MakeFile

### create `config.yml`

```yaml
app:
  host: localhost
  port: "8080"
smtp:
  host:
  port:
  user:
  password:
  sender: <sender_mail>
  from: <sender_name>
```

### Run server

serve on `http://localhost:8080`

```sh
make run
```
