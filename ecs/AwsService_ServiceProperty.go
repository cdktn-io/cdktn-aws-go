package ecs


// Experimental.
type AwsService_ServiceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#port_name AwsService#port_name}.
	// Experimental.
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// client_alias block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#client_alias AwsService#client_alias}
	// Experimental.
	ClientAlias *AwsService_ClientAliasProperty `field:"optional" json:"clientAlias" yaml:"clientAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#discovery_name AwsService#discovery_name}.
	// Experimental.
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#ingress_port_override AwsService#ingress_port_override}.
	// Experimental.
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#timeout AwsService#timeout}
	// Experimental.
	Timeout *AwsService_TimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#tls AwsService#tls}
	// Experimental.
	Tls *AwsService_TlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

