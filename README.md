# Monkey Language Interpreter (Go)

This repository contains a complete interpreter for the Monkey programming language, built from scratch in Go. This project was a focused learning exercise to understand the fundamentals of programming language design and implementation.

---

## Guide

The interpreter was developed by following the excellent guide:

**Writing an INTERPRETER In Go**
by Thorsten Ball

Following this book provided a deep, hands-on understanding of the components and design patterns that make an interpreter work, from lexical analysis to program evaluation.

---

## Features

This interpreter is fully functional and supports a wide range of features from the Monkey language:

* **Lexical Analysis (Lexer):** Converts source code into a stream of tokens.
* **Parsing (Parser):** Transforms the token stream into a hierarchical Abstract Syntax Tree (AST).
* **Data Types:** Supports integers, booleans, arrays, strings, and hash maps.
* **Functions & Closures:** Allows for user-defined functions and the creation of closures.
* **Control Flow:** Includes support for `if/else` conditional expressions and `return` statements.
* **Built-in Functions:** Provides a set of core built-in functions for working with data, such as `len`, `first`, `last`, `rest`, `push`, and `puts`.

---

## Getting Started

To run the interpreter's Read-Eval-Print Loop (REPL), follow these steps:

1.  Clone this repository to your local machine:
    ```bash
    git clone https://github.com/itsdrelolz/go-interpreter-monkey.git](https://github.com/itsdrelolz/go-interpreter-monkey.git
    cd monkey
    ```

2.  Run the main Go file:
    ```bash
    go run main.go
    ```

You should now be in the interactive Monkey REPL:
let five = 5;
five * 5;
25
let add = fn(x, y) { x + y; };
add(5, 5);
10

---

## Ideas for the Future

This project is a solid foundation, and there are several ways it could be expanded upon to further explore programming language concepts.

* **Bytecode and Virtual Machine:** The book mentions that a significant performance increase can be achieved by compiling the AST down to bytecode and running it on a virtual machine. This is a concept I would like to explore.
* **Syntax Changes:** I would like to experiment with modifying the language's syntax to better align with my personal preferences, such as adding a new way to declare variables or different syntax for function definitions.
* ** The process of creating a compiler. 
