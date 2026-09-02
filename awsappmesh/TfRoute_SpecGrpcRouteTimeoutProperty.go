package awsappmesh


// Experimental.
type TfRoute_SpecGrpcRouteTimeoutProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#idle TfRoute#idle}
	// Experimental.
	Idle *TfRoute_SpecGrpcRouteTimeoutIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#per_request TfRoute#per_request}
	// Experimental.
	PerRequest *TfRoute_SpecGrpcRouteTimeoutPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

