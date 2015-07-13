package api

// Structures common between different uris of the Qubole REST API

type PagingInfo struct {
	NextPage     uint   `json:"next_page"`
	PerPage      string `json:"per_page"`
	PreviousPage uint   `json:"previous_page"`
}
