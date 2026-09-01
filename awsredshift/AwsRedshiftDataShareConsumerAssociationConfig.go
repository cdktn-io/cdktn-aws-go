package awsredshift

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRedshiftDataShareConsumerAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_data_share_consumer_association#data_share_arn AwsRedshiftDataShareConsumerAssociation#data_share_arn}.
	// Experimental.
	DataShareArn *string `field:"required" json:"dataShareArn" yaml:"dataShareArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_data_share_consumer_association#allow_writes AwsRedshiftDataShareConsumerAssociation#allow_writes}.
	// Experimental.
	AllowWrites interface{} `field:"optional" json:"allowWrites" yaml:"allowWrites"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_data_share_consumer_association#associate_entire_account AwsRedshiftDataShareConsumerAssociation#associate_entire_account}.
	// Experimental.
	AssociateEntireAccount interface{} `field:"optional" json:"associateEntireAccount" yaml:"associateEntireAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_data_share_consumer_association#consumer_arn AwsRedshiftDataShareConsumerAssociation#consumer_arn}.
	// Experimental.
	ConsumerArn *string `field:"optional" json:"consumerArn" yaml:"consumerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_data_share_consumer_association#consumer_region AwsRedshiftDataShareConsumerAssociation#consumer_region}.
	// Experimental.
	ConsumerRegion *string `field:"optional" json:"consumerRegion" yaml:"consumerRegion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_data_share_consumer_association#region AwsRedshiftDataShareConsumerAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

