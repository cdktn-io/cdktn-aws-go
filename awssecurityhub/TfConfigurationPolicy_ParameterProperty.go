package awssecurityhub


// Experimental.
type TfConfigurationPolicy_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#name TfConfigurationPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#value_type TfConfigurationPolicy#value_type}.
	// Experimental.
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// bool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#bool TfConfigurationPolicy#bool}
	// Experimental.
	Bool *TfConfigurationPolicy_BoolProperty `field:"optional" json:"bool" yaml:"bool"`
	// double block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#double TfConfigurationPolicy#double}
	// Experimental.
	Double *TfConfigurationPolicy_DoubleProperty `field:"optional" json:"double" yaml:"double"`
	// enum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enum TfConfigurationPolicy#enum}
	// Experimental.
	Enum *TfConfigurationPolicy_EnumProperty `field:"optional" json:"enum" yaml:"enum"`
	// enum_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enum_list TfConfigurationPolicy#enum_list}
	// Experimental.
	EnumList *TfConfigurationPolicy_EnumListProperty `field:"optional" json:"enumList" yaml:"enumList"`
	// int block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#int TfConfigurationPolicy#int}
	// Experimental.
	Int *TfConfigurationPolicy_IntProperty `field:"optional" json:"int" yaml:"int"`
	// int_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#int_list TfConfigurationPolicy#int_list}
	// Experimental.
	IntList *TfConfigurationPolicy_IntListProperty `field:"optional" json:"intList" yaml:"intList"`
	// string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#string TfConfigurationPolicy#string}
	// Experimental.
	String *TfConfigurationPolicy_StringProperty `field:"optional" json:"string" yaml:"string"`
	// string_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#string_list TfConfigurationPolicy#string_list}
	// Experimental.
	StringList *TfConfigurationPolicy_StringListProperty `field:"optional" json:"stringList" yaml:"stringList"`
}

