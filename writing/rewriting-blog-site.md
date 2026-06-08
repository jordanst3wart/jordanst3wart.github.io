---
title: Rewriting the website
date: '2026-05-17'
tags: ['next 16', 'react', 'Typescript']
summary: 'Moving to simpler setup'
---

I recently wrote a blog, before publishing the new blog the build broke. I haven't updated my website in a while. But admitterently it is just my personal blog site, so it isn't a critical piece of software infrastructure. I'm a bit confused why the vercel build just suddenly broke. I think it might have automatically bumped my dependencies or something, which would be weird because the dependencies that I have are pinned. Maybe it was some security concern.

So I tried to just bump the npm dependencies, but apparently going to next 15 to 16 is a big upgrade. I thought I just did a big upgrade previously with react server components, or something... When picking this website I wanted to be as simple, and as standard as possible, using the industry standard react, and next. I felt like that would make it easy to handle, so I could focus on just writing. It seems I was wrong.

From first using this blog template I was mortified to find that where was a lot of tiny dependencies the author wrote, and put into separate dependencies, before upgrading the dependencies in a hard to upgrade way, such as [pliny](https://www.npmjs.com/package/pliny?activeTab=readme), and [contentlayer](https://www.npmjs.com/package/next-contentlayer2). **I have probably spent a few hours trying to remove this dependencies, and failed.** The code is mostly a lot of tiny utility type features, which are easy to write, but tracing the types, and the code across dependencies is a pain. Particularly when there is a lot of small functions.

There is also the issue of having **48 security vulnerabilities**. I have 26 dependencies, I tried to remove more, and tried. These security vulnerabilities are mostly high. Apparently, I closed 41 security vulnerabilities. This is tiny blog website. How could it possibly have 48 security vulnerabilities!

This is wild. I also get confused by the build process, and the postcss processors, and other silly things. I have maybe 10 config files?! `.eslintrc.js`, `contentlayer.config.ts`, `jsconfig.json`, `next-env.d.ts`, `next.config.js`, `package.json`, `postcss.config.js`, `prettier.config.js`, `tailwind.config.js`, and `tsconfig.json`. Not including the ignore files...

So I'm moving the entire site, and maybe doing a rewrite. Let's see. Instead of following the most common technology I might try:
As few dependencies as possible - so maybe not js
As simple as possible - static site...
Using fundamental web technologies if possible... html, js, css, so maybe not something like elm...
I also don't want to learn too many new things... so not rust.

I could potentially look at using bun, with something like solid.js... I could look at using go with templ, or go template rendering... I could use htmx or something too... I could use svelte, but I feel like it borders on too much learning. I might look at some sample projects, and see how complicated they look...

Here is a sample [go templ project](https://github.com/a-h/templ/tree/main/examples/blog).
It's very simple, but it worries me on how to handle js, and css. I guess I can build it automatically, and just test it... It has a risk of not having the best DX (but neither does anything in node.js really).

I was having a look at a building a static site with svelte, and it seems like a lot. I think people urrrr towards complicating a blog. I kind of want to keep it simple like: [Mitchell Hashimoto's blog](https://mitchellh.com/). I also like the style of terminal.shop (via ssh), how it uses keys to navigate... I might just use go.

I think I could really optimise the webpage - doing things like loading the page in the background - but I don't think that is worth it for my use case.

I guess there is two steps:
1. build the webpages (I don't need to build at runtime)
2. serve the webpages from the file system (which could be done with nginx, or anything...)
I probably want to serve the pages from a go process...
