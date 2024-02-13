# Interpreter in Go

A Tree-Walking Intepreter

This repository contains the C like Programming Language implementation. The lexer in the lexer package converts REPL input into tokens, and a Pratt parser in the parser package transforms these tokens into an Abstract Syntax Tree (AST). The evaluator, located in the evaluator package, uses the Eval function to traverse the AST and produce values. The result is printed to the REPL.

Currently extending the programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com) by [Thorsten Ball](https://github.com/mrnugget).

## It Supports the following list of features:

• integers
• booleans
• strings
• arrays
• hashes
• prefix-, infix- and index operators
• conditionals
• global and local bindings
• first-class functions
• return statements
• closures

### Let Statements

You can define identifiers with using the `let` keyword.

```sh
>> let x = 6;
>> x
6
>> let x = 99;
>> y
99
>> let subtract = fn(a,b) { return a - b; };
```

### If Expressions

You can define if statements using the `if` keyword.

```sh
>> let x = 5
>> let y = 10
>> if (x < y) {return x ; } else { return y; }
5
```

### Function Literals

You can define functions with using the `fn` keyword, followed by parameters and a block statement executed when the function is called.

```sh
>> let min = fn(x,y) { if (x < y) { x } else { y } };
>> min(5, 10)
5
```

### Integer Literals

```sh
>> let x = 6;
>> x
6
>> let x = 99;
>> y
99
```

### Boolean Literals

```sh
>> let a = true;
>> a
true
>> let b = false;
>> b
false
>> a == b
false
```

### Grouped Expressions

```sh
>> (5 + 5) * 2
20
>> -5 * (2 + 5)
-35
```

### Array Literals

```sh
>> let myArray = ["Jason", "Yao", 27, fn(x) {x + x}];
>> myArray[0]
Jason
>> myArray[2]
27
>> myArray[3](2);
4
```

### Hash Literals

```sh
>> let myHash = {"name": "Jason", "age": 18, "language": "Go"};
>> myHash["name"]
Jason
>> myHash["age"]
18
>> myHash["language"]
Go
```
