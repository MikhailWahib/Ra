# Functions

Functions are a core feature of Ra. You can define your own functions and use built-in functions.

## Defining Functions
Use the `fn` keyword to define functions:
```js
let add = fn(x, y) {
    return x + y;
};

puts(add(3, 7));  # Call the function
```

## Built-in Functions
Ra provides built-in functions like `puts`, `max`, and `len`:
```js
puts("Hello, Ra 𓋹!");  # Print to console

let maxInt = max(10, 20, 5);  # Find the maximum value
puts(maxInt);
```

## Recursion
Ra supports recursive functions:
```js
let factorial = fn(n) {
    if (n == 1) {
        return 1;
    } else {
        return n * factorial(n - 1);
    }
};

puts(factorial(5));  # Calculate factorial of 5
```

## Closures
Ra supports closures, which are functions that capture variables from their surrounding environment:
```js
let count = 0;
let add = fn() {
    count += 1;
    return count;
};

puts(add());  # Prints 1
```
