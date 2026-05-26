# Exercise 02 — Start here

> Your task: turn the three Pedelec services into a deployed application on Kubernetes.

## What's given to you

- The full Docker Compose stack from exercise 01, so you can still run things locally for sanity checks.
- Container images already published at `ghcr.io/mtze/pedelec-<service>:latest` (the workflow in `.github/workflows/build-and-publish.yml` builds them).

## What you need to produce

A `k8s/` directory in this folder containing:

1. **Three Deployments** — one per service. Each runs one replica of the corresponding container.
2. **Three Services** — one per Deployment. Each exposes port 80 inside the cluster, mapped to the service's container port.
3. **One Ingress** — fans out HTTP paths to the three Services as shown in the request-flow diagram in [`docs/architecture.md`](../../../docs/architecture.md).

Suggested final layout:

```
k8s/
├── deployments/
│   ├── damage-deployment.yml
│   ├── location-deployment.yml
│   └── reservation-deployment.yml
├── services/
│   ├── damage-service.yml
│   ├── location-service.yml
│   └── reservation-service.yml
└── pedelec-ingress.yml
```

## Step-by-step

### 1. Start with the damage Deployment

It's the simplest. The image is `ghcr.io/mtze/pedelec-damage:latest`, the container listens on `8082`, and you want one replica labeled `app: damage`.

<details><summary>Hint — minimal Deployment skeleton</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: damage
  namespace: my-namespace
spec:
  replicas: 1
  selector:
    matchLabels:
      app: damage
  template:
    metadata:
      labels:
        app: damage
    spec:
      containers:
        - name: damage
          image: ghcr.io/mtze/pedelec-damage:latest
          ports:
            - containerPort: 8082
              name: damage-api
```
</details>

Apply and verify:

```bash
kubectl apply -f k8s/deployments/damage-deployment.yml
kubectl get pods
kubectl logs deploy/damage
```

Expected: one pod in `Running`, logs say `Starting damage service`.

### 2. Add a Service for damage

A Service is just a label selector + a port mapping. Target the container port by its **name** (`damage-api`), not its number — it makes refactors safer.

<details><summary>Hint — minimal Service skeleton</summary>

```yaml
apiVersion: v1
kind: Service
metadata:
  name: pedelec-damage-service
  namespace: my-namespace
spec:
  selector:
    app: damage
  ports:
    - protocol: TCP
      port: 80
      targetPort: damage-api
```
</details>

Verify the Service found the pod:

```bash
kubectl get endpoints pedelec-damage-service
```

If you see an IP listed, the selector works. If `<none>`, the labels don't match.

### 3. Repeat for location and reservation

Same pattern, different image / port. The ports are:

| Service | Container port | Port name |
|---|---|---|
| damage | 8082 | damage-api |
| location | 8081 | location-api |
| reservation | 8080 | reservation-api |

### 4. Write the Ingress

The Ingress routes a single hostname into the right Service based on path. The path → service mapping is:

- `/pedelec` → `pedelec-reservation-service`
- `/reservation` → `pedelec-reservation-service`
- `/locations` → `pedelec-location-service`
- `/damage` → `pedelec-damage-service`

You'll need to replace the `host:` value with your own (see [`docs/cluster-access.md`](../../../docs/cluster-access.md)).

### 5. Apply everything

```bash
kubectl apply -f k8s/deployments/
kubectl apply -f k8s/services/
kubectl apply -f k8s/pedelec-ingress.yml
```

Or in one shot if you like:

```bash
kubectl apply -R -f k8s/
```

## How to verify

Once the ingress has an address, curl your hostname:

```bash
curl https://<your-host>/pedelec | jq
curl -X POST https://<your-host>/locations \
  -H 'Content-Type: application/json' \
  -d '{"id":"1","lat":48.149,"lng":11.568}'
```

If you'd rather skip the ingress while debugging, port-forward:

```bash
kubectl port-forward service/pedelec-reservation-service 8080:80
curl http://localhost:8080/pedelec
```

## Common pitfalls

> [!WARNING] **Endpoints empty** — your Service selector and your Pod labels don't match. Run `kubectl get pod --show-labels` and `kubectl describe service <name>` and compare.

> [!WARNING] **ImagePullBackOff** — usually a typo in the image name or tag. The upstream images are `ghcr.io/mtze/pedelec-<service>:latest` and are *public* — no `imagePullSecrets` needed.

> [!WARNING] **Ingress address never appears** — your cluster might not have an ingress controller installed. On a local cluster, follow [`docs/prerequisites.md`](../../../docs/prerequisites.md).

> [!TIP] **`kubectl describe` is your friend** — it shows the most recent events for a resource. Most "why isn't this working" questions get answered by it.

## Done?

Open [`../solution/README.md`](../solution/README.md) for the reference manifests and a walkthrough of the design decisions.
