package awscloudfrontkeyvaluestore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfKeyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The key to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfrontkeyvaluestore_key#key TfKey#key}
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Amazon Resource Name (ARN) of the Key Value Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfrontkeyvaluestore_key#key_value_store_arn TfKey#key_value_store_arn}
	// Experimental.
	KeyValueStoreArn *string `field:"required" json:"keyValueStoreArn" yaml:"keyValueStoreArn"`
	// The value to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfrontkeyvaluestore_key#value TfKey#value}
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

