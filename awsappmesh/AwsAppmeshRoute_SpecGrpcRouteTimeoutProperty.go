package awsappmesh


// Experimental.
type AwsAppmeshRoute_SpecGrpcRouteTimeoutProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#idle AwsAppmeshRoute#idle}
	// Experimental.
	Idle *AwsAppmeshRoute_SpecGrpcRouteTimeoutIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#per_request AwsAppmeshRoute#per_request}
	// Experimental.
	PerRequest *AwsAppmeshRoute_SpecGrpcRouteTimeoutPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

