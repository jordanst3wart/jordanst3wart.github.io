---
title: Why open banking failed in Australia, but succeeded in the UK
date: '2025-08-16'
tags: ['open banking', 'cdr']
summary: They didn't use FAPI v1 standard, and added a consent model
---

tldr: the regulatory's pushed for a bespoke OAuth stand and didn't use FAPI v1, making it harder for start ups to use the standard.

I rarely think big complex things, have simple solutions, but this is a rare case I think that does have a simple solution. I think open banking in Australia would have been a success if the open banking standard had just used the FAPI standard, or the financial API standard for OAuth, instead of making up a different standard.

Here is the signs of open banking has failed in Australia:

- with open banking apps the login process commonly fails, and you have to re-consent every 12 months
- the number of third party providers using the data is abysmal

With the login process: I was using the Frollo budgeting app. It was the first app to get open banking working.

With the following it:

- commbank -> works
- westpac -> fails
- citi credit card -> fails
- up bank -> works

It works for about 50% of the data holders I tried, which is a bit embarrassing.

With the number of third party providers in Australia's opening is:

- 12 surrendered or given up
- 29 active
- 9 accredited

Whereas the UK has 229 fintech, or third party providers, 10 times the amount.

Reference:
Australian Third party providers:
https://www.cdr.gov.au/find-a-provider?page=1&providerType=Data%2520Recipient&status=SURRENDERED
UK Third party providers:
https://www.openbanking.org.uk/regulated-providers/?filter-provider-type=third-party-providers

So open banking in Australia failed, but openbanking in the UK succeeded.

In the UK, the open banking standard used FAPI v1, or stayed very close to common standards. It highly valued not "re-inventing the wheel".
reference:
https://openbankinguk.github.io/read-write-api-site3/v4.0/profiles/read-write-data-api-profile.html#standards

This allowed for easy adoption from fintechs. The FAPI standard is well known and easy to understand. It's a well known extension of OAuth.

Whereas Australian open banking re-invented the wheel. It made a "consent model", or consent standards. It adds an entirely new entity to the OAuth spec `cdr_arrangement_id`, and the concept of consent. Normally in OAuth, if you have an active access token you can access data. With Australia's open banking, that logic becomes more complicated. It becomes if you have an active access token, and a valid consent, then you can access data. It's slightly more complicated than that with scopes, and claims, but that's not needed here.

So normally with OAuth you have an access token you can check is signed, and has not expired. You have a refresh token, which can re-generated an access token. With openid connect you have an id_token, which has additional user information. The OAuth server doesn't need to persist much information as it can dynamically check access tokens and refresh tokens are signed based on the keys it has (or json web keys JWKs). With consent added via an arrangement id, it makes every interaction just slightly more complicated.

Instead of the OAuth server (or Authorization Server) not needing to store data, now with every transaction the consent id needs to be checked as valid. This means the consent object needs to be stored and accessed by the Authorization server on every request. This adds another level of complexity on top of OAuth, which is not apparent anywhere else.

The FAPI (or financial API spec) does add in mutual TLS, but that is very common in security conscious applications.

The added layer of complexity prevented a lot of fintechs from joining the already regulatory heavy eco-system in Australia, and has resulted in only a few fintechs taking part in the eco-system. A lot of the startups looking to use open banking ever did not manage to implement the Australian open banking spec, or had more success using legacy web scraping technology.

Adding in the consent field would have looked like a bright idea to the regulatory, but in reality adding additional requirements is rarely a good idea.

I believe consults where hired to help the startups succeed in open banking, but when you have a team of one or two engineers, getting an email from a consultant asking for a status update is a waste of time.

Reference:
https://consumerdatastandardsaustralia.github.io/standards/#consumer-experience_consent-standards

The open banking specification in Australia was more technically complicated than what it needed to be, and didn't stick to open or familiars standards. That is way it failed in Australia, but succeeded in the UK.
