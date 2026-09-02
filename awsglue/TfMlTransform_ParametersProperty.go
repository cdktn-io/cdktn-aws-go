package awsglue


// Experimental.
type TfMlTransform_ParametersProperty struct {
	// find_matches_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#find_matches_parameters TfMlTransform#find_matches_parameters}
	// Experimental.
	FindMatchesParameters *TfMlTransform_FindMatchesParametersProperty `field:"required" json:"findMatchesParameters" yaml:"findMatchesParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#transform_type TfMlTransform#transform_type}.
	// Experimental.
	TransformType *string `field:"required" json:"transformType" yaml:"transformType"`
}

