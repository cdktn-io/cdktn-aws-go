package awsappmesh


// Experimental.
type AwsAppmeshRoute_SpecTcpRouteTimeoutProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#idle AwsAppmeshRoute#idle}
	// Experimental.
	Idle *AwsAppmeshRoute_SpecTcpRouteTimeoutIdleProperty `field:"optional" json:"idle" yaml:"idle"`
}

