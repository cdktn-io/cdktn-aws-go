package finspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKxClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#az_mode AwsKxCluster#az_mode}.
	// Experimental.
	AzMode *string `field:"required" json:"azMode" yaml:"azMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#environment_id AwsKxCluster#environment_id}.
	// Experimental.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#name AwsKxCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#release_label AwsKxCluster#release_label}.
	// Experimental.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#type AwsKxCluster#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#vpc_configuration AwsKxCluster#vpc_configuration}
	// Experimental.
	VpcConfiguration *AwsKxCluster_VpcConfigurationProperty `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// auto_scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#auto_scaling_configuration AwsKxCluster#auto_scaling_configuration}
	// Experimental.
	AutoScalingConfiguration *AwsKxCluster_AutoScalingConfigurationProperty `field:"optional" json:"autoScalingConfiguration" yaml:"autoScalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#availability_zone_id AwsKxCluster#availability_zone_id}.
	// Experimental.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// cache_storage_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#cache_storage_configurations AwsKxCluster#cache_storage_configurations}
	// Experimental.
	CacheStorageConfigurations interface{} `field:"optional" json:"cacheStorageConfigurations" yaml:"cacheStorageConfigurations"`
	// capacity_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#capacity_configuration AwsKxCluster#capacity_configuration}
	// Experimental.
	CapacityConfiguration *AwsKxCluster_CapacityConfigurationProperty `field:"optional" json:"capacityConfiguration" yaml:"capacityConfiguration"`
	// code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#code AwsKxCluster#code}
	// Experimental.
	Code *AwsKxCluster_CodeProperty `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#command_line_arguments AwsKxCluster#command_line_arguments}.
	// Experimental.
	CommandLineArguments *map[string]*string `field:"optional" json:"commandLineArguments" yaml:"commandLineArguments"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#database AwsKxCluster#database}
	// Experimental.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#description AwsKxCluster#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#execution_role AwsKxCluster#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#id AwsKxCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#initialization_script AwsKxCluster#initialization_script}.
	// Experimental.
	InitializationScript *string `field:"optional" json:"initializationScript" yaml:"initializationScript"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#region AwsKxCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// savedown_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#savedown_storage_configuration AwsKxCluster#savedown_storage_configuration}
	// Experimental.
	SavedownStorageConfiguration *AwsKxCluster_SavedownStorageConfigurationProperty `field:"optional" json:"savedownStorageConfiguration" yaml:"savedownStorageConfiguration"`
	// scaling_group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#scaling_group_configuration AwsKxCluster#scaling_group_configuration}
	// Experimental.
	ScalingGroupConfiguration *AwsKxCluster_ScalingGroupConfigurationProperty `field:"optional" json:"scalingGroupConfiguration" yaml:"scalingGroupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#tags AwsKxCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#tags_all AwsKxCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tickerplant_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#tickerplant_log_configuration AwsKxCluster#tickerplant_log_configuration}
	// Experimental.
	TickerplantLogConfiguration interface{} `field:"optional" json:"tickerplantLogConfiguration" yaml:"tickerplantLogConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#timeouts AwsKxCluster#timeouts}
	// Experimental.
	Timeouts *AwsKxCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

