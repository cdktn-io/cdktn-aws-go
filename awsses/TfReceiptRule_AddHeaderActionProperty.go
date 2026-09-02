package awsses


// Experimental.
type TfReceiptRule_AddHeaderActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#header_name TfReceiptRule#header_name}.
	// Experimental.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#header_value TfReceiptRule#header_value}.
	// Experimental.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#position TfReceiptRule#position}.
	// Experimental.
	Position *float64 `field:"required" json:"position" yaml:"position"`
}

