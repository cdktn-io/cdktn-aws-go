package awssecurityhub


// Experimental.
type AwsSecurityhubConfigurationPolicy_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#name AwsSecurityhubConfigurationPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#value_type AwsSecurityhubConfigurationPolicy#value_type}.
	// Experimental.
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// bool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#bool AwsSecurityhubConfigurationPolicy#bool}
	// Experimental.
	Bool *AwsSecurityhubConfigurationPolicy_BoolProperty `field:"optional" json:"bool" yaml:"bool"`
	// double block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#double AwsSecurityhubConfigurationPolicy#double}
	// Experimental.
	Double *AwsSecurityhubConfigurationPolicy_DoubleProperty `field:"optional" json:"double" yaml:"double"`
	// enum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enum AwsSecurityhubConfigurationPolicy#enum}
	// Experimental.
	Enum *AwsSecurityhubConfigurationPolicy_EnumProperty `field:"optional" json:"enum" yaml:"enum"`
	// enum_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enum_list AwsSecurityhubConfigurationPolicy#enum_list}
	// Experimental.
	EnumList *AwsSecurityhubConfigurationPolicy_EnumListProperty `field:"optional" json:"enumList" yaml:"enumList"`
	// int block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#int AwsSecurityhubConfigurationPolicy#int}
	// Experimental.
	Int *AwsSecurityhubConfigurationPolicy_IntProperty `field:"optional" json:"int" yaml:"int"`
	// int_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#int_list AwsSecurityhubConfigurationPolicy#int_list}
	// Experimental.
	IntList *AwsSecurityhubConfigurationPolicy_IntListProperty `field:"optional" json:"intList" yaml:"intList"`
	// string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#string AwsSecurityhubConfigurationPolicy#string}
	// Experimental.
	String *AwsSecurityhubConfigurationPolicy_StringProperty `field:"optional" json:"string" yaml:"string"`
	// string_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#string_list AwsSecurityhubConfigurationPolicy#string_list}
	// Experimental.
	StringList *AwsSecurityhubConfigurationPolicy_StringListProperty `field:"optional" json:"stringList" yaml:"stringList"`
}

