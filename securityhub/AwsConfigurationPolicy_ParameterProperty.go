package securityhub


// Experimental.
type AwsConfigurationPolicy_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#name AwsConfigurationPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#value_type AwsConfigurationPolicy#value_type}.
	// Experimental.
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// bool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#bool AwsConfigurationPolicy#bool}
	// Experimental.
	Bool *AwsConfigurationPolicy_BoolProperty `field:"optional" json:"bool" yaml:"bool"`
	// double block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#double AwsConfigurationPolicy#double}
	// Experimental.
	Double *AwsConfigurationPolicy_DoubleProperty `field:"optional" json:"double" yaml:"double"`
	// enum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enum AwsConfigurationPolicy#enum}
	// Experimental.
	Enum *AwsConfigurationPolicy_EnumProperty `field:"optional" json:"enum" yaml:"enum"`
	// enum_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enum_list AwsConfigurationPolicy#enum_list}
	// Experimental.
	EnumList *AwsConfigurationPolicy_EnumListProperty `field:"optional" json:"enumList" yaml:"enumList"`
	// int block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#int AwsConfigurationPolicy#int}
	// Experimental.
	Int *AwsConfigurationPolicy_IntProperty `field:"optional" json:"int" yaml:"int"`
	// int_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#int_list AwsConfigurationPolicy#int_list}
	// Experimental.
	IntList *AwsConfigurationPolicy_IntListProperty `field:"optional" json:"intList" yaml:"intList"`
	// string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#string AwsConfigurationPolicy#string}
	// Experimental.
	String *AwsConfigurationPolicy_StringProperty `field:"optional" json:"string" yaml:"string"`
	// string_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#string_list AwsConfigurationPolicy#string_list}
	// Experimental.
	StringList *AwsConfigurationPolicy_StringListProperty `field:"optional" json:"stringList" yaml:"stringList"`
}

