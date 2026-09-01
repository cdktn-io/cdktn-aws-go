package awsemr


// Experimental.
type AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_block_public_access_configuration#max_range AwsEmrBlockPublicAccessConfiguration#max_range}.
	// Experimental.
	MaxRange *float64 `field:"required" json:"maxRange" yaml:"maxRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_block_public_access_configuration#min_range AwsEmrBlockPublicAccessConfiguration#min_range}.
	// Experimental.
	MinRange *float64 `field:"required" json:"minRange" yaml:"minRange"`
}

