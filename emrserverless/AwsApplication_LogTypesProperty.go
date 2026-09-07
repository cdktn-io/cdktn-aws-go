package emrserverless


// Experimental.
type AwsApplication_LogTypesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#name AwsApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#values AwsApplication#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

