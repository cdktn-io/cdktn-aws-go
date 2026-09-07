package appmesh


// Experimental.
type AwsRoute_Http2RouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action AwsRoute#action}
	// Experimental.
	Action *AwsRoute_SpecHttp2RouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsRoute#match}
	// Experimental.
	Match *AwsRoute_SpecHttp2RouteMatchProperty `field:"required" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#retry_policy AwsRoute#retry_policy}
	// Experimental.
	RetryPolicy *AwsRoute_SpecHttp2RouteRetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout AwsRoute#timeout}
	// Experimental.
	Timeout *AwsRoute_SpecHttp2RouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

