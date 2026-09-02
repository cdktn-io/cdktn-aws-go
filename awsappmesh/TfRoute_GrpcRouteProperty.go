package awsappmesh


// Experimental.
type TfRoute_GrpcRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action TfRoute#action}
	// Experimental.
	Action *TfRoute_SpecGrpcRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match TfRoute#match}
	// Experimental.
	Match *TfRoute_SpecGrpcRouteMatchProperty `field:"optional" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#retry_policy TfRoute#retry_policy}
	// Experimental.
	RetryPolicy *TfRoute_SpecGrpcRouteRetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout TfRoute#timeout}
	// Experimental.
	Timeout *TfRoute_SpecGrpcRouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

