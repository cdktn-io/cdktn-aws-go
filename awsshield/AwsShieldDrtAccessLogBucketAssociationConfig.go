package awsshield

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsShieldDrtAccessLogBucketAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_drt_access_log_bucket_association#log_bucket AwsShieldDrtAccessLogBucketAssociation#log_bucket}.
	// Experimental.
	LogBucket *string `field:"required" json:"logBucket" yaml:"logBucket"`
	// Unused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_drt_access_log_bucket_association#role_arn_association_id AwsShieldDrtAccessLogBucketAssociation#role_arn_association_id}
	// Experimental.
	RoleArnAssociationId *string `field:"required" json:"roleArnAssociationId" yaml:"roleArnAssociationId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_drt_access_log_bucket_association#timeouts AwsShieldDrtAccessLogBucketAssociation#timeouts}
	// Experimental.
	Timeouts *AwsShieldDrtAccessLogBucketAssociation_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

