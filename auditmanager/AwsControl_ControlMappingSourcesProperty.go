package auditmanager


// Experimental.
type AwsControl_ControlMappingSourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_control#source_name AwsControl#source_name}.
	// Experimental.
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_control#source_set_up_option AwsControl#source_set_up_option}.
	// Experimental.
	SourceSetUpOption *string `field:"required" json:"sourceSetUpOption" yaml:"sourceSetUpOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_control#source_type AwsControl#source_type}.
	// Experimental.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_control#source_description AwsControl#source_description}.
	// Experimental.
	SourceDescription *string `field:"optional" json:"sourceDescription" yaml:"sourceDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_control#source_frequency AwsControl#source_frequency}.
	// Experimental.
	SourceFrequency *string `field:"optional" json:"sourceFrequency" yaml:"sourceFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_control#source_keyword AwsControl#source_keyword}.
	// Experimental.
	SourceKeyword interface{} `field:"optional" json:"sourceKeyword" yaml:"sourceKeyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_control#troubleshooting_text AwsControl#troubleshooting_text}.
	// Experimental.
	TroubleshootingText *string `field:"optional" json:"troubleshootingText" yaml:"troubleshootingText"`
}

