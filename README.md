## workflow-service-sdk

Go module: [`github.com/soft-optimum-services/workflow-service-sdk`](https://github.com/soft-optimum-services/workflow-service-sdk)

This project is a **client SDK for the `workflow-service` HTTP API**: it wraps HTTP calls to the REST routes exposed by your workflow server (templates + execution).

It provides a Go HTTP client, dependency interfaces, and data models aligned with the JSON shape returned by `workflow-service` so you can call:

- **Template** endpoints (states / transitions / template)
- **Execution** endpoints (can/execute transition, bind, history, etc.)

### Goals

- **Centralize** HTTP consumption of `workflow-service` in a reusable module.
- **Stabilize** consumer-side contracts via interfaces (straightforward mocks in tests).
- **Avoid** duplicated logic (URLs, JSON envelopes, transition sorting/filtering, etc.).

### What this module is not (yet)

- A workflow engine.
- A server implementation.
- The single source of domain types for your consuming application (your app can keep its own DB/UI models).

### Important Go constraint: `internal`

Applications often place structs under `…/internal/…`.
In Go, an `internal/...` package **cannot be imported** by another module.

As a result, this SDK exposes its models under `github.com/soft-optimum-services/workflow-service-sdk/model` by **duplicating** the structure and JSON tags expected by the `workflow-service` API.
If you also maintain types in your app, you must keep **JSON parity** (tags, fields) when the API contract changes.

### Installation (module)

```bash
go get github.com/soft-optimum-services/workflow-service-sdk
```

For local development before publishing, add a `replace` directive in the consumer’s `go.mod`:

```go
replace github.com/soft-optimum-services/workflow-service-sdk => ../path/to/workflow-service-sdk
```

### Configuration

The client is built from a `BaseURL` (HTTP origin of `workflow-service`).
The `NewFromEnv()` helper reads `WORKFLOW_SERVICE_BASE_URL`.

Example:

```go
package main

import (
	"context"
	"fmt"

	workflow "github.com/soft-optimum-services/workflow-service-sdk"
)

func main() {
	ctx := context.Background()

	c := workflow.NewFromEnv()
	states, err := c.ListStates(ctx, "YOUR_TEMPLATE_CODE")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(states))
}
```

### Interfaces: modular dependencies

The module exposes small interfaces plus a composite interface:

- `TemplateReader`: template reads (`ListStates`, `ListTransitions`, `GetTemplateByCode`)
- `ExecutionClient`: execution (`CanTransition`, `ExecuteTransition`, `BindEntity`, etc.)
- `WorkflowQueries`: derived helpers (e.g. `ListManualTransitionsFromState`, `ResolveTransitionTargetStateCode`)
- `WorkflowClient`: composite interface (combines the three above)

Recommendation: depend on the **small interfaces** when possible, and use `WorkflowClient` only when a component truly needs the full surface.

### Covered endpoints

#### Template

- `GET /api/v1/templates/by-code/{templateCode}/states`
- `GET /api/v1/templates/by-code/{templateCode}/transitions`
- `GET /api/v1/templates/by-code/{templateCode}`

#### Execution

- `POST /api/v1/execution/can-transition`
- `POST /api/v1/execution/execute-transition`
- `GET /api/v1/execution/state/{entityType}/{entityId}`
- `GET /api/v1/execution/history/{entityType}/{entityId}`
- `POST /api/v1/execution/bind`
- `GET /api/v1/execution/transitions/{entityType}/{entityId}?role=...`

### “Queries” (helpers)

Some operations (UI, orchestration, batch) need extra logic:

- `ListManualTransitionsFromState`: loads states+transitions, resolves the `from_state` ID, filters enabled `manual` transitions, sorts by priority then code.
- `ResolveTransitionTargetStateCode`: resolves the target `State.code` from a `transition_code`.

### Code layout

- `client.go`: `Client` struct, constructors
- `http_transport.go`: HTTP GET/POST JSON transport
- `template_reader.go`: template endpoints
- `execution.go`, `execution_types.go`: execution endpoints
- `queries.go`: derived helpers
- `model/`: JSON types (Template/State/Transition/…)

### Versioning / compatibility

This module is an SDK: compatibility follows the HTTP/JSON contract of `workflow-service`.
When the API introduces breaking changes (renamed fields, altered envelopes), update:

- response envelopes in the client
- types under `model/`
