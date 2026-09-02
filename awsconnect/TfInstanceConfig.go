package awsconnect

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#identity_management_type TfInstance#identity_management_type}.
	// Experimental.
	IdentityManagementType *string `field:"required" json:"identityManagementType" yaml:"identityManagementType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#inbound_calls_enabled TfInstance#inbound_calls_enabled}.
	// Experimental.
	InboundCallsEnabled interface{} `field:"required" json:"inboundCallsEnabled" yaml:"inboundCallsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#outbound_calls_enabled TfInstance#outbound_calls_enabled}.
	// Experimental.
	OutboundCallsEnabled interface{} `field:"required" json:"outboundCallsEnabled" yaml:"outboundCallsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#auto_resolve_best_voices_enabled TfInstance#auto_resolve_best_voices_enabled}.
	// Experimental.
	AutoResolveBestVoicesEnabled interface{} `field:"optional" json:"autoResolveBestVoicesEnabled" yaml:"autoResolveBestVoicesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#contact_flow_logs_enabled TfInstance#contact_flow_logs_enabled}.
	// Experimental.
	ContactFlowLogsEnabled interface{} `field:"optional" json:"contactFlowLogsEnabled" yaml:"contactFlowLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#contact_lens_enabled TfInstance#contact_lens_enabled}.
	// Experimental.
	ContactLensEnabled interface{} `field:"optional" json:"contactLensEnabled" yaml:"contactLensEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#directory_id TfInstance#directory_id}.
	// Experimental.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#early_media_enabled TfInstance#early_media_enabled}.
	// Experimental.
	EarlyMediaEnabled interface{} `field:"optional" json:"earlyMediaEnabled" yaml:"earlyMediaEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#id TfInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#instance_alias TfInstance#instance_alias}.
	// Experimental.
	InstanceAlias *string `field:"optional" json:"instanceAlias" yaml:"instanceAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#multi_party_conference_enabled TfInstance#multi_party_conference_enabled}.
	// Experimental.
	MultiPartyConferenceEnabled interface{} `field:"optional" json:"multiPartyConferenceEnabled" yaml:"multiPartyConferenceEnabled"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#region TfInstance#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#tags TfInstance#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#tags_all TfInstance#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#timeouts TfInstance#timeouts}
	// Experimental.
	Timeouts *TfInstance_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

