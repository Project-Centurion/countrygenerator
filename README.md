# Country Generator

Simple generator to create maps in Go using [http://country.io](http://country.io) for ISO codes.

It generates a map of country ISO code to country english name, and reversed.

## Usage

Install it by running:

```shell
go install github.com/Project-Centurion/countrygenerator@latest
```

Then run:

```shell
countrygenerator filename.go packagename
```

filename.go is the file (relative to the current directory) where the generator will write the code. 
It does not need to exist.

packagename is the name of the package that will be placed at the top of the generated code file.
