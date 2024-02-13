# Interpreter in Go

Currently extending the Monkey programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com) and [_Writing a Compiler in Go_](https://compilerbook.com) by [Thorsten Ball](https://github.com/mrnugget).

## It supports:

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

### If Expressions

```sh
>> let x = 5
>> let y = 10
>> if (x < y) {return x ; } else { return y; }
5
```

### Function Literals

```sh
>> let min = fn(x,y) { if (x < y) { x } else { y } };
>> min(5, 10)
5
```
