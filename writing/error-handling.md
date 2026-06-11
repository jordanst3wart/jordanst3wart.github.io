---
title: Error handling of different languages 
date: '2026-06-11'
tags: ['error handling']
summary: 'Error handling in different languages'
---

I saw this go code:

```go
func convertAndPrint(a, b string) error {
    num1, err := strconv.Atoi(a)
    if err != nil {
        return err
    }
    num2, err := strconv.Atoi(b)
    if err != nil {
        return err
    }
    fmt.Println(num1 + num2)
    return nil
}
```

If has 7 lines of error handling code.

Let's see what it looks like in different languages.

Python:
```python
from typing import Optional

def convert_and_print_with_error(a: str, b: str) -> Optional[ValueError]:
    try:
        num1 = int(a)
        num2 = int(b)
        print(num1 + num2)
        return None
    except ValueError as err:
        return err
```
This explicitly handles the error. You could just go:
```python
from typing import Optional

def convert_and_print_with_error(a: str, b: str):
    num1 = int(a)
    num2 = int(b)
    print(num1 + num2)
```
But that doesn't handle the error, which might be okay for short scripts. I would rather just handle the error though.

TypeScript:
```TypeScript
function convertAndPrint(a: string, b: string): void {
    const num1 = Number(a);
    if (isNaN(num1) || a.trim() === "")) {
        throw new Error(`Failed to parse '${value}' into a valid number.`);
    }
    const num2 = Number(b);
    if (isNaN(num2) || b.trim() === "")) {
        throw new Error(`Failed to parse '${value}' into a valid number.`);
    }
    console.log(num1 + num2);
}
```

I normally associate a `NaN` with a float. It's weird having it with a int. I really dislike `NaN` values sneeking in. I feel like this is worst than go somehow, but I guess that is due to the choice of parseInt returning a `NaN` instead of an error, which most languages don't do.

Apparently, `parseInt("10px",10)` returns `10` as well... So you use `Number("10px")`. This might not be TypeScript's fault, it seems like poor JavaScript functions defined to not crash.

Rust:
```rust
use std::num::ParseIntError;

fn convert_and_print(a: &str, b: &str) -> Result<(), ParseIntError> {
    let num1: i32 = a.parse()?;
    let num2: i32 = b.parse()?;

    println!("{}", num1 + num2);
    Ok(())
}
```
The type `Result<(), ParseIntError>` seems complicated already. The `?` seems nice, but i guess that leads you to optimistic error handling. I'm a bit confused with the `.parse()` not having a type as well.

Zig:
```zig
const std = @import("std");

fn convertAndPrint(a: []const u8, b: []const u8) !void {
    const num1 = try std.fmt.parseInt(i32, a, 10);
    const num2 = try std.fmt.parseInt(i32, b, 10);

    std.debug.print("{d}\n", .{num1 + num2});
}
```

I don't like the namespace `std.fmt.parseInt`. I feel like you should just be able to import that function. I'm confused about a string being a `[]const u8` as well. I do like `!void` being neither an error, or nothing. I do like the `try` statement, better than `?` at the end. I don't get the `.{num1 + num2}`, why is it not just `num1 + num2`.


Kotlin:
```kotlin
fun convertAndPrint(a: String, b: String) {
    val num1 = a.toIntOrNull() ?: return
    val num2 = b.toIntOrNull() ?: return

    println(num1 + num2)
}
```
Or:
```
@Throws(NumberFormatException::class)
fun convertAndPrintOrThrow(a: String, b: String) {
    val num1 = a.toInt()
    val num2 = b.toInt()

    println(num1 + num2)
}
```
I guess the first isn't valid as it doesn't handle the error, and the second one doesn't explicity show where the error is coming from, which could make it harder with longer functions.


Scala:
```scala
def convertAndPrintEarlyReturn(a: String, b: String): Unit =
  val num1 = a.toIntOption.getOrElse(return)
  val num2 = b.toIntOption.getOrElse(return)

  println(num1 + num2)
```
Or
```scala
def convertAndPrintEarlyReturn(a: String, b: String): Unit =
  val num1 = a.toIntOption.getOrElse(throw new IllegalArgumentException("Failed to parse int a"))
  val num2 = b.toIntOption.getOrElse(throw new IllegalArgumentException("Failed to parse int b"))

  println(num1 + num2)
```

I thought scala, and kotlin would be like rust, and zig, but they are all different in some odd way. Kotlin is almost more similar to rust, despite not using optionals.

I don't feel like there is a clear winner, and loser. I might try to combine the best though... let's see.


Best:
```text
fn convertAndPrint(a, b string) !void {
  num1 := try parseInt(a)
  num2 := try parseInt(b)
  println(num1 + num2)
}
```
That is `fn` from rust or zig, the `a, b string` from go, the `!void` from zig, the `try` and `parseInt` from zig, the `println(num1 + num2)` from scala or kotlin. That is a lot. It's so weird taking from all the languages. I don't know if a compiler could be written for something like that. It is obviously missing an import.

It's also worth mentioning a lot of the programming language is the tools, security, and how to publish packages too.
