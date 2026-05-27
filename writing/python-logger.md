---
title: Python Logger Gotcha
date: '2024-02-18'
tags: ['python', 'logger', 'prototype pattern']
summary: Python logger using the prototype pattern
---

I was having some issues with the logging module in python, so I decided to write a small note about it.

The python logging module uses oddly uses prototyping to create loggers at the module level. It can make it confusing to understand how to the root logger applied to all loggers, and how a prototyped logger does not have this affect.

But normally with prototyping once instantiated the prototype should not be impacted by the changes to the original object. But with the logging module, it's the opposite. The prototype is impacted by the changes to the original object.

So you should see this:

```python
logging.getLogger(__name__)
```

This prototypes a logger from the root logger for the python module, creating one if none exist.

It gets confusing when you want to change the logger, and

```python
# let's set the logging level to info
logging.basicConfig(level=logging.INFO)

# let's get a logger for the module
loggerMain = logging.getLogger(__name__)
loggerMain.debug("This shouldn't log as it's debug level, and the level is set to info")  # it doesn't log
# `logging.getLogger()` is confusing as, it gets the root logger
# it's the same as `logging.root`
loggerRoot = logging.getLogger()
# if we set the level of the root logger to debug, then the debug message will log
# is it is using the prototype of the root logger
loggerRoot.setLevel(logging.DEBUG)
# now let's log a debug message again
loggerMain.debug("This logs now...")  # it logs
# if we set the level of the main logger to info, then the debug message will not log
loggerMain.setLevel(logging.INFO)
loggerMain.debug("This doesn't log now...")  # it doesn't log
```

So if you want to change root logger and have that applied across logger level, you should do it like this:

```python
rootLogger = logging.root
# make changes
rootLogger.setLevel(logging.DEBUG)
# have them applied across all loggers in all modules
logger = logging.getLogger(__name__).
```
