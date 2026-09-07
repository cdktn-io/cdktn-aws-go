package endusermessagingsms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPhoneNumberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#iso_country_code AwsPhoneNumber#iso_country_code}.
	// Experimental.
	IsoCountryCode *string `field:"required" json:"isoCountryCode" yaml:"isoCountryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#message_type AwsPhoneNumber#message_type}.
	// Experimental.
	MessageType *string `field:"required" json:"messageType" yaml:"messageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#number_capabilities AwsPhoneNumber#number_capabilities}.
	// Experimental.
	NumberCapabilities *[]*string `field:"required" json:"numberCapabilities" yaml:"numberCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#number_type AwsPhoneNumber#number_type}.
	// Experimental.
	NumberType *string `field:"required" json:"numberType" yaml:"numberType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#deletion_protection_enabled AwsPhoneNumber#deletion_protection_enabled}.
	// Experimental.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#force_disassociate AwsPhoneNumber#force_disassociate}.
	// Experimental.
	ForceDisassociate interface{} `field:"optional" json:"forceDisassociate" yaml:"forceDisassociate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#opt_out_list_name AwsPhoneNumber#opt_out_list_name}.
	// Experimental.
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#region AwsPhoneNumber#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#registration_id AwsPhoneNumber#registration_id}.
	// Experimental.
	RegistrationId *string `field:"optional" json:"registrationId" yaml:"registrationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#self_managed_opt_outs_enabled AwsPhoneNumber#self_managed_opt_outs_enabled}.
	// Experimental.
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#tags AwsPhoneNumber#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#timeouts AwsPhoneNumber#timeouts}
	// Experimental.
	Timeouts *AwsPhoneNumber_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#two_way_channel_arn AwsPhoneNumber#two_way_channel_arn}.
	// Experimental.
	TwoWayChannelArn *string `field:"optional" json:"twoWayChannelArn" yaml:"twoWayChannelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#two_way_channel_enabled AwsPhoneNumber#two_way_channel_enabled}.
	// Experimental.
	TwoWayChannelEnabled interface{} `field:"optional" json:"twoWayChannelEnabled" yaml:"twoWayChannelEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#two_way_channel_role AwsPhoneNumber#two_way_channel_role}.
	// Experimental.
	TwoWayChannelRole *string `field:"optional" json:"twoWayChannelRole" yaml:"twoWayChannelRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_phone_number#wait_for_active AwsPhoneNumber#wait_for_active}.
	// Experimental.
	WaitForActive interface{} `field:"optional" json:"waitForActive" yaml:"waitForActive"`
}

