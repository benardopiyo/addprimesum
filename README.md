# addprimesum

## Instructions

* The `addprimesum` program takes a single positive integer as an argument and calculates the sum of all prime numbers that are less than or equal to the given integer. The result is then printed followed by a newline ('\n').
Requirements

* If the program receives exactly one argument and it is a positive integer, it will output the sum of all prime numbers less than or equal to that integer.
    
* If the number of arguments is different from 1, or if the argument is not a positive number, the program outputs 0 followed by a newline.

## Usage

To run the program, use the following commands in your terminal:


```bash
$ go run . 5
10
```
```bash
$ go run . 7
17
```
```bash
$ go run . -2
0
```
```bash
$ go run . 0
0
```
```bash
$ go run .
0
```
```bash
$ go run . 5 7
0
```

## Running the Program

Ensure you have Go installed on your system. Place the program in a file (e.g., main.go) and use the Go command to run it with the desired input.

