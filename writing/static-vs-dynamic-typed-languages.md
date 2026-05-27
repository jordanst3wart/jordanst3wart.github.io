---
title: Python Programming Blog
date: '2024-07-06'
tags: ['python', 'static', 'prototype pattern']
summary: Exploring trade-offs of dynamic vs static typing. Every new programming language has either a type system or has a type spec to support compile time checks.
---

I've been working with Python, and I've started aggressively adding in optional types to functions. I've worked with other dynamic type languages, and didn't really believe the type system was not critical, but now I seem to be relying on it more and more.

I feel two things:

1. Getting language server feedback at development time is amazing.
2. Some types are over-complicated, and I lose time to them (but rarely)

As an example of the first: I was inputting a dictionary, when a list, to a rarely called python function that was handling errors. So the flow should something like: network error-> error handling code -> typed error. It took me a week to find the error.

```Python
def some_rarely_called_function_requiring_a_list(b):
    return a[0]
some_rarely_called_function_requiring_a_list({})
```

The second is when you have spend time understanding the type system, or do type casting such as in TypeScript:

```TypeScript
const address = server.server.address() as AddressInfo
```

I also find the use of `Omit` in Typescript confusing. I have to do maths on the type system rather than the type system do maths for me.

I wonder how many large (or relatively small) python projects have taken a similar response to adding in types in retrospect.

Some of the top python projects on github are:
Photon doesn't use types, but is relatively simple with types (ie. url, host)
https://github.com/s0md3v/Photon/tree/master
Manim does use types, and it's complex:
https://github.com/3b1b/manim
Apache airflow uses types, it's a large codebase:
https://github.com/apache/airflow

Interestingly types for Python was added in 3.5 at 12 Sep 2015. So that makes it relatively new. I'm used to seeing Python code looking like:

```Python
def run(*args, **kwargs):
```

But that might of changed with Python 3.

The requests library doesn't use types, but it's old:
https://github.com/psf/requests

Fastapi uses types:
https://github.com/tiangolo/fastapi

It looks like most larger Python codebases use types, which is odd.

Another camp which was against types was Clojure. Clojure devs are generally very highly performing, and leading with thinking.

They also came up with ClojureSpec, which adds types, and more. That could just be selection bias. ie. if a new language comes out highly intelligent people will pick it up fast, and learn it faster - as that is basically the definition of IQ.

Clojure did have [schema](https://github.com/plumatic/schema), and [herbert](https://github.com/miner/herbert) to add types on before ClojureSpec.

Clojure did largely encourage the use of maps, and lists, which generally don't get as much value out of typing.

Looking at newer programming languages a lot seem to have inferred types, which is what i prefer, like:

- Dart (2011): inferred types
- scala (2004): inferred types
- Ocaml (1996): inferred types
- purescript (2013): inferred types
- go (2009): inferred types
- kotlin (2011): inferred types
- swift (2014): inferred types
- rust (2010): inferred types
- zig (2015): inferred types
- TypeScript (2012): inferred types
- julia (2012): dynamic, but supports type declarations (which are used in libraries)
- Clojure (2007): dynamic (has a type spec with clojureSpec)
- Elixir (2011): dynamic (has a type spec)

So basically every new programming language, either:

1. has inferred types
2. sleepishly adds in a type spec, despite bold claims of being dynamically typed :/

A lot of languages are even adding more compile time checks. I think this kind of settles the typed vs dynamic debate: every new programming language has either a type system or has a type spec to support compile time checks.

Note from 12th, Feb 2025:
I think language server protocols changed the arguement from dynamic to static types.
