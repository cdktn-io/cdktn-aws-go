package ecs


// Experimental.
type AwsService_SecretOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name AwsService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#value_from AwsService#value_from}.
	// Experimental.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

