package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_ValidationSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#validation_role AwsSagemakerAlgorithm#validation_role}.
	// Experimental.
	ValidationRole *string `field:"required" json:"validationRole" yaml:"validationRole"`
	// validation_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#validation_profiles AwsSagemakerAlgorithm#validation_profiles}
	// Experimental.
	ValidationProfiles interface{} `field:"optional" json:"validationProfiles" yaml:"validationProfiles"`
}

