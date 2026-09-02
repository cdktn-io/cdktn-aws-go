package awssns

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSmsPreferencesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#default_sender_id TfSmsPreferences#default_sender_id}.
	// Experimental.
	DefaultSenderId *string `field:"optional" json:"defaultSenderId" yaml:"defaultSenderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#default_sms_type TfSmsPreferences#default_sms_type}.
	// Experimental.
	DefaultSmsType *string `field:"optional" json:"defaultSmsType" yaml:"defaultSmsType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#delivery_status_iam_role_arn TfSmsPreferences#delivery_status_iam_role_arn}.
	// Experimental.
	DeliveryStatusIamRoleArn *string `field:"optional" json:"deliveryStatusIamRoleArn" yaml:"deliveryStatusIamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#delivery_status_success_sampling_rate TfSmsPreferences#delivery_status_success_sampling_rate}.
	// Experimental.
	DeliveryStatusSuccessSamplingRate *string `field:"optional" json:"deliveryStatusSuccessSamplingRate" yaml:"deliveryStatusSuccessSamplingRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#id TfSmsPreferences#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#monthly_spend_limit TfSmsPreferences#monthly_spend_limit}.
	// Experimental.
	MonthlySpendLimit *float64 `field:"optional" json:"monthlySpendLimit" yaml:"monthlySpendLimit"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#region TfSmsPreferences#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_sms_preferences#usage_report_s3_bucket TfSmsPreferences#usage_report_s3_bucket}.
	// Experimental.
	UsageReportS3Bucket *string `field:"optional" json:"usageReportS3Bucket" yaml:"usageReportS3Bucket"`
}

