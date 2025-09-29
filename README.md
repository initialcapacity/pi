# Pi Estimation

Estimates the value of π using the [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method).

## Method

We generate random points within the unit square and count how many fall within a circle of radius 1 centered at the origin.
The ratio of the number of points within the circle to the total number of points approaches π/4.

![Illustration of a Monte Carlo simulation](./readme_images/monte-carlo.svg)

## Run Estimation

Build and run an estimation with the default configuration.

```shell
go build ./cmd/pi
./pi
```

Use the `--help` flag to see configuration options.

```shell
./pi --help
```

## Development

1.  Run tests
    ```shell
    go test ./...
    ```

1.  Check formatting
    ```shell
    gofmt -d .
    ```

1.  Check linking
    ```shell
    go vet ./...
    ```

1.  Generate new diagram for the readme
    ```shell
    go run ./cmd/generatediagram
    ```
