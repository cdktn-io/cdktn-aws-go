package awselb


// Experimental.
type TfAlbListenerRule_FixedResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#content_type TfAlbListenerRule#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#message_body TfAlbListenerRule#message_body}.
	// Experimental.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#status_code TfAlbListenerRule#status_code}.
	// Experimental.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

