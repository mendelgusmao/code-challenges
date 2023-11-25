# Stark Bank Code Challenge

This is the final product for Stark Bank's Webhook challenge.

It's a monorepo using Domain Driven Design, so the main parts - **invoice issuer** and **invoice webhook** - can share some common components and make deployment and running easier.

## Challenge questions

#### How much have you learned during the trial?

I took the opportunity to practice and learn a bit more of **Domain Driven Design** - extensively used in the project, as well as some **Docker** and **Go** specific traits. Also, reading the **Stark Bank API** documentation and SDK code was pretty fun. I planned to deploy it using **Terraform**, but as I already have a home server running **Docker**, I found it faster and easier to use **docker-compose** locally. Also, I finally experimented with **Tailscale Funnel**. **Tailscale** is my preferred VPN and its **Funnel** service allow services running in my machines to be accessible from the external world.

#### How well does your code run? Did you create unit tests or did you leave bugs hidden in your code?

I think it runs pretty OK for the use case. I had not much time to write tests, but I'd have no trouble writing them as the code is designed to make testing easier. There are probably some bugs derived from unthought or untreated edge cases.

#### How much did you rely on us to help with technical issues you face? Less is better!

I relied on the team only for getting the sandbox credentials. The API documentation and SDK code were pretty much enough for finishing the task.

#### How readable and efficient is your code?

It uses **DDD** and has components whose implementations are very close to the Single Responsibility Principle. It has some structural overhead but the code inside every component is very clear and straightforward.

#### How quickly did you deliver the finished task? Good code is better than quick code, but if you can score on both ends, we will be impressed.

I made it in about 6 days, totalling 14 hours of work, I guess.

## Running

### Configuring

* Put a private key named `privateKey.pem` under `infrastructure/certs`
* Update the `.env` files under `infrastructure/env`. The `*_PROJECT_ID` env vars should contain your project id.

### Running with Docker

This project's `Dockerfile` is a multistage build with multiple image output, so the invoice issuer and invoice webhook can be built at the same time but run each in its own container, as we normally do using a microservices architecture.

Make sure you have **Make**, **Docker** and **docker-compose** installed in your machine.

To build and run the services, run:

```
make run
```

## Some considerations

Doing the task was pretty fun and challenging! **Kudos to Stark Bank's tech team for designing a challenge that sounds pretty close to a day doing real work**.

I had some difficulties with outdated documentation. For example: in the website documentation, `event.Parse` returns an object representing the callback event. But, after having some runtime errors when trying to unmarshal data into an Event object, I discovered that `event.Parse` actually returns the full JSON string. This made me write a workaround to parse this JSON inside a component that in principle shouldn't deal with JSON. 
