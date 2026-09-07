package glue


// Experimental.
type AwsClassifier_XmlClassifierProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#classification AwsClassifier#classification}.
	// Experimental.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#row_tag AwsClassifier#row_tag}.
	// Experimental.
	RowTag *string `field:"required" json:"rowTag" yaml:"rowTag"`
}

