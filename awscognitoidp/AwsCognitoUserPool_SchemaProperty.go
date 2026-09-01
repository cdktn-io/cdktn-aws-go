package awscognitoidp


// Experimental.
type AwsCognitoUserPool_SchemaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#attribute_data_type AwsCognitoUserPool#attribute_data_type}.
	// Experimental.
	AttributeDataType *string `field:"required" json:"attributeDataType" yaml:"attributeDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#name AwsCognitoUserPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#developer_only_attribute AwsCognitoUserPool#developer_only_attribute}.
	// Experimental.
	DeveloperOnlyAttribute interface{} `field:"optional" json:"developerOnlyAttribute" yaml:"developerOnlyAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#mutable AwsCognitoUserPool#mutable}.
	// Experimental.
	Mutable interface{} `field:"optional" json:"mutable" yaml:"mutable"`
	// number_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#number_attribute_constraints AwsCognitoUserPool#number_attribute_constraints}
	// Experimental.
	NumberAttributeConstraints *AwsCognitoUserPool_NumberAttributeConstraintsProperty `field:"optional" json:"numberAttributeConstraints" yaml:"numberAttributeConstraints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#required AwsCognitoUserPool#required}.
	// Experimental.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// string_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#string_attribute_constraints AwsCognitoUserPool#string_attribute_constraints}
	// Experimental.
	StringAttributeConstraints *AwsCognitoUserPool_StringAttributeConstraintsProperty `field:"optional" json:"stringAttributeConstraints" yaml:"stringAttributeConstraints"`
}

