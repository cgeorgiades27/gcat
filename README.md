# gca

## Description

gcat is a Go implementation of the Unix `cat` utility. It reads files sequentially, writing them to standard output. It includes additional features such as line numbering and adding blank spaces after lines.

## Installation

To install gcat, you need to have Go installed on your machine. Then, you can build the project using the following commands:

```sh
git clone https://github.com/cgeorgiades27/gcat.git
cd gcat
go build -o gcat main.go
```

## Usage

Here are some basic usage examples:

#### Example 1: Display contents of a file

`./gcat filename.txt`

#### Example 2: Display contents with line numbers

`./gcat -n filename.txt`

#### Example 3: Display contents with line numbers and line breaks

`./gcat -n -b filename.txt`

#### Example 4: Pipe stdout

`echo "hello there" | ./gcat`

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
