package awsmacie


// Experimental.
type AwsMacie2ClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#comparator AwsMacie2ClassificationJob#comparator}.
	// Experimental.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#key AwsMacie2ClassificationJob#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#values AwsMacie2ClassificationJob#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

