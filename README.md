## Kubernetes Deployment Examples

This repository offers a simplistic [Go app](app/app.go) along with resources useful for deploying the app to [Kubernetes](https://kubernetes.io).

The Go application exposes 2 endpoints:

 * `/health`: Responds with HTTP 200, useful for talking to Kubernetes [Readiness Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/).
 * `/version`: Exposes the version number of the application, which is set as a [constant](app/app.go#L8).