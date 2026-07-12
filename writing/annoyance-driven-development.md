---
title: Annoyance driven development - Oblivion
date: '2026-07-12'
tags: []
summary: Do I need this?
---

I've been thinking of working on a new side project that delete things. I've been deleting the old git branches on github of some repositories I work in at work. It's annoying. Colleagues don't delete their branches once they merge in their code, leaving hundreds of useful branches hanging around.

At first I just deleted the branches in a fun frenzy... but now it feels lame. It feels a bit annoying. The fix for this is to enable 'Automatically delete head branches' in the admin console of the git repository, or for the developers to just delete their own branches when they merge, but I don't have permission to do that. So I just have to manually delete them in the UI, or via the console. Someone wrote a git action to delete branches that were older than 6 months, but I feel like that is more maintanence.

This is why I want to create a tool that helps with the deletion of dead software artefacts. I want to call it oblivion. As in it moves the dead code/software/data to oblivion. I just don't have the hubris, or passion to knuckle down and write hundreds or thousands of lines of code on a another side project... but I do have the annoyance, and disgust though.

I feel like to write a tool to delete things is to admit you are coming in second. You are not the perfect solution. The perfect solution is to just have not created the thing to begin with. Or to have deleted it in a more timely manner. Or to have blocked it from being created.

For instance, when I shop online, and a parcel is posted with Fedex I get at least four emails:
1. Your parcel is being prepared by the sender
2. Your parcel from Federal Express Australia is on its way
3. Your parcel from Federal Express Australia is ready for collection
4. Your parcel from Federal Express Australia has been collected
In the idea, of having a notifiation per action, I only need one email "Your parcel is ready for collection", or "We lost your parcel sorry - giving you a refund". The ideal solution is for AusPost to just not send three emails. I have an app where I can check the status of parcels anyways if I need to know. The second best solutions is to filter out these emails after they have been sent which hopefully oblivion will be able to do. Ideally, this should happen before I get a notification of receiving a new email, but that might not be possible - so it might even be the third best solution.

Maybe this needs to be pushed up, and told to people, that is all your users delete this - so added a dumb feature. I don't know. I just don't feel like this can be worked on with the same enthusiasm as other projects.
