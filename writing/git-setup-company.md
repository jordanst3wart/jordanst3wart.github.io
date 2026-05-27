---
title: Git personal and company profile
date: '2024-07-10'
tags: ['git']
summary: How I setup git with a personal and company profiles
---

I just wanted to show how I got a company and personal git profile on the same computer.

This is my ~/.gitconfig file (take note of the `gitdir`):

```toml
[core]
  excludesFile = ~/.gitignore_global
[alias]
  squash = "!f(){ git reset --soft $(git merge-base main $(git branch --show-current)) && git commit -m \"${1}\";};f"
  squashm = "!f(){ git reset --soft $(git merge-base master $(git branch --show-current)) && git commit -m \"${1}\";};f"
  update = "!f(){ git checkout master && git pull && git checkout $(git branch --show-current) && git rebase;};f"
[push]
  autoSetupRemote = true
[pull]
  rebase = true
[includeIf "gitdir:~/company/"]
  path = .gitconfig_company
[includeIf "gitdir:~/Documents/"]
  path = .gitconfig_personal
```

So any work I do in `~/company/` uses the company git profile, whilst any work I do in `~/Documents`, uses my personal git profile.

The file `~/.gitconfig_company` is just:

```toml
[User]
  email = username@company.com
  name = company-git-username
  signingkey = XXXXXXXXXXXXXXXXXXX
[gpg]
  program = gpg2
[commit]
  gpgsign = true
```

And my personal gitconfig is at `.gitconfig_personal` and looks like:

```
[user]
   name = jordanst3wart
   email = nil@gmail.com
```

And that setup works well.
