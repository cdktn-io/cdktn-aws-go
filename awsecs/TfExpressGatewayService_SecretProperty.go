package awsecs


// Experimental.
type TfExpressGatewayService_SecretProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#name TfExpressGatewayService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#value_from TfExpressGatewayService#value_from}.
	// Experimental.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

