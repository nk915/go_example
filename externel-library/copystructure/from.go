package main

type GetOrderHisResponse struct {
	CloudId    *string             `json:"cloud_id,omitempty"`
	ProductId  *string             `json:"product_id,omitempty"`
	ServiceId  *string             `json:"service_id,omitempty"`
	License    *License            `json:"license,omitempty"`
	AddService []AddService        `json:"add_service,omitempty"`
	Autoscale  *OrderInfoAutoscale `json:"autoscale,omitempty"`
	Config     []Config            `json:"config,omitempty"`
	//	Server     []Server            `json:"server,omitempty"`
	//	NicSubnet  *NicSubnet          `json:"nic_subnet,omitempty"`
	//	Lb         []Lb                `json:"lb,omitempty"`
	//	Acg        []Acg               `json:"acg,omitempty"`
}

type LicenseGroup struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

type AddService struct {
	Id          *string `json:"id,omitempty" validate:"required"`
	PlanType    *string `json:"plan_type,omitempty" validate:"required"`
	Price       *int32  `json:"price,omitempty" validate:"gt=0"`
	UseType     *string `json:"use_type,omitempty"`
	Description *string `json:"description,omitempty" validate:"required"`
}

type OrderInfoAutoscale struct {
	Group  []OrderInfoAutoscaleGroupInner  `json:"group,omitempty"`
	System []OrderInfoAutoscaleSystemInner `json:"system,omitempty"`
}

type OrderInfoAutoscaleGroupInner struct {
	PolicyId string `json:"policy_id" validate:"required"`
	ActValue *int32 `json:"act_value,omitempty" validate:"gt=0"`
}

type OrderInfoAutoscaleSystemInner struct {
	SystemType string                                     `json:"system_type" validate:"required"`
	System     []OrderInfoAutoscaleSystemInnerSystemInner `json:"system"`
}

type OrderInfoAutoscaleSystemInnerSystemInner struct {
	PolicyId string `json:"policy_id" validate:"required"`
	ActValue int32  `json:"act_value" validate:"gt=0"`
}

type Config struct {
	SystemId string       `json:"system_id" validate:"required"`
	System   ConfigSystem `json:"system" validate:"required"`
}

type ConfigSystem struct {
	// 클라우드 VPC명
	CloudVpcName string `json:"cloud_vpc_name" validate:"required"`
	// 서버위치를 나타내는 VPC 고유코드
	CloudVpcCode string `json:"cloud_vpc_code" validate:"required"`
	// VPC의 대역정보
	CloudVpcNetworkRange string `json:"cloud_vpc_network_range" validate:"required"`
	// Subnet명
	CloudSubnetName string `json:"cloud_subnet_name" validate:"required"`
	// 서버위치(subnet)를 나타내는 subnet 고유코드
	CloudSubnetCode string `json:"cloud_subnet_code" validate:"required"`
	// Subnet의 대역정보
	CloudSubnetNetworkRange string `json:"cloud_subnet_network_range" validate:"required"`
	// Subnet 용도
	CloudSubnetUsageType *string `json:"cloud_subnet_usage_type,omitempty" default:"GENERAL"`
	// Subnet의 네트워크 유형
	CloudSubnetNetworkType string `json:"cloud_subnet_network_type" validate:"required"`
}
