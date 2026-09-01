package awsappmesh


// Experimental.
type AwsAppmeshRoute_SpecHttpRouteRetryPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#max_retries AwsAppmeshRoute#max_retries}.
	// Experimental.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// per_retry_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#per_retry_timeout AwsAppmeshRoute#per_retry_timeout}
	// Experimental.
	PerRetryTimeout *AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutProperty `field:"required" json:"perRetryTimeout" yaml:"perRetryTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#http_retry_events AwsAppmeshRoute#http_retry_events}.
	// Experimental.
	HttpRetryEvents *[]*string `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#tcp_retry_events AwsAppmeshRoute#tcp_retry_events}.
	// Experimental.
	TcpRetryEvents *[]*string `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}

