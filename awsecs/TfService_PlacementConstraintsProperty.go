package awsecs


// Experimental.
type TfService_PlacementConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#type TfService#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#expression TfService#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

