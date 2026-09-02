package awsroute53


// Experimental.
type TfRecord_CoordinatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#latitude TfRecord#latitude}.
	// Experimental.
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#longitude TfRecord#longitude}.
	// Experimental.
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
}

