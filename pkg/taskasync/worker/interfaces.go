package worker

type IWorker interface {
	Run() error
	Map(options ...ConfigMap) *Worker
}
