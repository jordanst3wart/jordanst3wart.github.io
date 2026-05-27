---
title: Exploring the strategy pattern
date: '2025-02-06'
tags: ['kotlin', 'strategy pattern']
summary: Exploring the strategy pattern in kotlin
---

I wanted to explore the strategy pattern in a modern software concept, and in kotlin in particular.

If you look at the strategy design pattern from a source like refactoring guru, and convert it to kotlin you get this:

```kotlin
package org.example

interface Strategy {
    fun execute(a: Int, b: Int): Int
}

class ConcreteStrategyAdd : Strategy {
    override fun execute(a: Int, b: Int): Int {
        return a + b
    }
}

class ConcreteStrategySubtract : Strategy {
    override fun execute(a: Int, b: Int): Int {
        return a - b
    }
}

class ConcreteStrategyMultiply : Strategy {
    override fun execute(a: Int, b: Int): Int {
        return a * b
    }
}

class Context {
    private lateinit var strategy: Strategy

    fun setStrategy(strategy: Strategy) {
        this.strategy = strategy
    }

    fun executeStrategy(a: Int, b: Int): Int {
        return strategy.execute(a, b)
    }
}

fun main() {
    val context = Context()
    // actually get an input
    println("Enter first number:")
    val firstNumber = readlnOrNull()?.toIntOrNull() ?: 0

    println("Enter second number:")
    val secondNumber = readlnOrNull()?.toIntOrNull() ?: 0

    // if statements...
    println("Enter operation (addition/subtraction/multiplication):")
    when (readlnOrNull()?.lowercase()) {
        "addition" -> context.setStrategy(ConcreteStrategyAdd())
        "subtraction" -> context.setStrategy(ConcreteStrategySubtract())
        "multiplication" -> context.setStrategy(ConcreteStrategyMultiply())
        else -> {
            println("Invalid operation")
            return
        }
    }

    val result = context.executeStrategy(firstNumber, secondNumber)
    println("Result: $result")
}
```

So it's an interface basically. In this example you need to use a `when`, or switch statement to set the strategy, which is pretty common. A nicer way of doing this is probably using a map for the constructor of the strategy like:

```kotlin
package org.example

enum class Operations(val text: String) : Strategy {
    ADDITION("addition") {
        override fun execute(a: Int, b: Int): Int {
            return a + b
        }
    },
    SUBTRACTION("subtraction") {
        override fun execute(a: Int, b: Int): Int {
            return a - b
        }
    },
    MULTIPLICATION("multiplication") {
        override fun execute(a: Int, b: Int): Int {
            return a * b
        }
    };

    companion object {
        fun fromString(value: String?): Strategy {
            value ?: throw IllegalArgumentException("Input is null")
            val operation = entries.find { it.text.equals(value, ignoreCase = true) }
            operation ?: throw IllegalArgumentException("Invalid operation")
            return operation
        }
    }
}

fun main() {
    println("Enter first number:")
    val firstNumber = readlnOrNull()?.toIntOrNull() ?: 0

    println("Enter second number:")
    val secondNumber = readlnOrNull()?.toIntOrNull() ?: 0

    println("Enter operation (addition/subtraction/multiplication):")
    val strategy = Operations.fromString(readlnOrNull()?.lowercase())
    val result = strategy.execute(firstNumber, secondNumber)
    println("Result: $result")
}
```

This avoids the `when`, and saves 10 lines of code. I think it is a bit nicer, but uses more advanced language features, like enums, and companion objects.

This is what the code looks like inlined:

```
package org.example

fun main() {
    println("Enter first number:")
    val firstNumber = readlnOrNull()?.toIntOrNull() ?: 0

    println("Enter second number:")
    val secondNumber = readlnOrNull()?.toIntOrNull() ?: 0

    // if statements...
    println("Enter operation (addition/subtraction/multiplication):")
    val result = when (readlnOrNull()?.lowercase()) {
        "addition" -> firstNumber + secondNumber
        "subtraction" -> firstNumber - secondNumber
        "multiplication" -> firstNumber * secondNumber
        else -> {
            throw Exception("Invalid operation")
        }
    }
    println("Result: $result")
}
```

It saves a further 20 lines of code, doesn't use advanced language features, has less redirection, and is easier to read. This example is potentially not that straight forward as generally a strategy would have significant more code, and more complexity than addition, and substraction. It's worth pointing out though that you shouldn't just default to a design pattern.

I've been thinking about Elon Musk's 5 step process in design and engineering, and step one and two are:

1. Make requirements less dumb: Question and verify each requirement to make it less error-prone.
2. Delete the part or process: Remove unnecessary parts or process steps.

I think in the example above you can use a negative number for substraction, like instead of inputting '1', '2', and 'subtraction', you can just input '1', and '-2'. I think you could also argue that multiple might not be needed, or why is division not there? I think generally in software you have several things you need like addition, subtraction, multiplication, division, exponentiation (like 2^3), root operations (cube root of 8), modulo (7 mod 3 = 1), logarithms (like log₂(8) = 3), sine and cosine operations for a partial calculator app. Or you just need one, like addition. So in this case I would say you just need addition, and subtraction.

So for just addition, and subtraction, the code looks like this:

```kotlin
package org.example

fun main() {
    println("Enter first number:")
    val firstNumber = readlnOrNull()?.toIntOrNull() ?: 0

    println("Enter second number:")
    val secondNumber = readlnOrNull()?.toIntOrNull() ?: 0

    // assuming multiplication isn't needed, and can just enter a negative number for subtraction
    val result = firstNumber + secondNumber
    println("Result: $result")
}
```

That's 13 lines of code. That's quiet simple, but a lot of things are.
