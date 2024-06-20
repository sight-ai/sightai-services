package lib_subgraph

type Config struct {
	Nodes []*Node `json:"nodes"`
}

type Node struct {
	Endpoint  string `json:"endpoint"`
	ChainID   uint   `json:"chainID"`
	ChainName string `json:"chainName"`
}
