---
title: Don't use terraform workspaces
date: '2025-1-19'
tags: ['terraform', 'developer experience', 'workspace']
summary: 'Terraform workspaces are dangerous and there are better alternatives'
---

Don't use terraform workspaces. It's dangerous, and there is better alternatives. Workspaces should probably be removed from terraform.

Let's look at an example, so you think I want an exact replica of a dev environment, and prod environment with as few files as possible.

So let's have:

```
infrastructure/
├── main.tf
├── variables.tf
├── outputs.tf
└── terraform.tfvars
```

And you create two workspaces, dev and prod:

```sh
# Create workspaces
terraform workspace new dev
terraform workspace new prod

# List workspaces (* shows current workspace)
terraform workspace list

# Switch between workspaces
terraform workspace select dev
```

Once done you want a plan, and assuming everything is well apply. So:

```sh
terraform plan
terraform apply
```

Okay great, dev is done. Now, prod:

```sh
terraform workspace select prod
terraform plan
terraform apply
```

Cool all done. Now, let's say you come back, and want to do something quick in dev, and forget you left it in prod:

```sh
terraform plan
terraform apply
```

Oops... you just deployed to prod. If you forget `terraform workspace select dev`, you can easily break prod, and if you want to do some small changes in one environment it's easy to get confused.

Instead, try this:

```
infrastructure/
├── environments/
│   ├── dev/
│   │   └── main.tf
│   └── prod/
│       └── main.tf
├── modules/
│   └── base_infrastructure/
│       ├── main.tf
│       ├── variables.tf
│       └── outputs.tf
└── provider.tf
```

And use:

```sh
cd environments/dev
terraform init
terraform plan
terraform apply
```

This might be easily in a script:

```sh
terraform -chdir=environments/dev plan
terraform -chdir=environments/dev apploy
```

For prod:

```sh
terraform -chdir=environments/prod plan
terraform -chdir=environments/prod apploy
```

This approach has more files, but it is a lot easier, and less error prone.
