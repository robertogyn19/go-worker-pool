package worker_pool

type GenericJob interface {
	Start(id int)
}
