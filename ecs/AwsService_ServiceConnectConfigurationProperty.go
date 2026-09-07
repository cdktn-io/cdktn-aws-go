package ecs


// Experimental.
type AwsService_ServiceConnectConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#enabled AwsService#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// access_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#access_log_configuration AwsService#access_log_configuration}
	// Experimental.
	AccessLogConfiguration *AwsService_AccessLogConfigurationProperty `field:"optional" json:"accessLogConfiguration" yaml:"accessLogConfiguration"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#log_configuration AwsService#log_configuration}
	// Experimental.
	LogConfiguration *AwsService_LogConfigurationProperty `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#namespace AwsService#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#service AwsService#service}
	// Experimental.
	Service interface{} `field:"optional" json:"service" yaml:"service"`
}

