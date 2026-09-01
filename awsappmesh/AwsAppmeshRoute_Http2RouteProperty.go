package awsappmesh


// Experimental.
type AwsAppmeshRoute_Http2RouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action AwsAppmeshRoute#action}
	// Experimental.
	Action *AwsAppmeshRoute_SpecHttp2RouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsAppmeshRoute#match}
	// Experimental.
	Match *AwsAppmeshRoute_SpecHttp2RouteMatchProperty `field:"required" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#retry_policy AwsAppmeshRoute#retry_policy}
	// Experimental.
	RetryPolicy *AwsAppmeshRoute_SpecHttp2RouteRetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout AwsAppmeshRoute#timeout}
	// Experimental.
	Timeout *AwsAppmeshRoute_SpecHttp2RouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

