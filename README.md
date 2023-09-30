# petnamesd

## Motivation

A small webservice that provides a GET /petnames endpoint that returns petnames looking like this:

```
{
    "pet_name": "cool-cat"
}
```

## Installation

Simply do:

```
go install github.com/4thel00z/petnamesd/...@latest
```

### Usage

```
PORT=1234 petnamesd &
curl "http://localhost:1234/petnames?words=3&separator=_"
```
## License

This project is licensed under the GPL-3 license.
