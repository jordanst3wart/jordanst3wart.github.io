---
title: Wordpress dilemma
date: '2025-02-15'
tags: ['technology', 'wordpress']
summary: Recommending software to people who are not software engineers
---

I have this dilemma. It starts with a friend who has no technology skills ask me how they can setup a website. I think if I was doing it I would write HTML, CSS, and JavaScript, as these technologies are what a webpage is, but for them I think oh, I guess you can use a third party software platform like Wix, or Wordpress. However, i know if, or when, they want to do something fancy it would be a lot harder for me to help them out, and they are creating a dependency on a huge ecosystem that would be hard to move off.

Their journey generally looks like:

1. buy a domain (difficulty: easy)
2. create a template website (difficulty: easy)
3. try to change website to do the thing they want (difficulty: medium to hard)

It's generally in the changing the website to be exactly how they want were it gets really hard. They obviously don't know how to style a website, or even the important web concepts like a "call to action". It's this oddness of creating a website is extremely easy, but changing it is extremely hard.

There is also this huge dependency burden they take on. Wordpress may need to be upgraded every year. The plugins that they use may become unsupported or have security vulnerabilities. If something goes wrong, how do you debug a Wordpress site for someone non-technical? I think these are issues you quickly run into if you are trying to do something that isn't perfectly cookie cutter.

On the other hand, do I recommend then to use react, and go into using npm packages. They may need to update there version of react, and other npm libraries every year. It would take them a lot longer to get started, and they would have to learn a lot for something is potentially not that strategy for their business.

I could recommend they just use html, css, and javascript, to avoid all unnecessary dependencies, but a that stage it might take them even longer to get a basic webpage setup.

I see this happening in the worst way at companies. So you have a company that is heavy motivated to save money, and deliver a project fast. So instead of hiring software who knows HTML, CSS, and JavaScript, or even hiring someone a software engineer that knows React. They look at wordpress, and they quickly generate a wordpress sight. This allows them to capture a space in the short term, but in the long term it really slows them down. They them struggle to update the website as instead of having a frontend engineer, who has a little react knowledge, they struggle to update the php plugins driving their website. It's a lot more difficult for a frontend engineer to engage with a wordpress website, than it is for a frontend engineer to engage with a react website.

There is this secondary problem of the frontend engineer now. If they update the website in an odd way, it might make the website less "vanilla", and more customised. If wordpress offered a custom incuberating feature to make websites load two times fast with static page generation, and your boss wants to dive user engage, you are motivated to use this incuberating feature knowing if will more than likely get changed. Whereas if you have svelte, it's already optimised for fast page reloads. If you use this new incuberating feature, when the next engineer comes along, and trys to update wordpress, they could easily get an error with something like "wordpress static site generation has changed". So it motivates subsequent engineers to try to keep wordpress as vanilla as possible.

It seems to end up in a contradiction where you need a really skilled engineer to not just do the simple code change, but they do complicated code changes in some opensource codebase, or opensource tool, that might not completely fit your use case.

The company I work has setup Librechat as a large language model chat bot. Librechat is basically an openai chat clone that just calls the model from Azure, or from Claude. The codebase is in TypeScript, and there is python code for the RAG server (Retrieve augmented something). When I think of a company adopting this tool, I kind of immediately think "why not just pay for the SAAS platform?". It saves the maintenance, and setup costs. I them think given the codebase is TypeScript updating the dependencies might be a pain. There is actually 2703 versioned dependencies in the package-lock.json file for the codebase (https://github.com/danny-avila/LibreChat/blob/main/package-lock.json), which is a bit wild. If any one of those dependencies has a severe vulnerability the engineers maintaining it are probably on the hook for that, and because it is an open source codebase there is only two engineers looking into it.

Whilst there is two engineers responsible for editing the codebase, there is potentially five managers, and directors who want to get credit for the project, which is at odds. As managers, and directors, also want to have a lot of people under them. If you are a manager, and you manage one person, you must not be that important. Managers, and directors need to look important, i think makes them want to hire out sourced workers from India, from Tara Consulting Services (TCS). I think as a manager I would look a lot more important if I had a team of 20 "engineers", rather than just having two engineers, who would potentially be more productive. At this moment, there is no doubt in my head, that three well skilled, and suited engineers can outperform 20 less skilled engineers.
