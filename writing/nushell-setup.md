---
title: Nushell setup
date: '2024-03-03'
tags: ['nushell', 'shell']
summary: How I setup nushell
---

I decided to try out [nushell](https://www.nushell.sh/). The reason for this is I dislike bash/zsh if statements, and other similarities.

For reference here is a bash if statement:

```bash
file="file.txt"
#Checking if a file exists
if [[ -e "$file" ]]; then
  echo "$file file exists."
fi
```

Whenever I am writing a bash script, I have to pause and look up this syntax. It's just like the `tar` command.

Here is the same thing in nushell:

```nu
let file = "file.txt"
if $file | path exists {
  $"($file) file exists"
}
```

I think I can learn to write that whilst not breaking my flow.

For reference, I tried out Elvish shell as well, and I found the missing `history` command to be a deal breaker for me. You can use Ctrl-R in Elvish for history (I think), but I have to look that up every time as well.

I found installing nushell easy, with `nix` package manager on MacOs, but I had two problems:

1. adding custom commands in a modular way
2. setting up PATH
   You can see how I solved these problems below.

You can find the config directory of nushell with this command:

```nu
$nu.default-config-dir
```

The `config.nu` file is quiet long:

```nu
❯ wc -l config.nu
     743 config.nu
```

So I didn't want to edit that file, and I wanted to keep my custom commands separate. I created a `custom.nu` file in that directory, and added the following command in `env.nu` to load the custom file on startup:

```nu
source ($nu.default-config-dir | path join 'custom.nu')
```

I think a nice way to add to PATH is:

```nu
use std "path add"
$env.PATH = ($env.PATH | split row (char esep))
path add /nix/var/nix/profiles/default/bin
path add ($env.HOME | path join "Library" "pnpm")
path add ($env.HOME | path join ".nix-profile" "bin")
path add ($env.HOME | path join ".bun" "bin")
path add ($env.HOME | path join ".npm-global" "bin")
path add ($env.CARGO_HOME | path join "bin")
path add /usr/local/go/bin
path add ($env.HOME | path join ".local" "bin")
$env.PATH = ($env.PATH | uniq)
```

I think this is another nice way to adding to path:

```nu
$env.PATH = ($env.PATH |
    split row (char esep) |
    append /usr/local/bin |
    append ($env.CARGO_HOME | path join "bin") |
    append ($env.HOME | path join ".local" "bin")
)
# filter so the paths are unique
$env.PATH = ($env.PATH | uniq)
```

I also added some custom commands to `custom.nu`, like:

```nu
def search [query: string] {
  let search = "https://google.com.au/search?q=" + $query
  ^open -a '/Applications/Firefox Developer Edition.app' $search
}
```

And that's my current `nushell` setup. It is also nice to have these configuration files versioned in git. I like `stow` to symlink and version this.
