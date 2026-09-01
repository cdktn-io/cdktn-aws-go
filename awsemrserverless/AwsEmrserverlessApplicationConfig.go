package awsemrserverless

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmrserverlessApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#name AwsEmrserverlessApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#release_label AwsEmrserverlessApplication#release_label}.
	// Experimental.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#type AwsEmrserverlessApplication#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#architecture AwsEmrserverlessApplication#architecture}.
	// Experimental.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// auto_start_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#auto_start_configuration AwsEmrserverlessApplication#auto_start_configuration}
	// Experimental.
	AutoStartConfiguration *AwsEmrserverlessApplication_AutoStartConfigurationProperty `field:"optional" json:"autoStartConfiguration" yaml:"autoStartConfiguration"`
	// auto_stop_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#auto_stop_configuration AwsEmrserverlessApplication#auto_stop_configuration}
	// Experimental.
	AutoStopConfiguration *AwsEmrserverlessApplication_AutoStopConfigurationProperty `field:"optional" json:"autoStopConfiguration" yaml:"autoStopConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#id AwsEmrserverlessApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#image_configuration AwsEmrserverlessApplication#image_configuration}
	// Experimental.
	ImageConfiguration *AwsEmrserverlessApplication_ImageConfigurationProperty `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// initial_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#initial_capacity AwsEmrserverlessApplication#initial_capacity}
	// Experimental.
	InitialCapacity interface{} `field:"optional" json:"initialCapacity" yaml:"initialCapacity"`
	// interactive_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#interactive_configuration AwsEmrserverlessApplication#interactive_configuration}
	// Experimental.
	InteractiveConfiguration *AwsEmrserverlessApplication_InteractiveConfigurationProperty `field:"optional" json:"interactiveConfiguration" yaml:"interactiveConfiguration"`
	// job_level_cost_allocation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#job_level_cost_allocation_configuration AwsEmrserverlessApplication#job_level_cost_allocation_configuration}
	// Experimental.
	JobLevelCostAllocationConfiguration *AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationProperty `field:"optional" json:"jobLevelCostAllocationConfiguration" yaml:"jobLevelCostAllocationConfiguration"`
	// maximum_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#maximum_capacity AwsEmrserverlessApplication#maximum_capacity}
	// Experimental.
	MaximumCapacity *AwsEmrserverlessApplication_MaximumCapacityProperty `field:"optional" json:"maximumCapacity" yaml:"maximumCapacity"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#monitoring_configuration AwsEmrserverlessApplication#monitoring_configuration}
	// Experimental.
	MonitoringConfiguration *AwsEmrserverlessApplication_MonitoringConfigurationProperty `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#network_configuration AwsEmrserverlessApplication#network_configuration}
	// Experimental.
	NetworkConfiguration *AwsEmrserverlessApplication_NetworkConfigurationProperty `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#region AwsEmrserverlessApplication#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// runtime_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#runtime_configuration AwsEmrserverlessApplication#runtime_configuration}
	// Experimental.
	RuntimeConfiguration interface{} `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// scheduler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#scheduler_configuration AwsEmrserverlessApplication#scheduler_configuration}
	// Experimental.
	SchedulerConfiguration *AwsEmrserverlessApplication_SchedulerConfigurationProperty `field:"optional" json:"schedulerConfiguration" yaml:"schedulerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#tags AwsEmrserverlessApplication#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#tags_all AwsEmrserverlessApplication#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

