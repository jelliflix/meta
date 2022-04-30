Meta
====

Meta provider for movies and TV series by Cinemeta.

## Usage

### Installation

#### `go get`

```shell
$ go get -u -v github.com/jelliflix/meta
```

#### `go mod` (Recommended)

```go
import "github.com/jelliflix/meta"
```

```shell
$ go mod tidy
```

### API

#### Meta provider

```go
GetMovie(id string) (Meta, error)
GetSeries(id string) (Meta, error)
```

GetX returns meta for movie or tv series.

##### Examples

```go
import "github.com/jelliflix/meta"

client := meta.NewCinemeta(meta.DefaultOptions)
response, err := client.GetMovie("tt9170516")
if err != nil {
    log.Fatal(err)
}

log.Println(response)
// Output:
// {tt9170516 The Shadow in My Eye movie The fates of several Copenhagen...}
```
