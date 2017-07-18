# Rest Api with Authentication, Payload validation, Push notifications and embedded Documentation

This library is a wrapper around [Gin Gonic](https://github.com/wearkinetic/gin) that provides the following built-in features:
* Build an API with `REST` in mind
* Authentication at the endpoint level, and not at the resource level, using `JWT` tokens.
* Push notifications through WebSockets
* Embedded API documentation

## Endpoint structure

Every new API endpoint has the following structure:

### Resource
Because we have `REST` in mind, the URIs should be built around resources. A resource can be defined as an atom of the application logic. In most cases it will represent a specific entity represented in the database, whether it expresses an idea by itself, or is a relationship between multiple ideas.

**Example**
Imagine a database storing soccer players, teams, and soccer player agents. We can identify three resources which are `/players`, `/teams` and `/agents`. You would access methods centered around each of these resources using the previous URI. Now, what if the database models the relationship "An agent negociated the contract between a player and a team". That's a three-way relationship, and we could represent it using `/agent/AGENT_ID/player`, or maybe `/agent/AGENT_ID/team` etc. We decide to represent this relationship as a resource itself, by using the URI `/contract` for instance. It gives room for more complexity as this relationship gets more complex (imagine we'd like to save in which country the contract was negociated).

Also, every resource needs to have a description. The program will compile but not run otherwise.

### Verbs
Again, because we have `REST` in mind, once we have decided which resources are made available by this API, we can implement the four `HTTP` methods (or verbs):

* `GET`, to retrieve the resource, based on `GET` parameters (such as `?hello=world&hola=mundo`). Idempotent, cacheable.
* `POST`, to create a new resource, based on the request body (we recommend using `JSON`).
* `PUT`, to update a given resource, based on the request body. Idempotent.
* `DELETE`, to delete a given resource, based on the request body.

We allow "multiple `GET` methods" to be implemented, each of them having a different name, encoded in the `_n` `GET` parameter.

**Example**
`/players?_n=all` would return all players.
`/players?_n=team&teamID=xyz` would returns all players for team `xyz`.

Depending on the `_n` get parameter, the API will figure out which **action** the API user wants to do. Let us explain the concept of **action** in this library.

### Actions
An action represents what the user wants to do, given a resource, and an HTTP verb. Each action need to implement the following steps:
1. Payload validation
2. Authentication validation
3. Handler (the actual thing the user wants to do, if he gets access, with the right payload)

In our examples above `/players?_n=all` and `/players?_n=team&teamID=xyz`, this allows us to fetch:
* all players at once, and this should be only accessible for a super-user, managing all teams
* just a specific team's players, which should be accessible not only by a super-user, but a team manager too.

Putting authentication at the action level enables complex authentication defined depending on what the user is trying to access for `GET` methods. For `POST`, `PUT` and `DELETE` the authentication is at the resource level, but can be custom to the payload sent in the request.

Also, every action needs to have a description. The program will compile but not run otherwise.

## Documentation

The endpoint `/doc`, with the appropriate authentication, called with `GET` method, will give you a documentation of the endpoints accessible to you, resource by resource, and action by action for each of these resources.


## Nota Bene

* Request payload and body are used to mean the same thing.
* HTTP verb and method are used to mean the same thing.
