package awscloudtrail


// Experimental.
type TfCloudtrail_FieldSelectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#field TfCloudtrail#field}.
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#ends_with TfCloudtrail#ends_with}.
	// Experimental.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#equals TfCloudtrail#equals}.
	// Experimental.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#not_ends_with TfCloudtrail#not_ends_with}.
	// Experimental.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#not_equals TfCloudtrail#not_equals}.
	// Experimental.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#not_starts_with TfCloudtrail#not_starts_with}.
	// Experimental.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#starts_with TfCloudtrail#starts_with}.
	// Experimental.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

