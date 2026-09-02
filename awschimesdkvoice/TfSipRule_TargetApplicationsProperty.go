package awschimesdkvoice


// Experimental.
type TfSipRule_TargetApplicationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkvoice_sip_rule#aws_region TfSipRule#aws_region}.
	// Experimental.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkvoice_sip_rule#priority TfSipRule#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkvoice_sip_rule#sip_media_application_id TfSipRule#sip_media_application_id}.
	// Experimental.
	SipMediaApplicationId *string `field:"required" json:"sipMediaApplicationId" yaml:"sipMediaApplicationId"`
}

