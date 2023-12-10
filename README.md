# learn-clean-architecture

- [Clean ArchitectureでAPI Serverを構築してみる #Go - Qiita](https://l.pg1x.com/1PcLM6zT7HkYVGiH8)
- [hirotakan/go-cleanarchitecture-sample: Example of a Go application using a clean architecture.](https://l.pg1x.com/8EdLZ6dnP7jHvqMt7)

- [ ] module周りで詰み

# tree
```
.
├── README.md
└── app
    ├── domain
    │   └── user.go
    ├── go.mod
    ├── go.sum
    ├── infrastructure
    │   ├── router.go
    │   └── sqlhandler.go
    ├── interfaces
    │   ├── controllers
    │   │   ├── context.go
    │   │   └── user_controller.go
    │   └── database
    │       ├── sqlhandler.go
    │       └── user_repository.go
    ├── server.go
    └── usecase
        ├── user_interactor.go
        └── user_repository.go
```
