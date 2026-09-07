package networkmanager


// Experimental.
type AwsDevice_LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_device#address AwsDevice#address}.
	// Experimental.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_device#latitude AwsDevice#latitude}.
	// Experimental.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_device#longitude AwsDevice#longitude}.
	// Experimental.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

