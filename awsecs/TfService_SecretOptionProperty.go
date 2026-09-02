package awsecs


// Experimental.
type TfService_SecretOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name TfService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#value_from TfService#value_from}.
	// Experimental.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

