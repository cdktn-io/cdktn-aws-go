package awsecs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTaskDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#container_definitions TfTaskDefinition#container_definitions}.
	// Experimental.
	ContainerDefinitions *string `field:"required" json:"containerDefinitions" yaml:"containerDefinitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#family TfTaskDefinition#family}.
	// Experimental.
	Family *string `field:"required" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#cpu TfTaskDefinition#cpu}.
	// Experimental.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#enable_fault_injection TfTaskDefinition#enable_fault_injection}.
	// Experimental.
	EnableFaultInjection interface{} `field:"optional" json:"enableFaultInjection" yaml:"enableFaultInjection"`
	// ephemeral_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#ephemeral_storage TfTaskDefinition#ephemeral_storage}
	// Experimental.
	EphemeralStorage *TfTaskDefinition_EphemeralStorageProperty `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#execution_role_arn TfTaskDefinition#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#id TfTaskDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#ipc_mode TfTaskDefinition#ipc_mode}.
	// Experimental.
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#memory TfTaskDefinition#memory}.
	// Experimental.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#network_mode TfTaskDefinition#network_mode}.
	// Experimental.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#pid_mode TfTaskDefinition#pid_mode}.
	// Experimental.
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// placement_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#placement_constraints TfTaskDefinition#placement_constraints}
	// Experimental.
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// proxy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#proxy_configuration TfTaskDefinition#proxy_configuration}
	// Experimental.
	ProxyConfiguration *TfTaskDefinition_ProxyConfigurationProperty `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#region TfTaskDefinition#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#requires_compatibilities TfTaskDefinition#requires_compatibilities}.
	// Experimental.
	RequiresCompatibilities *[]*string `field:"optional" json:"requiresCompatibilities" yaml:"requiresCompatibilities"`
	// runtime_platform block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#runtime_platform TfTaskDefinition#runtime_platform}
	// Experimental.
	RuntimePlatform *TfTaskDefinition_RuntimePlatformProperty `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#skip_destroy TfTaskDefinition#skip_destroy}.
	// Experimental.
	SkipDestroy interface{} `field:"optional" json:"skipDestroy" yaml:"skipDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#tags TfTaskDefinition#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#tags_all TfTaskDefinition#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#task_role_arn TfTaskDefinition#task_role_arn}.
	// Experimental.
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#track_latest TfTaskDefinition#track_latest}.
	// Experimental.
	TrackLatest interface{} `field:"optional" json:"trackLatest" yaml:"trackLatest"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#volume TfTaskDefinition#volume}
	// Experimental.
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

