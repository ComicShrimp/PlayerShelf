# Architecture

The Goal is to use Clean Architecture separated by domains. Every domain will have all the layers of the clean Architecture (app, domain, infra).

The layers are:

- App (app): Store interfaces to be implemented in infra and have the use cases
- Domain (domain): Have the entities and value objects
- Infra (infra): Have the implementations with the external world
