package awselasticbeanstalk


// Experimental.
type TfEnvironment_SettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_environment#name TfEnvironment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_environment#namespace TfEnvironment#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_environment#value TfEnvironment#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_environment#resource TfEnvironment#resource}.
	// Experimental.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

