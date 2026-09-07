package networkmanager


// Experimental.
type AwsSite_LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_site#address AwsSite#address}.
	// Experimental.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_site#latitude AwsSite#latitude}.
	// Experimental.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_site#longitude AwsSite#longitude}.
	// Experimental.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

