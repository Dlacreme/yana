package gen

type Select2 struct {
	Pagination bool          `json:"pagination"`
	Results    []Select2Item `json:"results"`
}

type Select2Item struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}
