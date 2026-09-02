package awsbatch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#name TfJobDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#type TfJobDefinition#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#container_properties TfJobDefinition#container_properties}.
	// Experimental.
	ContainerProperties *string `field:"optional" json:"containerProperties" yaml:"containerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#deregister_on_new_revision TfJobDefinition#deregister_on_new_revision}.
	// Experimental.
	DeregisterOnNewRevision interface{} `field:"optional" json:"deregisterOnNewRevision" yaml:"deregisterOnNewRevision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#ecs_properties TfJobDefinition#ecs_properties}.
	// Experimental.
	EcsProperties *string `field:"optional" json:"ecsProperties" yaml:"ecsProperties"`
	// eks_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#eks_properties TfJobDefinition#eks_properties}
	// Experimental.
	EksProperties *TfJobDefinition_EksPropertiesProperty `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#id TfJobDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#node_properties TfJobDefinition#node_properties}.
	// Experimental.
	NodeProperties *string `field:"optional" json:"nodeProperties" yaml:"nodeProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#parameters TfJobDefinition#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#platform_capabilities TfJobDefinition#platform_capabilities}.
	// Experimental.
	PlatformCapabilities *[]*string `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#propagate_tags TfJobDefinition#propagate_tags}.
	// Experimental.
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#region TfJobDefinition#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#retry_strategy TfJobDefinition#retry_strategy}
	// Experimental.
	RetryStrategy *TfJobDefinition_RetryStrategyProperty `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#scheduling_priority TfJobDefinition#scheduling_priority}.
	// Experimental.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#tags TfJobDefinition#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#tags_all TfJobDefinition#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#timeout TfJobDefinition#timeout}
	// Experimental.
	Timeout *TfJobDefinition_TimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

