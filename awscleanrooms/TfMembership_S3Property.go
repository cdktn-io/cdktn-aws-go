package awscleanrooms


// Experimental.
type TfMembership_S3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_membership#bucket TfMembership#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_membership#result_format TfMembership#result_format}.
	// Experimental.
	ResultFormat *string `field:"required" json:"resultFormat" yaml:"resultFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_membership#key_prefix TfMembership#key_prefix}.
	// Experimental.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

