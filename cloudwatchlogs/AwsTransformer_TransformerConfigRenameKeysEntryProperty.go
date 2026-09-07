package cloudwatchlogs


// Experimental.
type AwsTransformer_TransformerConfigRenameKeysEntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#key AwsTransformer#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#rename_to AwsTransformer#rename_to}.
	// Experimental.
	RenameTo *string `field:"required" json:"renameTo" yaml:"renameTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#overwrite_if_exists AwsTransformer#overwrite_if_exists}.
	// Experimental.
	OverwriteIfExists interface{} `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
}

