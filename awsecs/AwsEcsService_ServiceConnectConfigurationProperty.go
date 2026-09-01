package awsecs


// Experimental.
type AwsEcsService_ServiceConnectConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#enabled AwsEcsService#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// access_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#access_log_configuration AwsEcsService#access_log_configuration}
	// Experimental.
	AccessLogConfiguration *AwsEcsService_AccessLogConfigurationProperty `field:"optional" json:"accessLogConfiguration" yaml:"accessLogConfiguration"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#log_configuration AwsEcsService#log_configuration}
	// Experimental.
	LogConfiguration *AwsEcsService_LogConfigurationProperty `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#namespace AwsEcsService#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#service AwsEcsService#service}
	// Experimental.
	Service interface{} `field:"optional" json:"service" yaml:"service"`
}

