package awselb


// Experimental.
type TfAlbListener_FixedResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#content_type TfAlbListener#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#message_body TfAlbListener#message_body}.
	// Experimental.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#status_code TfAlbListener#status_code}.
	// Experimental.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

