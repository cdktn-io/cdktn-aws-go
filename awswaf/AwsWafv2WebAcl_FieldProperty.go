package awswaf


// Experimental.
type AwsWafv2WebAcl_FieldProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#field_type AwsWafv2WebAcl#field_type}.
	// Experimental.
	FieldType *string `field:"required" json:"fieldType" yaml:"fieldType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#field_keys AwsWafv2WebAcl#field_keys}.
	// Experimental.
	FieldKeys *[]*string `field:"optional" json:"fieldKeys" yaml:"fieldKeys"`
}

