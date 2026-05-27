---
title: Odd Things with terraform
date: '2024-10-22'
tags: ['terraform', 'dsl']
summary: Somethings that I find odd with Terraform
---

I've been using terraform for a bit, and in general it's good... but there is some things that I find very rough around the edges.

These are:

- control structure `for_each` and `count`, missing (missing `if`)
- providers not allowed to use `for_each`
- backend state not being parameterised
- looping over lists has to be a set
- lack of tfvar (HCL) support in other languages
- templating is weird

### Control structure `for_each` and `count`, missing (missing `if`)

Sometimes it's just nice having code like:

```golang
if flag {
  # do stuff
}
```

It's really nice, and in terraform you have to do:

```hcl
module "servers" {
  count  = var.flag ? 1 : 0
  source = "./server"
  servers = 5
}
```

This just isn't as nice.

Count is also nasty. Let's say I have this:

```hcl
module "database" {
  count  = 5
  source = "./database"
  size = "large"
}
```

And I want to remove the 2nd database, but keep the 5th. It's hard to do (you have to use imports). So I try to always use `for_each`.

### providers not allowed to use `for_each`

This is basically just [this](https://github.com/hashicorp/terraform/issues/19932) github issue open since 2019.

It's basically just isn't allowed:

```
provider "aws" {
  for_each   = var.region
```

And this would be really handy when you need to declare a provider for each AWS region, like when you are using AWS guardduty in multiple regions.

So instead of just having a list of providers you have to have something like:

```hcl
provider "aws" {
  region = "us-east-1"
}

provider "aws" {
  region = "us-east-2"
}

provider "aws" {
  region = "us-west-1"
}
# and on and on...
```

### backend state not being parameterised

So with terraform you have things like:

```hcl
resource "aws_instance" "web" {
  ami           = var.ami
  instance_type = var.instance_size
}
```

And then you suddenly get to backend state, and it looks like:

```hcl
terraform {
  backend "s3" {}
}
```

... which is confusing.

You would expect this:

```hcl
# This will NOT work
terraform {
  backend "s3" {
    bucket = var.state_bucket  # This is not allowed!
    key    = var.state_key     # This is not allowed!
  }
}
```

But instead you have to do something like this:

```sh
terraform init \
  -backend-config=backend-config/sandbox.tfvars.json
```

And pass in the variables. It's just a completely different way of doing things compared to the rest of the DSL.

### looping over lists has to be a set

This one is odd. With the terraform for_each is works on either a set of strings (not a list), or a map.

So this is invalid:

```hcl
variable "sandboxes" {
  type    = set(number) # not strings
  default = [1, 2, 3]
}

resource "aws_instance" "web" {
  for_each      = var.sandboxes # fails as not strings
  ami           = var.ami
  instance_type = var.instance_size
}
```

So is this:

```hcl
variable "sandboxes" {
  type    = list(string) # it's a list!
  default = [1, 2, 3]
}

resource "aws_instance" "web" {
  for_each      = var.sandboxes # fails as a set
  ami           = var.ami
  instance_type = var.instance_size
}
```

But there is also no way to do:

```hcl
variable "sandboxes" {
  type    = set(string)
  default = ["1", "2", "3"]
}

resource "aws_instance" "web" {
  for_each      = [for sandbox in var.sandboxes : sandbox if sandbox != "3"] # cannot do if statements here
  ami           = var.ami
  instance_type = var.instance_size
}
```

But this is fine:

```hcl
# using a map
variable "sandboxes" {
  type = map(object({
    index = number
  }))
}

resource "aws_instance" "web" {
  for_each = {
    for name, sandbox in var.sandboxes : name => sandbox
    if sandbox.index != 3
  } # can do if statements with maps... but not lists... :?
  ami           = var.ami
  instance_type = var.instance_size
}
```

These above examples, kind of make the DSL look really half-baked.

The map syntax is a bit confusing too, like what `name` is in the above example is a bit confusing.

### lack of tfvar (HCL) support in other languages

Enough said... tfvars are

Like there is this 250 star python library on github for python:
https://github.com/amplify-education/python-hcl2

And this struggling java library:
https://github.com/bertramdev/hcl4j

If you compare that to the serialization of json, and yaml, it is incredibly lack luster.

### templating is weird

So with Ansible they use the common, and vanilla jinja library, which is really easy to use both in ansible templates, and from python.

Go (or Golang) also has a really easy to use templating library (https://pkg.go.dev/text/template), which is used a lot. But terraform templates decided to have a half/half approach, and be mostly go text templates, but add

So it looks like this:

```tftpl
{
  "port": "${addr}:${port}",
}
```

When it could of been generic go templates as this:

```
{
  "backends": "{{.addr}}:{{.port}}",
}
```

This has the file extension of `.tftpl`, which is a bit odd as well. It's a next of not invented here.

As you can see terraform has some odd characteristics in a few ways.
