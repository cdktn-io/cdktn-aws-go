package awsendusermessagingsms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPoolConfig struct {
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
	// Type of message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#message_type TfPool#message_type}
	// Experimental.
	MessageType *string `field:"required" json:"messageType" yaml:"messageType"`
	// Set of origination identity ARNs to associate with the pool. At least one origination identity is required at creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#origination_identities TfPool#origination_identities}
	// Experimental.
	OriginationIdentities *[]*string `field:"required" json:"originationIdentities" yaml:"originationIdentities"`
	// Whether deletion protection is enabled. When `true`, the pool cannot be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#deletion_protection_enabled TfPool#deletion_protection_enabled}
	// Experimental.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Two-character code, in ISO 3166-1 alpha-2 format, for the country or region of the pool.
	//
	// This field is optional for origination identity types that are not country-specific.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#iso_country_code TfPool#iso_country_code}
	// Experimental.
	IsoCountryCode *string `field:"optional" json:"isoCountryCode" yaml:"isoCountryCode"`
	// Name of the opt-out list to associate with the pool. Inherited from the initial origination identity when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#opt_out_list_name TfPool#opt_out_list_name}
	// Experimental.
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#region TfPool#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Whether the pool relies on self-managed opt-out handling.
	//
	// When `false`, AWS auto-replies to HELP/STOP requests and manages the opt-out list. Inherited from the initial origination identity when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#self_managed_opt_outs_enabled TfPool#self_managed_opt_outs_enabled}
	// Experimental.
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// Whether shared routes are enabled for the pool.
	//
	// When `true`, messages may use shared phone numbers or sender IDs in countries that allow it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#shared_routes_enabled TfPool#shared_routes_enabled}
	// Experimental.
	SharedRoutesEnabled interface{} `field:"optional" json:"sharedRoutesEnabled" yaml:"sharedRoutesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#tags TfPool#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#timeouts TfPool#timeouts}
	// Experimental.
	Timeouts *TfPool_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ARN of the two-way channel that receives inbound messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#two_way_channel_arn TfPool#two_way_channel_arn}
	// Experimental.
	TwoWayChannelArn *string `field:"optional" json:"twoWayChannelArn" yaml:"twoWayChannelArn"`
	// ARN of the IAM role that End User Messaging SMS assumes to publish inbound messages to the two-way channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#two_way_channel_role TfPool#two_way_channel_role}
	// Experimental.
	TwoWayChannelRole *string `field:"optional" json:"twoWayChannelRole" yaml:"twoWayChannelRole"`
	// Whether inbound message reception is enabled for the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool#two_way_enabled TfPool#two_way_enabled}
	// Experimental.
	TwoWayEnabled interface{} `field:"optional" json:"twoWayEnabled" yaml:"twoWayEnabled"`
}

