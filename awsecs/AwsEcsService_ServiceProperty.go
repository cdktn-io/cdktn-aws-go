package awsecs


// Experimental.
type AwsEcsService_ServiceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#port_name AwsEcsService#port_name}.
	// Experimental.
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// client_alias block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#client_alias AwsEcsService#client_alias}
	// Experimental.
	ClientAlias *AwsEcsService_ClientAliasProperty `field:"optional" json:"clientAlias" yaml:"clientAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#discovery_name AwsEcsService#discovery_name}.
	// Experimental.
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#ingress_port_override AwsEcsService#ingress_port_override}.
	// Experimental.
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#timeout AwsEcsService#timeout}
	// Experimental.
	Timeout *AwsEcsService_TimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#tls AwsEcsService#tls}
	// Experimental.
	Tls *AwsEcsService_TlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

