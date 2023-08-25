# Notes on Go Modules

From https://go.dev/doc/modules/managing-dependencies

* To find latest version of all modules

```bash
go list -m -u all | grep ']$' | cat -
```

* To get a specific numbered version, append the module path with an @ sign followed by the version you want:

```bash
go get example.com/theirmodule@v1.3.4
```

* To get the latest version, append the module path with @latest:

```bash
go get example.com/theirmodule@latest
```

See [Stackoverflow](https://stackoverflow.com/questions/65683782/how-to-identify-dependency-chain-using-go-modules)for info on finding out *why* a package is needed.

```bash
go mod graph | grep github.com/russross/blackfriday
```