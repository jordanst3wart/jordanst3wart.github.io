---
title: Result of rewriting the website
date: '2026-06-06'
tags: ['templ', 'react', 'Typescript']
summary: 'Moved to simpler setup'
---

So I rewrote the blog from scratch. I used templ, and go to compile static html pages, and tailwind css v4 for the styling. The static html pages are then served with a web server.

[![Some artwork](/silent_paths_of_japan_by_bogi380.jpg "Silent Paths of Japan by Bogi380")](https://www.deviantart.com/bogi380/art/Silent-Paths-of-Japan-1342210952)

This is the result:
- total lines of code from `1839` to `517`
- number of dependencies from `1511` to `118` with depth of 2 for the old codebase, and about 50 dependencies coming in from `air` cli tool in go
- time taken from '2026-05-17' to '06-06-2026', I haven't added tags yet though, and it is intentionally very minimistic, with more effort on styling. It was less than a month.
- number of config files from `10` to `2`

So the number of lines of code was reduced by 4x... and the stewart.bot codebase was already quiet minimistic anyways. It would of been very hard to have reduced the typescript any more. The number of dependencies in `air`, the go cli tool, is surpising, removing it would mean the dependencies look like roughly like going from `1511` dependencies to `66`. Even with `118` dependencies though that is roughly 13 times more dependencies, which might mean 14 times less security vulnerabilities, and 14 times less time upgrading dependencies. I might took to replace air with a generic command line runner. It just watches files, and runs cli tools. I could just replace it with a makefile or justfile.

I think this exercise was worth it. I went from having a broken next.js build, and not being able to publish blogs, to having a lot more minimistic styled "writings" website. I learnt a lot about styling blogs, and not just have a generic webpage. I found it interesting comparing my blog website to other dev blog websites with similar domains as me, such as [stewart.io](https://stewart.io) - minimistic and nice, and [stewart.dev](https://stewart.dev) - boring and plain consulting site. I didn't think i realised how many boring generic blog website there are. Here is another two nice blog website:
- [k10s](https://blog.k10s.dev/im-going-back-to-writing-code-by-hand/)
- [Mo Bitar](https://atmoio.substack.com/p/after-two-years-of-vibecoding-im)
- [Bryan Cantrill](https://bcantrill.dtrace.org/2026/04/12/the-peril-of-laziness-lost/#)

Both of those actually use blog templates, or 3rd parties.

### Calculations
#### Number of lines of code in new codebase
- sitemap.go 58
- cmd/server/main.go 21
- main.go 217
- main_test.go 36
- blog.templ 185
(blog_templ.go is generated)
total 517 (4 files, 2 are not strictly needed to build the website)

#### Number of lines in old codebase
- app: 655 lines
- components: 485 lines
- data: 36 lines
- layouts: 663 lines
Total: 1,839 lines


#### Number of dependencies
#### new codebase
27 in my go.sum file
109 in [air](https://github.com/air-verse/air/blob/master/go.sum)
24 in [templ](https://github.com/a-h/templ/blob/main/go.sum)
160 in total, filtering out duplicates means there is 118 dependencies

Or:
```sh
go mod graph | wc -l
66
```

#### old codebase
```sh
pnpm ls --parseable | wc -l
946
```
Or including all dependencies:
```sh
pnpm ls --parseable --depth Infinity | wc -l
2005
```
Depth two seems reasonable:
```sh
pnpm ls --parseable --depth 2 | wc -l
1511
```



#### Number of config files
.air.toml
go.mod
2

Old
jsconfig.json
next.config.js
next-env.d.ts
package.json
pnpm-workspace.yaml
postcss.config.js
prettier.config.js
tsconfig.json
tailwind.config.js
contentlayer.config.js
10
