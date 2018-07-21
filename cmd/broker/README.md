# ab-lane broker

Broker's purpose is to receive byte messages, store them in a topic and serve them to subscribed consumers.

## Usage

```text
Usage of ./broker:
  -port int
    	Port to bind process (default 9000)
```

## Development

Clone the repository

```
git@github.com:ggirtsou/ab-lane.git
make install # download go dependencies
make protos # generate protos
make test # run tests
make build # generate binary
```
