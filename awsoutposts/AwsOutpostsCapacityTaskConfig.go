package awsoutposts

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOutpostsCapacityTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#outpost_identifier AwsOutpostsCapacityTask#outpost_identifier}.
	// Experimental.
	OutpostIdentifier *string `field:"required" json:"outpostIdentifier" yaml:"outpostIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#asset_id AwsOutpostsCapacityTask#asset_id}.
	// Experimental.
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// instance_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#instance_pool AwsOutpostsCapacityTask#instance_pool}
	// Experimental.
	InstancePool interface{} `field:"optional" json:"instancePool" yaml:"instancePool"`
	// instances_to_exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#instances_to_exclude AwsOutpostsCapacityTask#instances_to_exclude}
	// Experimental.
	InstancesToExclude interface{} `field:"optional" json:"instancesToExclude" yaml:"instancesToExclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#order_id AwsOutpostsCapacityTask#order_id}.
	// Experimental.
	OrderId *string `field:"optional" json:"orderId" yaml:"orderId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#region AwsOutpostsCapacityTask#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#task_action_on_blocking_instances AwsOutpostsCapacityTask#task_action_on_blocking_instances}.
	// Experimental.
	TaskActionOnBlockingInstances *string `field:"optional" json:"taskActionOnBlockingInstances" yaml:"taskActionOnBlockingInstances"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#timeouts AwsOutpostsCapacityTask#timeouts}
	// Experimental.
	Timeouts *AwsOutpostsCapacityTask_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

