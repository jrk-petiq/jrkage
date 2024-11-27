# Installation

Ensure you have go installed on your machine. If not, you can install it with homebrew on MacOS:

```bash
brew install go
```

Ensure go is added to your path. In .zshrc or .bashrc:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

```

Then, you can install the package with the following command:

```bash
go install github.com/jrk-petiq/jrkage@latest
```

You should then be able to run the program with the following command:

```bash
jrkage
```
