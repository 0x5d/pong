# Pong

1/2 of a simple async ping pong app built for my Component-Based Software Development course, using Golang
and RabbitMQ.

## Run it

- Make sure [Golang](https://golang.org/dl/) is properly installed.

- pong requires a running instance of RabbitMQ. If there's one running already, don't worry. If not, you can run it with the [Docker RabbitMQ container](https://hub.docker.com/_/rabbitmq/):
  ```sh
  docker run -d --hostname ping-pong-broker --name broker rabbitmq:3
  ```

- Clone this repo.

  ```sh
  git clone https://github.com/castillobg/pong.git
  ```

- Build the `pong` executable.

  ```sh
  # Inside the cloned pong repo:
  go build
  ```

- Run `pong`:
  ```
  ./pong -port 8080 -broker rabbit -address localhost:5672
  ```
  or simply
  ```
  ./pong
  ```

- You can now go over to the [ping](https://github.com/castillobg/pong) repo and start ping-ponging.
