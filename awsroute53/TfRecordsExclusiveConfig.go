package awsroute53

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRecordsExclusiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#zone_id TfRecordsExclusive#zone_id}.
	// Experimental.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// resource_record_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#resource_record_set TfRecordsExclusive#resource_record_set}
	// Experimental.
	ResourceRecordSet interface{} `field:"optional" json:"resourceRecordSet" yaml:"resourceRecordSet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#timeouts TfRecordsExclusive#timeouts}
	// Experimental.
	Timeouts *TfRecordsExclusive_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

