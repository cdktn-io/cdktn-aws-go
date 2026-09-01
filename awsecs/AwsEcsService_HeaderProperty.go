package awsecs


// Experimental.
type AwsEcsService_HeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name AwsEcsService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#value AwsEcsService#value}
	// Experimental.
	Value *AwsEcsService_ValueProperty `field:"required" json:"value" yaml:"value"`
}

