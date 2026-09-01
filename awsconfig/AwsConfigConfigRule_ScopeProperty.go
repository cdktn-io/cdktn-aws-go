package awsconfig


// Experimental.
type AwsConfigConfigRule_ScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#compliance_resource_id AwsConfigConfigRule#compliance_resource_id}.
	// Experimental.
	ComplianceResourceId *string `field:"optional" json:"complianceResourceId" yaml:"complianceResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#compliance_resource_types AwsConfigConfigRule#compliance_resource_types}.
	// Experimental.
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#tag_key AwsConfigConfigRule#tag_key}.
	// Experimental.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#tag_value AwsConfigConfigRule#tag_value}.
	// Experimental.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

