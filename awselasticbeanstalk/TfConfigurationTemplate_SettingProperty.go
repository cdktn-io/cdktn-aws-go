package awselasticbeanstalk


// Experimental.
type TfConfigurationTemplate_SettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_configuration_template#name TfConfigurationTemplate#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_configuration_template#namespace TfConfigurationTemplate#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_configuration_template#value TfConfigurationTemplate#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_configuration_template#resource TfConfigurationTemplate#resource}.
	// Experimental.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

