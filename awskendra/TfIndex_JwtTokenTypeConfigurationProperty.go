package awskendra


// Experimental.
type TfIndex_JwtTokenTypeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#key_location TfIndex#key_location}.
	// Experimental.
	KeyLocation *string `field:"required" json:"keyLocation" yaml:"keyLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#claim_regex TfIndex#claim_regex}.
	// Experimental.
	ClaimRegex *string `field:"optional" json:"claimRegex" yaml:"claimRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#group_attribute_field TfIndex#group_attribute_field}.
	// Experimental.
	GroupAttributeField *string `field:"optional" json:"groupAttributeField" yaml:"groupAttributeField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#issuer TfIndex#issuer}.
	// Experimental.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#secrets_manager_arn TfIndex#secrets_manager_arn}.
	// Experimental.
	SecretsManagerArn *string `field:"optional" json:"secretsManagerArn" yaml:"secretsManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#url TfIndex#url}.
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#user_name_attribute_field TfIndex#user_name_attribute_field}.
	// Experimental.
	UserNameAttributeField *string `field:"optional" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

