package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecGrpcRouteActionProperty struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#target AwsAppmeshGatewayRoute#target}
	// Experimental.
	Target *AwsAppmeshGatewayRoute_SpecGrpcRouteActionTargetProperty `field:"required" json:"target" yaml:"target"`
}

