# BBB MS

Step to run the project:

Run the docker compose up with build flag:

This will build the rabbit mq queue and the postgres db

```sh
docker compose up --build
```

Run the makefile in the terminal, this will up the app and the load test:

```make
  make run & make load-test
```

> If you haven't the air installed for golang.
> then run this in the terminal:

```make
  make run-no-air & make load-test
```
