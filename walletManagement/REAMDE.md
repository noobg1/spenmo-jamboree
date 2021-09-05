## WalletManagement service

### Usage

```
# go version 1.15 and above
export MONGO_PASS=XXXXX # mandatory
export ENABLE_RATE_LIMITING=true # optional
export REDIS_PASS=XXXXX # optional
./kickstart.sh # Or go run server.go
```

### Tests

```
go test ./coreDomains/cards/  -v
```

#### Examples

```
 # Get cards 
 curl 127.0.0.1:8080/cards
 # Create a card
 curl -d '{"name": "travel", "walletType": "team", "walletId": "6134895c65039b0e7bd841bc", "dailyLimit": 12, "monthlyLimit": 20  }' 127.0.0.1:8080/cards
 # Delete card
 curl -X DELETE 127.0.0.1:8080/cards/613488eeaa19c77af4baf816
```

### Docs

```
# swagger doc
http://localhost:8080/swagger/doc.json
```

### Design/ Approach

1. DDD(Domain Driven Design) - each domains (users, wallets, cards etc) have clear boundaries
2. Layered flow of data: one way street => `controller -> service -> repository -> dataSources`
3. Folder/ file structure
```
.
├── REAMDE.md
├── common
│   ├── appError.go
│   └── constants.go
├── coreDomains
│   ├── cards
│   │   ├── controller.go
│   │   ├── controller_test.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── repository_todo.go
│   │   ├── service.go
│   │   └── service_test.go
│   ├── users
│   │   └── model.go
│   └── wallets
│       └── model.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── kickstart.sh
├── middlewares
│   ├── rateLimiter.go
│   └── requestTracer.go
├── server.go
├── utils
│   └── db.go
└── walletManagement

```
4. Tests next to respective files
5. Each domain is one package ex: `cards` domain means translates to `cards` package
6. Each domain contains it's one model, controller, service, repository files
7. Further the domains can be broken down to subdomains as code grows
8. Communication between domains strictly by calling `Service` layer
9. REST endpoints (structured to be extensible for grpc, event driveb/ pub-sub, graphql)
### TODOs in mind

Functional requirements:
Crud APIs
- [x] Request validation
- [x] Idiomatic response codes
- [x] `/cards` GET 
- [ ] `/cards/:d` GET
- [x] `/cards` POST
- [ ] `/cards` UPDATE (got throttled by mocking lib issues)
- [x] `/cards` DELETE
- [x] controller unit tests
- [x] service unit tests
- [ ] repo unit tests (stuck with mocking lib issues)
- [x] Integration tests (is as simple as removing mocks in controllers and setting env vars)
- [x] Middleware to set contexts (traceability by picking from request params(tracer id, auth headers, idempotency key etc) and passing on)
- [ ] Middleware tests
- [x] Integration tests (is as simple as removing mocks in controllers and setting env vars)
- [x] Traceability/ open tracing (implemented using context API to be passed as args for chain of function, method calls)
- [x] pagination for `/cards GET `
- [x] API doc using swagger (param descriptions, response detailing, disabling in certain envs to be enhanced)
- [x] Ratelimitting using `ulule/limiter` via middelware with redis as datastore
- [x] Design catering no circular dependency
- [x] 12 factor app
- [x] Singleton db connection pool reuse
- [x] Domains boundaries are clear to be extensible
- [x] Global error handling/ recovery using middleware (to be improvised for custom errors) 

Note:
- Quite struggled with mocking in tests, still there are gaps
- Tried balancing between delivering and learning :P This was fun
