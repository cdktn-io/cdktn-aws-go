package appmesh


// Experimental.
type AwsRoute_SpecTcpRouteTimeoutProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#idle AwsRoute#idle}
	// Experimental.
	Idle *AwsRoute_SpecTcpRouteTimeoutIdleProperty `field:"optional" json:"idle" yaml:"idle"`
}

