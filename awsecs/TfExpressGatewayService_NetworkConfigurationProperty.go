package awsecs


// Experimental.
type TfExpressGatewayService_NetworkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#security_groups TfExpressGatewayService#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#subnets TfExpressGatewayService#subnets}.
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

