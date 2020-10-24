M9 / LOGIN
----------

A module with features to sign in and sign up using default method using API and service provider, such as Facebook & Google.

### GET STARTED

1. Copy the `config.yaml.example` to `config.yaml`
2. Update the configuration according to your requirement.
3. Create database (manually) and migrate the schema using this command
```
m9 migrate -m login --verbose
```
4. Start the application server
```
m9 server start login

* use --verbose to verbose the sql query
```

### FEATURES

| Features | Status |
|----------|--------|
| Feature to login/register using Facebook | DONE |
| Feature to login/register using Google | DONE |
| Feature to redirect to target URL feature (if enabled) after access token is genereted when login using service provider | DONE |
| Feature to signup using normal method via REST API| DONE |
| Feature to login using normal method via REST | DONE |
| Feature to validate access token | DONE |
| Feature to reset password | BACKLOG |
