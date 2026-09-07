package backup


// Experimental.
type AwsSelection_StringNotEqualsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_selection#key AwsSelection#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_selection#value AwsSelection#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

