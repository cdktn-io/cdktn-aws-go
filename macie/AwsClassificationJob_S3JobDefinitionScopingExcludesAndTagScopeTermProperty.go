package macie


// Experimental.
type AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#comparator AwsClassificationJob#comparator}.
	// Experimental.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#key AwsClassificationJob#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_values AwsClassificationJob#tag_values}
	// Experimental.
	TagValues interface{} `field:"optional" json:"tagValues" yaml:"tagValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#target AwsClassificationJob#target}.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

