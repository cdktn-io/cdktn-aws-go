package awsfinspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFinspaceKxClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#az_mode AwsFinspaceKxCluster#az_mode}.
	// Experimental.
	AzMode *string `field:"required" json:"azMode" yaml:"azMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#environment_id AwsFinspaceKxCluster#environment_id}.
	// Experimental.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#name AwsFinspaceKxCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#release_label AwsFinspaceKxCluster#release_label}.
	// Experimental.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#type AwsFinspaceKxCluster#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#vpc_configuration AwsFinspaceKxCluster#vpc_configuration}
	// Experimental.
	VpcConfiguration *AwsFinspaceKxCluster_VpcConfigurationProperty `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// auto_scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#auto_scaling_configuration AwsFinspaceKxCluster#auto_scaling_configuration}
	// Experimental.
	AutoScalingConfiguration *AwsFinspaceKxCluster_AutoScalingConfigurationProperty `field:"optional" json:"autoScalingConfiguration" yaml:"autoScalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#availability_zone_id AwsFinspaceKxCluster#availability_zone_id}.
	// Experimental.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// cache_storage_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#cache_storage_configurations AwsFinspaceKxCluster#cache_storage_configurations}
	// Experimental.
	CacheStorageConfigurations interface{} `field:"optional" json:"cacheStorageConfigurations" yaml:"cacheStorageConfigurations"`
	// capacity_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#capacity_configuration AwsFinspaceKxCluster#capacity_configuration}
	// Experimental.
	CapacityConfiguration *AwsFinspaceKxCluster_CapacityConfigurationProperty `field:"optional" json:"capacityConfiguration" yaml:"capacityConfiguration"`
	// code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#code AwsFinspaceKxCluster#code}
	// Experimental.
	Code *AwsFinspaceKxCluster_CodeProperty `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#command_line_arguments AwsFinspaceKxCluster#command_line_arguments}.
	// Experimental.
	CommandLineArguments *map[string]*string `field:"optional" json:"commandLineArguments" yaml:"commandLineArguments"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#database AwsFinspaceKxCluster#database}
	// Experimental.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#description AwsFinspaceKxCluster#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#execution_role AwsFinspaceKxCluster#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#id AwsFinspaceKxCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#initialization_script AwsFinspaceKxCluster#initialization_script}.
	// Experimental.
	InitializationScript *string `field:"optional" json:"initializationScript" yaml:"initializationScript"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#region AwsFinspaceKxCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// savedown_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#savedown_storage_configuration AwsFinspaceKxCluster#savedown_storage_configuration}
	// Experimental.
	SavedownStorageConfiguration *AwsFinspaceKxCluster_SavedownStorageConfigurationProperty `field:"optional" json:"savedownStorageConfiguration" yaml:"savedownStorageConfiguration"`
	// scaling_group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#scaling_group_configuration AwsFinspaceKxCluster#scaling_group_configuration}
	// Experimental.
	ScalingGroupConfiguration *AwsFinspaceKxCluster_ScalingGroupConfigurationProperty `field:"optional" json:"scalingGroupConfiguration" yaml:"scalingGroupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#tags AwsFinspaceKxCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#tags_all AwsFinspaceKxCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tickerplant_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#tickerplant_log_configuration AwsFinspaceKxCluster#tickerplant_log_configuration}
	// Experimental.
	TickerplantLogConfiguration interface{} `field:"optional" json:"tickerplantLogConfiguration" yaml:"tickerplantLogConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#timeouts AwsFinspaceKxCluster#timeouts}
	// Experimental.
	Timeouts *AwsFinspaceKxCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

