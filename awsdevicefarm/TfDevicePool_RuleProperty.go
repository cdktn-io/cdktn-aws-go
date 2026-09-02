package awsdevicefarm


// Experimental.
type TfDevicePool_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_device_pool#attribute TfDevicePool#attribute}.
	// Experimental.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_device_pool#operator TfDevicePool#operator}.
	// Experimental.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_device_pool#value TfDevicePool#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

