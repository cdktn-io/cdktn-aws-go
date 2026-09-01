package awsglue


// Experimental.
type AwsGlueMlTransform_ParametersProperty struct {
	// find_matches_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#find_matches_parameters AwsGlueMlTransform#find_matches_parameters}
	// Experimental.
	FindMatchesParameters *AwsGlueMlTransform_FindMatchesParametersProperty `field:"required" json:"findMatchesParameters" yaml:"findMatchesParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#transform_type AwsGlueMlTransform#transform_type}.
	// Experimental.
	TransformType *string `field:"required" json:"transformType" yaml:"transformType"`
}

