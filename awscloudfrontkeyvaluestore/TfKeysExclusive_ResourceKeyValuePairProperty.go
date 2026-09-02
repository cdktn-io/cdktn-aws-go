package awscloudfrontkeyvaluestore


// Experimental.
type TfKeysExclusive_ResourceKeyValuePairProperty struct {
	// The key to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#key TfKeysExclusive#key}
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#value TfKeysExclusive#value}
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

