package awsecs


// Experimental.
type TfService_ServiceConnectConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#enabled TfService#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// access_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#access_log_configuration TfService#access_log_configuration}
	// Experimental.
	AccessLogConfiguration *TfService_AccessLogConfigurationProperty `field:"optional" json:"accessLogConfiguration" yaml:"accessLogConfiguration"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#log_configuration TfService#log_configuration}
	// Experimental.
	LogConfiguration *TfService_LogConfigurationProperty `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#namespace TfService#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#service TfService#service}
	// Experimental.
	Service interface{} `field:"optional" json:"service" yaml:"service"`
}

