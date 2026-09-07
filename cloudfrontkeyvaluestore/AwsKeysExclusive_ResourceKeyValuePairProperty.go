package cloudfrontkeyvaluestore


// Experimental.
type AwsKeysExclusive_ResourceKeyValuePairProperty struct {
	// The key to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#key AwsKeysExclusive#key}
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#value AwsKeysExclusive#value}
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

