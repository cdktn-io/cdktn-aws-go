package ecs


// Experimental.
type AwsService_HeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name AwsService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#value AwsService#value}
	// Experimental.
	Value *AwsService_ValueProperty `field:"required" json:"value" yaml:"value"`
}

