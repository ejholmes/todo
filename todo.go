package todo

// Client wraps services.
type Client struct {
	Todos TodosService
}

// New returns a new Client.
func New() *Client {
	return &Client{
		Todos: NewTodosService(),
	}
}
