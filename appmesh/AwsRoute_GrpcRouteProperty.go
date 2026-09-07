package appmesh


// Experimental.
type AwsRoute_GrpcRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action AwsRoute#action}
	// Experimental.
	Action *AwsRoute_SpecGrpcRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsRoute#match}
	// Experimental.
	Match *AwsRoute_SpecGrpcRouteMatchProperty `field:"optional" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#retry_policy AwsRoute#retry_policy}
	// Experimental.
	RetryPolicy *AwsRoute_SpecGrpcRouteRetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout AwsRoute#timeout}
	// Experimental.
	Timeout *AwsRoute_SpecGrpcRouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

