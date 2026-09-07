package config


// Experimental.
type AwsConfigRule_SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#owner AwsConfigRule#owner}.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// custom_policy_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#custom_policy_details AwsConfigRule#custom_policy_details}
	// Experimental.
	CustomPolicyDetails *AwsConfigRule_CustomPolicyDetailsProperty `field:"optional" json:"customPolicyDetails" yaml:"customPolicyDetails"`
	// source_detail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#source_detail AwsConfigRule#source_detail}
	// Experimental.
	SourceDetail interface{} `field:"optional" json:"sourceDetail" yaml:"sourceDetail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#source_identifier AwsConfigRule#source_identifier}.
	// Experimental.
	SourceIdentifier *string `field:"optional" json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

