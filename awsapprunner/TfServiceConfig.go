package awsapprunner

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#service_name TfService#service_name}.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#source_configuration TfService#source_configuration}
	// Experimental.
	SourceConfiguration *TfService_SourceConfigurationProperty `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#auto_scaling_configuration_arn TfService#auto_scaling_configuration_arn}.
	// Experimental.
	AutoScalingConfigurationArn *string `field:"optional" json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#encryption_configuration TfService#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *TfService_EncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// health_check_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#health_check_configuration TfService#health_check_configuration}
	// Experimental.
	HealthCheckConfiguration *TfService_HealthCheckConfigurationProperty `field:"optional" json:"healthCheckConfiguration" yaml:"healthCheckConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#id TfService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#instance_configuration TfService#instance_configuration}
	// Experimental.
	InstanceConfiguration *TfService_InstanceConfigurationProperty `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#network_configuration TfService#network_configuration}
	// Experimental.
	NetworkConfiguration *TfService_NetworkConfigurationProperty `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// observability_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#observability_configuration TfService#observability_configuration}
	// Experimental.
	ObservabilityConfiguration *TfService_ObservabilityConfigurationProperty `field:"optional" json:"observabilityConfiguration" yaml:"observabilityConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#region TfService#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#tags TfService#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#tags_all TfService#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

