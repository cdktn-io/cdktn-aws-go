package awsglue


// Experimental.
type TfClassifier_XmlClassifierProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#classification TfClassifier#classification}.
	// Experimental.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#row_tag TfClassifier#row_tag}.
	// Experimental.
	RowTag *string `field:"required" json:"rowTag" yaml:"rowTag"`
}

