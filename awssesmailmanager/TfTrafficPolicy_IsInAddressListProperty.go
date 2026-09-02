package awssesmailmanager


// Experimental.
type TfTrafficPolicy_IsInAddressListProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#address_lists TfTrafficPolicy#address_lists}.
	// Experimental.
	AddressLists *[]*string `field:"required" json:"addressLists" yaml:"addressLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#attribute TfTrafficPolicy#attribute}.
	// Experimental.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

