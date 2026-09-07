package appmesh


// Experimental.
type AwsRoute_SpecHttpRouteTimeoutProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#idle AwsRoute#idle}
	// Experimental.
	Idle *AwsRoute_SpecHttpRouteTimeoutIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#per_request AwsRoute#per_request}
	// Experimental.
	PerRequest *AwsRoute_SpecHttpRouteTimeoutPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

