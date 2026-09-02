package awsecs


// Experimental.
type TfService_TimeoutProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#idle_timeout_seconds TfService#idle_timeout_seconds}.
	// Experimental.
	IdleTimeoutSeconds *float64 `field:"optional" json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#per_request_timeout_seconds TfService#per_request_timeout_seconds}.
	// Experimental.
	PerRequestTimeoutSeconds *float64 `field:"optional" json:"perRequestTimeoutSeconds" yaml:"perRequestTimeoutSeconds"`
}

