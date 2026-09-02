package awsappmesh


// Experimental.
type TfGatewayRoute_SpecGrpcRouteActionProperty struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#target TfGatewayRoute#target}
	// Experimental.
	Target *TfGatewayRoute_SpecGrpcRouteActionTargetProperty `field:"required" json:"target" yaml:"target"`
}

