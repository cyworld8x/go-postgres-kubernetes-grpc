package worker

type IWorker interface {
	Map(options ...ConfigMap) *Worker
}
