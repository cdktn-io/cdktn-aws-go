package cognitoidp


// Experimental.
type AwsUserPool_SchemaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#attribute_data_type AwsUserPool#attribute_data_type}.
	// Experimental.
	AttributeDataType *string `field:"required" json:"attributeDataType" yaml:"attributeDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#name AwsUserPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#developer_only_attribute AwsUserPool#developer_only_attribute}.
	// Experimental.
	DeveloperOnlyAttribute interface{} `field:"optional" json:"developerOnlyAttribute" yaml:"developerOnlyAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#mutable AwsUserPool#mutable}.
	// Experimental.
	Mutable interface{} `field:"optional" json:"mutable" yaml:"mutable"`
	// number_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#number_attribute_constraints AwsUserPool#number_attribute_constraints}
	// Experimental.
	NumberAttributeConstraints *AwsUserPool_NumberAttributeConstraintsProperty `field:"optional" json:"numberAttributeConstraints" yaml:"numberAttributeConstraints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#required AwsUserPool#required}.
	// Experimental.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// string_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#string_attribute_constraints AwsUserPool#string_attribute_constraints}
	// Experimental.
	StringAttributeConstraints *AwsUserPool_StringAttributeConstraintsProperty `field:"optional" json:"stringAttributeConstraints" yaml:"stringAttributeConstraints"`
}

