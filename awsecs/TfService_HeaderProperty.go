package awsecs


// Experimental.
type TfService_HeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name TfService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#value TfService#value}
	// Experimental.
	Value *TfService_ValueProperty `field:"required" json:"value" yaml:"value"`
}

