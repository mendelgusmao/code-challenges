# ZAP CHALLENGE APP

This is a simple API server to filter realties between realty portals

## Disclaimer

I'm very concerned with code quality. Not just about what is running, but what is
being pushed to the codebase and how. My way of pushing code is to create context-aware and concise commits. For the first sight the reader can be a little surprised about how such an codebase can run with so few commits, but the truth is that to achieve the conciseness I want, I became a fan of ammending and rebasing. Every single commit is in fact a collection of squashed commits made in different periods of time and the result of several `rebase -i` runs.

I agree that this only works well for single-maintainer codebases, although. Rebasing requires history rewriting and a lot of `git push --force` . Forced pushes are a bit harsh in this context and can lead to great problems if the team is not **very** aware of what is happening.

Also, there's the point of **team visibility** in which we can see the other progressing in the implementation of a feature or a big refactoring. But it kinda conflicts of the idea of conciseness because sometimes one would change a line, a word or even a letter and call it a commit, right? No worries! One should agree that if the extension of the modification allows, a squash inside a branch right after a pull request's approval is the way to achieve both *visibility* and *conciseness*. Thus, a beautiful history that makes sense when one is learning, evolving or hunting the origin of bugs is achieved.

## Architecture

The application consists of:
  * a backend server that downloads a JSON from EC2 and provides a REST API

## Testing

Go has a very simple and nice native testing framework for unit testing

Give it a try, at the root of the project, execute:

```
go test -v ./...
```

## Building & Running

#### Native Go
```
go run backend/main.go
```

#### Docker
```
docker build -t zapchallenge .
docker run -p 9091:9091 zapchallenge
```

## Querying the API

#### Listings

* **GET /listings/{portal}[?page=&size=]**

  Acquires the listings from source and filters it according to `{portal}`, which can be `zap` or `vivareal`.

  The parameter `page` is any number greater than 1 and defines the page number for the paginator while `size` defines the amount of listings per page and can be any number between 10 and 100.

  Parameters:
    - **page**: should be any number greater than 1 and defines the page number for the paginator. Default: 1
    - **size**: defines the amount of listings per page and can be any number between 10 and 100. Default: 10

  Responses:

  * *200 OK* with a *application/json* body if successful
  * *404 Not Found* if there's no such portal
  * *500 Internal Server Error* if there's an unexpected error

  Body:

```
  {
    pageNumber: int32,
    pageSize: int32,
    totalCount: int32,
    listings: [
      {
        "id": "some-id",
        ...
      }
    ]
  }
```

## Environment variables

* **ZAPCHALLENGE_ADDRESS**
 - The network address the server will be listening to (default: **:9091**)
* **ZAPCHALLENGE_SOURCE**
 - URL pointing to the source data
* **ZAPCHALLENGE_PORTALS**
 - Path to the YAML file that contains the business rules (default: **portals.yaml**)
