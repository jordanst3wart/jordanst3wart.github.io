---
title: Terraform Map
date: '2025-02-25'
tags: ['terraform', 'maps']
summary: Making terraform code look slightly nicer
---

Terraform has two ways to create multiple resources. There is the `for_each` option, and a `count` option. Count creates an array of resources that are number, whereas `for_each` uses the key of the key to store the state of the option against. That is `file[0], file[1]` for count, and `file[hello], file[world]` with `for_each`. It's generally a lot nicer having resources named explicitly. It's often more convenient to name things as well, even if it is just `ec2-instance-1`, and `ec2-instance-2`. Due to this I heavily prefer `for_each` loops over `counts`.

`for_each` loops have limited inputs. They allow either a set of strings, that is `toset(["hello", "world"])`, or a map object. Often a set of strings isn't ideal to create a resource, unless it's one resource you want to name.

This can look like (naming a single resource with a set):

```hcl
resource "some_resource" "foo" {
  for_each = toset(["some_name"])
  key = value
}
```

That would create a resource with the name `some_resource.foo["some_name"]`, or something close to that.

With a `count` statement, this looks like:

```hcl
resource "some_resource" "foo" {
  count = 1
  key = value
}
```

And the resource should be called `some_resource.foo[0]`.

I see a bit of an anti-pattern in terraform when someone has a list, and converts that list to a map to iterate over those resources. It looks like:

```hcl
locals {
  some_list_of_objects = [{
    name = "hi",
    size = "large"
  },
  {
    name = "bye",
    size = "medium"
  }]
}

resource "some_resource" "foo" {
  # have to convert list to map here
  for_each = { for index, value in local.some_list_of_objects: value.name => value }
  size = each.value.size
}
```

Whereas I think it is a lot simpler to just use maps instead of lists at the start, like:

```hcl
locals {
  some_map = {
    hi = {
      size = "large"
    },
    bye = {
      size = "medium"
    }
  }
}

resource "some_resource" "foo" {
  # just use the map
  for_each = local.some_map
  size = each.value.size
}
```

I think the second example is a lot simpler.

For completeness, here is an example with an if statement:

```hcl
# Define environments with their configurations
variable "environments" {
  type = map(object({
    cidr      = string
    is_public = bool
  }))
  default = {
    development = { cidr = "10.0.0.0/16", is_public = true },
    staging     = { cidr = "10.1.0.0/16", is_public = true },
    production  = { cidr = "10.2.0.0/16", is_public = false }
  }
}

# Create public IP addresses only for environments where is_public is true
resource "some_resource" "foo" {
  for_each = {
    for env_name, env_config in var.environments : env_name => env_config
    if env_config.is_public
  }

  vpc = true

  tags = {
    Name = "eip-${each.key}"
    Environment = each.key
  }
}
```
