# Interpreter in Go

Currently extending the programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com) and [_Writing a Compiler in Go_](https://compilerbook.com) by [Thorsten Ball](https://github.com/mrnugget).

## It supports:

## Let Statements

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
