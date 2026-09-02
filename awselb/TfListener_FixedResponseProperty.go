package awselb


// Experimental.
type TfListener_FixedResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#content_type TfListener#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#message_body TfListener#message_body}.
	// Experimental.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#status_code TfListener#status_code}.
	// Experimental.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

