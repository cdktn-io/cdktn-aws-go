package awsendusermessaging


// Experimental.
type TfEmailTemplate_EmailTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_email_template#default_substitutions TfEmailTemplate#default_substitutions}.
	// Experimental.
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_email_template#description TfEmailTemplate#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_email_template#header TfEmailTemplate#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_email_template#html_part TfEmailTemplate#html_part}.
	// Experimental.
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_email_template#recommender_id TfEmailTemplate#recommender_id}.
	// Experimental.
	RecommenderId *string `field:"optional" json:"recommenderId" yaml:"recommenderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_email_template#subject TfEmailTemplate#subject}.
	// Experimental.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_email_template#text_part TfEmailTemplate#text_part}.
	// Experimental.
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

