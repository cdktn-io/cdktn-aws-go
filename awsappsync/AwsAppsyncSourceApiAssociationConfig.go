package awsappsync

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppsyncSourceApiAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#description AwsAppsyncSourceApiAssociation#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#merged_api_arn AwsAppsyncSourceApiAssociation#merged_api_arn}.
	// Experimental.
	MergedApiArn *string `field:"optional" json:"mergedApiArn" yaml:"mergedApiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#merged_api_id AwsAppsyncSourceApiAssociation#merged_api_id}.
	// Experimental.
	MergedApiId *string `field:"optional" json:"mergedApiId" yaml:"mergedApiId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#region AwsAppsyncSourceApiAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#source_api_arn AwsAppsyncSourceApiAssociation#source_api_arn}.
	// Experimental.
	SourceApiArn *string `field:"optional" json:"sourceApiArn" yaml:"sourceApiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#source_api_association_config AwsAppsyncSourceApiAssociation#source_api_association_config}.
	// Experimental.
	SourceApiAssociationConfig interface{} `field:"optional" json:"sourceApiAssociationConfig" yaml:"sourceApiAssociationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#source_api_id AwsAppsyncSourceApiAssociation#source_api_id}.
	// Experimental.
	SourceApiId *string `field:"optional" json:"sourceApiId" yaml:"sourceApiId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_source_api_association#timeouts AwsAppsyncSourceApiAssociation#timeouts}
	// Experimental.
	Timeouts *AwsAppsyncSourceApiAssociation_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

