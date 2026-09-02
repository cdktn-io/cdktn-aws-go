package awsquicksight

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAccountSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#account_name TfAccountSubscription#account_name}.
	// Experimental.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#authentication_method TfAccountSubscription#authentication_method}.
	// Experimental.
	AuthenticationMethod *string `field:"required" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#edition TfAccountSubscription#edition}.
	// Experimental.
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#notification_email TfAccountSubscription#notification_email}.
	// Experimental.
	NotificationEmail *string `field:"required" json:"notificationEmail" yaml:"notificationEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#active_directory_name TfAccountSubscription#active_directory_name}.
	// Experimental.
	ActiveDirectoryName *string `field:"optional" json:"activeDirectoryName" yaml:"activeDirectoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#admin_group TfAccountSubscription#admin_group}.
	// Experimental.
	AdminGroup *[]*string `field:"optional" json:"adminGroup" yaml:"adminGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#admin_pro_group TfAccountSubscription#admin_pro_group}.
	// Experimental.
	AdminProGroup *[]*string `field:"optional" json:"adminProGroup" yaml:"adminProGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#author_group TfAccountSubscription#author_group}.
	// Experimental.
	AuthorGroup *[]*string `field:"optional" json:"authorGroup" yaml:"authorGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#author_pro_group TfAccountSubscription#author_pro_group}.
	// Experimental.
	AuthorProGroup *[]*string `field:"optional" json:"authorProGroup" yaml:"authorProGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#aws_account_id TfAccountSubscription#aws_account_id}.
	// Experimental.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#contact_number TfAccountSubscription#contact_number}.
	// Experimental.
	ContactNumber *string `field:"optional" json:"contactNumber" yaml:"contactNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#directory_id TfAccountSubscription#directory_id}.
	// Experimental.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#email_address TfAccountSubscription#email_address}.
	// Experimental.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#first_name TfAccountSubscription#first_name}.
	// Experimental.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#iam_identity_center_instance_arn TfAccountSubscription#iam_identity_center_instance_arn}.
	// Experimental.
	IamIdentityCenterInstanceArn *string `field:"optional" json:"iamIdentityCenterInstanceArn" yaml:"iamIdentityCenterInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#id TfAccountSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#last_name TfAccountSubscription#last_name}.
	// Experimental.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#reader_group TfAccountSubscription#reader_group}.
	// Experimental.
	ReaderGroup *[]*string `field:"optional" json:"readerGroup" yaml:"readerGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#reader_pro_group TfAccountSubscription#reader_pro_group}.
	// Experimental.
	ReaderProGroup *[]*string `field:"optional" json:"readerProGroup" yaml:"readerProGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#realm TfAccountSubscription#realm}.
	// Experimental.
	Realm *string `field:"optional" json:"realm" yaml:"realm"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#region TfAccountSubscription#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#timeouts TfAccountSubscription#timeouts}
	// Experimental.
	Timeouts *TfAccountSubscription_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

