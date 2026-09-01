package awscloudwatchlogs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudwatchLogDeliveryDestinationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_delivery_destination_policy#delivery_destination_name AwsCloudwatchLogDeliveryDestinationPolicy#delivery_destination_name}.
	// Experimental.
	DeliveryDestinationName *string `field:"required" json:"deliveryDestinationName" yaml:"deliveryDestinationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_delivery_destination_policy#delivery_destination_policy AwsCloudwatchLogDeliveryDestinationPolicy#delivery_destination_policy}.
	// Experimental.
	DeliveryDestinationPolicy *string `field:"required" json:"deliveryDestinationPolicy" yaml:"deliveryDestinationPolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_delivery_destination_policy#region AwsCloudwatchLogDeliveryDestinationPolicy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

