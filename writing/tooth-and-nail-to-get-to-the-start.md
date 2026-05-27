---
title: Python Coding Challenges
date: '2024-07-06'
tags: ['python']
summary: Tooth and nail to get to the start; python is only good with data scientist
---

Python is not ideal for building software, or building anything besides machine learning models. Python is about 3 times slower than older languages. It's slower writing larger programs with no types, and the type hints are not well supported. It's hard to deploy with its dependencies. It's hard to manage local environments with virtual environments. The packages are often lacking with error handling. Python relies on a lot of odd syntax.

Python requires a lot more compute resources to do the same amount of work than most other programming languages. As a rough number python is about 3 times slower than java.

Python is generally said to a fast language to development in. This is generally offset by bad library support for more advanced tasks, bad support for a type system, and don't look at async tasks. The common libraries are great, but things like writing to databases, grpc, and not ideal. A lot of libraries rely on `*args, **kwargs` and that does not help with using data classes, data transfer object, and entities. Having a lot of older program use `args` and `kwargs` prevents typing from being any good.

The typing tools upholding python are mypy (Python community), Pyre (Facebook), Pytype (Google), and Pyright (Microsoft). Mypy, and Pyright are the most popular, but they are barely supported by PyCharm, and just don't seem to function well. Pyright works with VS Code. A lot of older libraries in Python do not have types annotated. So it's hard to fully use them in a system.

Python is hard to deploy, and hard to install. It's common for a library to install C binaries and other things. I believe Pandas is 80MB in size, and LXML is over 20MB in size. Numpy at lest 20MB. A lot of python libraries sneak in C dependencies which are hard to manage moving to ARM systems, and on different operating systems. On top of C binaries, Python need to be compiled from C to install itself. It needs to be managed as a language version, you cannot have the operating systems Python of 3.9 running the apps Python 3.12 code. The dependencies need to be installed on the server. It's hard to package a deployment, like a Jar or a binary. Poetry goes do a good job of packaged a zip though.

A lot of python packages are targeted to small scripts, and not larger programs. This means their error handling is lack luster. With a small script, or in a jupyter notebook, it is fine to just log errors. In a larger program running on a server. If a package/library logs an error, it's extremely confusing. You don't know where that function was called, and you don't get any stack trace.

Python has gone through an odd 'pythonic' phrase, and has a lot of built magic methods or special methods. Pythonic code may have encouraged the use of really dense hard to understand code like:

```
def collect_account_ids_from_arns(arns):
    return { matched.groupdict()["account_id"] for arn in arns if (matched := re.match(ARN_REGEX, arn)) is not None }
```

Good luck understanding that on your Thursday afternoon.

Pythons built in methods are a bit unintuitive. For example`__repr__` and `__str__`, could it be just a string function, and how is `__repr__` different? Could `__init__` just be `init()`? Could `__del__` just be `delete`, and`__len__` could that just be `len()`. There is also `__contains__`, `__setitem__`, `__iter__`, `__next__`, `__call__`, `__eq__`, `__exit__`, `__enter__`, `__new__` and `__add__`. These really seem there just to confuse people.

I think if you are writing code in python you should try to move to another language, unless you are building machine learning models.

ps. I called this Tooth and nail to get to the start as the start of a program is when it starts adding business value, and python makes it hard to get to this starting spot.
