package docker

// 镜像结构
type Image struct {
	Created     uint64
	Id          string
	ParentId    string
	RepoTags    []string
	Size        uint64
	VirtualSize uint64
}

// 容器结构
type Container struct {
	Id              string                 `json:"Id"`
	Names           []string               `json:"Names"`
	Image           string                 `json:"Image"`
	ImageID         string                 `json:"ImageID"`
	Command         string                 `json:"Command"`
	Created         uint64                 `json:"Created"`
	State           string                 `json:"State"`
	Status          string                 `json:"Status"`
	Ports           []Port                 `json:"Ports"`
	Labels          map[string]string      `json:"Labels"`
	HostConfig      map[string]string      `json:"HostConfig"`
	NetworkSettings map[string]interface{} `json:"NetworkSettings"`
	Mounts          []Mount                `json:"Mounts"`
}

// docker 端口映射
type Port struct {
	IP          string `json:"IP"`
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}

// docker 挂载
type Mount struct {
	Type        string `json:"Type"`
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
	Propatation string `json:"Propagation"`
}
