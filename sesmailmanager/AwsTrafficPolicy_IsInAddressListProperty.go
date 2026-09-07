package sesmailmanager


// Experimental.
type AwsTrafficPolicy_IsInAddressListProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#address_lists AwsTrafficPolicy#address_lists}.
	// Experimental.
	AddressLists *[]*string `field:"required" json:"addressLists" yaml:"addressLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#attribute AwsTrafficPolicy#attribute}.
	// Experimental.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

