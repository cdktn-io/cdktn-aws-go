package awsecs


// Experimental.
type TfService_DeploymentCircuitBreakerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#enable TfService#enable}.
	// Experimental.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#rollback TfService#rollback}.
	// Experimental.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

