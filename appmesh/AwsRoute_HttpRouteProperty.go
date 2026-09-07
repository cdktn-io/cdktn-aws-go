package appmesh


// Experimental.
type AwsRoute_HttpRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action AwsRoute#action}
	// Experimental.
	Action *AwsRoute_SpecHttpRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsRoute#match}
	// Experimental.
	Match *AwsRoute_SpecHttpRouteMatchProperty `field:"required" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#retry_policy AwsRoute#retry_policy}
	// Experimental.
	RetryPolicy *AwsRoute_SpecHttpRouteRetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout AwsRoute#timeout}
	// Experimental.
	Timeout *AwsRoute_SpecHttpRouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

