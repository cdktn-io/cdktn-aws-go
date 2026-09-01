package awsappmesh


// Experimental.
type AwsAppmeshRoute_GrpcRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action AwsAppmeshRoute#action}
	// Experimental.
	Action *AwsAppmeshRoute_SpecGrpcRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsAppmeshRoute#match}
	// Experimental.
	Match *AwsAppmeshRoute_SpecGrpcRouteMatchProperty `field:"optional" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#retry_policy AwsAppmeshRoute#retry_policy}
	// Experimental.
	RetryPolicy *AwsAppmeshRoute_SpecGrpcRouteRetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout AwsAppmeshRoute#timeout}
	// Experimental.
	Timeout *AwsAppmeshRoute_SpecGrpcRouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

