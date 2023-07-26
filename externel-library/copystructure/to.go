package main

type Order struct {
	ProductListSeq int      `validate:"gt=0"`
	TenantID       string   `validate:"required"`
	Info           OrderHis `validate:"required,dive,required"`
}

type OrderHis struct {
	CloudID      string         `validate:"required" json:"cloud_id,omitempty"`
	ProductID    string         `validate:"required" json:"product_id,omitempty"`
	ServiceID    string         `validate:"required" json:"service_id,omitempty"`
	GroupName    string         `json:"group_name,omitempty"` // Custom 내부에 쓰기 위한 용도
	License      LicenseO       `validate:"required,dive,required" json:"license,omitempty"`
	AddService   []AddServiceO  `validate:"required,dive,required" json:"add_service,omitempty"`
	Autoscale    Autoscale      `validate:"required,dive,required" json:"autoscale,omitempty"`
	Config       []Config       `validate:"required,dive,required" json:"config,omitempty"`
	Server       []Server       `validate:"required,dive,required" json:"server,omitempty"`
	NicSubnet    NicSubnet      `validate:"required,dive,required" json:"nic_subnet,omitempty"`
	LoadBalancer []LoadBalancer `json:"lb,omitempty"`
	Acg          []Acg          `validate:"required,dive,required" json:"acg,omitempty"`
}

type LicenseO struct {
	ID          string `validate:"required" json:"id,omitempty"`
	Plan        string `json:"plan,omitempty"`
	Price       int    `json:"price,omitempty"`
	ProductID   string `json:"product_id,omitempty"`
	ServiceID   string `json:"service_id,omitempty"`
	Grade       string `json:"grade,omitempty"`
	Description string `json:"description,omitempty"`
	Group       struct {
		Min int `validate:"gt=0" json:"min,omitempty"`
		Max int `validate:"gt=0" json:"max,omitempty"`
	} `json:"group,omitempty"`
}

type AddServiceO struct {
	ID          string `validate:"required" json:"id,omitempty"`
	Plan        string `json:"plan,omitempty"`
	Price       int    `json:"price,omitempty"`
	Description string `json:"description,omitempty"`
}

// TODO: AutoScale 정책 필수요소 검토 필요
type Autoscale struct {
	Group []struct {
		PolicyID string `json:"policy_id,omitempty"`
		ActValue int    `json:"act_value,omitempty"`
	} `json:"group,omitempty"`
	System []struct {
		SystemType string `json:"system_type,omitempty"`
		System     []struct {
			PolicyID string `json:"policy_id,omitempty"`
			ActValue int    `json:"act_value,omitempty"`
		} `json:"system,omitempty"`
	} `json:"system,omitempty"`
}

type Config struct {
	SystemType string       `validate:"required" json:"system_type,omitempty"`
	System     ConfigSystem `validate:"required,dive,required" json:"system,omitempty"`
}
type ConfigSystem struct {
	CloudVpcName            string `json:"cloud_vpc_name,omitempty"`
	CloudVpcCode            string `json:"cloud_vpc_code,omitempty"`
	CloudVpcNetworkRange    string `validate:"required" json:"cloud_vpc_network_range,omitempty"`
	CloudSubnetName         string `json:"cloud_subnet_name,omitempty"`
	CloudSubnetCode         string `json:"cloud_subnet_code,omitempty"`
	CloudSubnetNetworkRange string `validate:"required" json:"cloud_subnet_network_range,omitempty"`
	CloudSubnetNetworkType  string `validate:"required" json:"cloud_subnet_network_type,omitempty"`
	CloudSubnetUsageType    string `validate:"required" json:"cloud_subnet_usage_type,omitempty"`
}

type NicSubnet struct {
	Subnet []Subnet `validate:"required,dive,required" json:"subnet,omitempty"`
	Group  []Group  `validate:"required,dive,required" json:"group,omitempty"`
}

type Subnet struct {
	Id        string `validate:"required" json:"id,omitempty"`
	Type      string `validate:"required" json:"network_type,omitempty"`
	Range     string `validate:"required" json:"cidr_block,omitempty"`
	UsageType string `json:"usage_type,omitempty"`
}
type Group struct {
	Index  int           `json:"index"`
	System []GroupSystem `validate:"required,dive,required" json:"system,omitempty"`
}
type GroupSystem struct {
	Index int        `json:"index"`
	Type  string     `validate:"required" json:"type,omitempty"`
	Nic   []GroupNic `validate:"required,dive,required" json:"nic,omitempty"`
}
type GroupNic struct {
	Id       string `validate:"required" json:"id,omitempty"`
	SubnetId string `validate:"required" json:"subnet_id,omitempty"`
}

type Server struct {
	SystemType  string `validate:"required" json:"system_type,omitempty"`
	PublicIpUse bool   `json:"public_ip_use,omitempty"`
	System      struct {
		CloudOsID      string `validate:"required" json:"cloud_os_id,omitempty"`
		CloudProductID string `validate:"required" json:"cloud_product_id,omitempty"`
		Min            int    `json:"min,omitempty"`
		Max            int    `json:"max,omitempty"`
	} `validate:"required,dive,required" json:"system,omitempty"`
	Storage []Storage `json:"storage,omitempty"`
}

type Storage struct {
	CloudStorageID string `validate:"required" json:"cloud_storage_id,omitempty"`
	Volume         int    `validate:"required" json:"volume,omitempty"`
	CreateNum      int    `json:"create_num,omitempty"`
}

type LoadBalancer struct {
	Id             string             `validate:"required" json:"id,omitempty"`
	Type           string             `validate:"required" json:"type,omitempty"`
	ThroughputType string             `validate:"required" json:"throughput_type,omitempty"`
	TargetSystem   string             `validate:"required" json:"target_system,omitempty"`
	NetworkType    string             `validate:"required" json:"network_type,omitempty"`
	Name           string             `json:"name,omitempty"`
	Description    string             `json:"description,omitempty"`
	Subnet         []Subnet           `validate:"required,dive,required" json:"subnet,omitempty"`
	Rule           []LoadBalancerRule `validate:"required,dive,required" json:"rule,omitempty"`
}
type LoadBalancerRule struct {
	Id           string `validate:"required" json:"id,omitempty"`
	Protocol     string `validate:"required" json:"protocol,omitempty"`
	ListenerPort int    `validate:"required" json:"listener_port,omitempty"`
	TargetPort   int    `validate:"required" json:"target_port,omitempty"`
	Description  string `json:"description,omitempty"`
}

type Acg struct {
	AcgSeq      int64     `json:"acg_seq,omitempty"` // Custom 내부에 쓰기 위한 용도
	AcgID       string    `validate:"required" json:"acg_id,omitempty"`
	ServiceType string    `json:"service_type,omitempty"`
	Name        string    `json:"name,omitempty"`
	ServiceID   string    `validate:"required" json:"service_id,omitempty"`
	System      string    `validate:"required" json:"system,omitempty"`
	NicID       string    `validate:"required" json:"nic_id,omitempty"`
	AclRule     []AclRule `json:"acl_rule,omitempty"`
	//} `validate:"required,dive,required" json:"acl_rule,omitempty"`
}

type AclRule struct {
	Direction   string `validate:"required" json:"direction,omitempty"`
	Protocol    string `validate:"required" json:"protocol,omitempty"`
	SourceType  string `json:"source_type,omitempty"`
	Source      string `validate:"required" json:"source,omitempty"`
	Ports       string `validate:"required" json:"ports,omitempty"`
	Description string `json:"description,omitempty"`
	AddInfo     struct {
	} `json:"add_info,omitempty"`
}
