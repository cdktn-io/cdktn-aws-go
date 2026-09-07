package appmesh


// Experimental.
type AwsRoute_SpecHttp2RouteTimeoutProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#idle AwsRoute#idle}
	// Experimental.
	Idle *AwsRoute_SpecHttp2RouteTimeoutIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#per_request AwsRoute#per_request}
	// Experimental.
	PerRequest *AwsRoute_SpecHttp2RouteTimeoutPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

