package route53


// Experimental.
type AwsRecord_CoordinatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#latitude AwsRecord#latitude}.
	// Experimental.
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#longitude AwsRecord#longitude}.
	// Experimental.
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
}

