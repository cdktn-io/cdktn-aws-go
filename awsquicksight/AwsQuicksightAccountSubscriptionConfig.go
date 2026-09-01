package awsquicksight

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightAccountSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#account_name AwsQuicksightAccountSubscription#account_name}.
	// Experimental.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#authentication_method AwsQuicksightAccountSubscription#authentication_method}.
	// Experimental.
	AuthenticationMethod *string `field:"required" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#edition AwsQuicksightAccountSubscription#edition}.
	// Experimental.
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#notification_email AwsQuicksightAccountSubscription#notification_email}.
	// Experimental.
	NotificationEmail *string `field:"required" json:"notificationEmail" yaml:"notificationEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#active_directory_name AwsQuicksightAccountSubscription#active_directory_name}.
	// Experimental.
	ActiveDirectoryName *string `field:"optional" json:"activeDirectoryName" yaml:"activeDirectoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#admin_group AwsQuicksightAccountSubscription#admin_group}.
	// Experimental.
	AdminGroup *[]*string `field:"optional" json:"adminGroup" yaml:"adminGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#admin_pro_group AwsQuicksightAccountSubscription#admin_pro_group}.
	// Experimental.
	AdminProGroup *[]*string `field:"optional" json:"adminProGroup" yaml:"adminProGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#author_group AwsQuicksightAccountSubscription#author_group}.
	// Experimental.
	AuthorGroup *[]*string `field:"optional" json:"authorGroup" yaml:"authorGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#author_pro_group AwsQuicksightAccountSubscription#author_pro_group}.
	// Experimental.
	AuthorProGroup *[]*string `field:"optional" json:"authorProGroup" yaml:"authorProGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#aws_account_id AwsQuicksightAccountSubscription#aws_account_id}.
	// Experimental.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#contact_number AwsQuicksightAccountSubscription#contact_number}.
	// Experimental.
	ContactNumber *string `field:"optional" json:"contactNumber" yaml:"contactNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#directory_id AwsQuicksightAccountSubscription#directory_id}.
	// Experimental.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#email_address AwsQuicksightAccountSubscription#email_address}.
	// Experimental.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#first_name AwsQuicksightAccountSubscription#first_name}.
	// Experimental.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#iam_identity_center_instance_arn AwsQuicksightAccountSubscription#iam_identity_center_instance_arn}.
	// Experimental.
	IamIdentityCenterInstanceArn *string `field:"optional" json:"iamIdentityCenterInstanceArn" yaml:"iamIdentityCenterInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#id AwsQuicksightAccountSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#last_name AwsQuicksightAccountSubscription#last_name}.
	// Experimental.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#reader_group AwsQuicksightAccountSubscription#reader_group}.
	// Experimental.
	ReaderGroup *[]*string `field:"optional" json:"readerGroup" yaml:"readerGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#reader_pro_group AwsQuicksightAccountSubscription#reader_pro_group}.
	// Experimental.
	ReaderProGroup *[]*string `field:"optional" json:"readerProGroup" yaml:"readerProGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#realm AwsQuicksightAccountSubscription#realm}.
	// Experimental.
	Realm *string `field:"optional" json:"realm" yaml:"realm"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#region AwsQuicksightAccountSubscription#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#timeouts AwsQuicksightAccountSubscription#timeouts}
	// Experimental.
	Timeouts *AwsQuicksightAccountSubscription_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

