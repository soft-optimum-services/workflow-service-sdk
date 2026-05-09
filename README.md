## workflow-service-sdk

Module Go : [`github.com/soft-optimum-services/workflow-service-sdk`](https://github.com/soft-optimum-services/workflow-service-sdk)

Ce projet est, pour l’instant, un **SDK client orienté API `workflow-service`** : il encapsule les appels HTTP vers les routes REST exposées par votre serveur workflow (templates + exécution).

Il fournit un client HTTP Go, des interfaces de dépendance, et des modèles de données alignés sur la forme JSON exposée par `workflow-service` afin de consommer :

- les endpoints « template » (states / transitions / template)
- les endpoints « execution » (can/execute transition, bind, history, etc.)

### Objectifs

- **Centraliser** la consommation HTTP de `workflow-service` dans un module réutilisable.
- **Stabiliser** les contrats côté consommateurs via des interfaces (mocks simples en tests).
- **Éviter** la duplication de logique (URLs, enveloppes JSON, tri/filtrage des transitions, etc.).

### Ce que ce module n’est pas (encore)

- Un moteur de workflow.
- Une implémentation serveur.
- Le catalogue unique des types métier de votre application consommatrice (celle-ci peut garder ses propres modèles DB/UI).

### Contrainte Go importante : `internal`

Souvent, une application place ses structs sous `…/internal/…`.
En Go, un package `internal/...` **ne peut pas être importé** par un autre module.

Conséquence : ce SDK expose ses modèles dans `github.com/soft-optimum-services/workflow-service-sdk/model` en **dupliquant** la structure et les tags JSON attendus par l’API `workflow-service`.
Si vous maintenez aussi des types dans votre app, il faut garder **la parité JSON** (tags, champs) lorsque le contrat API évolue.

### Installation (module)

```bash
go get github.com/soft-optimum-services/workflow-service-sdk
```

Pour un développement local avant publication, utilisez une directive `replace` dans le `go.mod` du consommateur :

```go
replace github.com/soft-optimum-services/workflow-service-sdk => ../chemin/vers/workflow-service-sdk
```

### Configuration

Le client est construit à partir d’un `BaseURL` (origin HTTP de `workflow-service`).
La helper `NewFromEnv()` lit `WORKFLOW_SERVICE_BASE_URL`.

Exemple :

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

### Interfaces : dépendances modulaires

Le module expose des petites interfaces + une interface composite :

- `TemplateReader` : lecture des templates (`ListStates`, `ListTransitions`, `GetTemplateByCode`)
- `ExecutionClient` : exécution (`CanTransition`, `ExecuteTransition`, `BindEntity`, etc.)
- `WorkflowQueries` : helpers dérivés (ex. `ListManualTransitionsFromState`, `ResolveTransitionTargetStateCode`)
- `WorkflowClient` : interface composite (composition des trois ci-dessus)

Recommandation : dépendre des **petites interfaces** lorsque c’est possible, et utiliser `WorkflowClient` seulement si un composant a réellement besoin de tout.

### Endpoints couverts

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

### « Queries » (helpers)

Certaines opérations (UI, orchestration, batch) ont besoin de logique additionnelle :

- `ListManualTransitionsFromState` : charge states+transitions, résout l’ID du `from_state`, filtre les transitions `manual` activées, trie par priorité puis code.
- `ResolveTransitionTargetStateCode` : résout le `State.code` cible à partir d’un `transition_code`.

### Layout du code

- `client.go` : struct `Client`, constructeurs
- `http_transport.go` : transport HTTP GET/POST JSON
- `template_reader.go` : endpoints template
- `execution.go`, `execution_types.go` : endpoints execution
- `queries.go` : helpers dérivés
- `model/` : types JSON (Template/State/Transition/…)

### Versioning / compatibilité

Ce module est un SDK : sa compatibilité dépend du contrat HTTP/JSON de `workflow-service`.
En cas de breaking change côté API (champs renommés, enveloppes modifiées), il faut adapter :

- les enveloppes de réponse dans le client
- les types dans `model/`
