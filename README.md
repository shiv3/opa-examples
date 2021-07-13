# opa-envoy examples 

- http
  - go server impl 
- grpc
  - go grpc server impl

```
.
├── README.md
├── component.drawio
├── grpc
│   ├── docker-compose.yaml
│   ├── envoy.yaml
│   ├── opa.yaml
│   ├── policy.rego
│   ├── server
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── grpc
│   │   │   ├── go
│   │   │   │   └── grpc.pb.go
│   │   │   ├── grpc.pb
│   │   │   ├── grpc.proto
│   │   │   └── grpcpb
│   │   │       ├── grpc.pb.go
│   │   │       └── grpc_grpc.pb.go
│   │   ├── main.go
│   │   ├── models
│   │   │   ├── boil_main_test.go
│   │   │   ├── boil_queries.go
│   │   │   ├── boil_queries_test.go
│   │   │   ├── boil_suites_test.go
│   │   │   ├── boil_table_names.go
│   │   │   ├── boil_types.go
│   │   │   ├── sqlite3_main_test.go
│   │   │   ├── user.go
│   │   │   └── user_test.go
│   │   ├── schema
│   │   │   ├── ddl.sql
│   │   │   └── models.sqlite
│   │   ├── sqlboiler.toml
│   │   └── test.sh
│   └── test
│       ├── inputs
│       │   └── jwt_verified_input.rego
│       └── policy_test.rego
└── http
    ├── docker-compose.yaml
    ├── envoy.yaml
    ├── opa.yaml
    ├── policy.rego
    ├── server
    │   ├── Dockerfile
    │   ├── go.mod
    │   ├── go.sum
    │   ├── index.html
    │   ├── main.go
    │   ├── models
    │   │   ├── boil_main_test.go
    │   │   ├── boil_queries.go
    │   │   ├── boil_queries_test.go
    │   │   ├── boil_suites_test.go
    │   │   ├── boil_table_names.go
    │   │   ├── boil_types.go
    │   │   ├── sqlite3_main_test.go
    │   │   ├── user.go
    │   │   └── user_test.go
    │   ├── schema
    │   │   ├── ddl.sql
    │   │   └── models.sqlite
    │   ├── sqlboiler.toml
    │   └── test.sh
    └── test
        ├── inputs
        │   ├── curl_ping.rego
        │   ├── curl_post_user.rego
        │   └── jwt_verified_input.rego
        └── policy_test.rego
```
