package awsdatasync

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDatasyncLocationHdfsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#agent_arns AwsDatasyncLocationHdfs#agent_arns}.
	// Experimental.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// name_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#name_node AwsDatasyncLocationHdfs#name_node}
	// Experimental.
	NameNode interface{} `field:"required" json:"nameNode" yaml:"nameNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#authentication_type AwsDatasyncLocationHdfs#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#block_size AwsDatasyncLocationHdfs#block_size}.
	// Experimental.
	BlockSize *float64 `field:"optional" json:"blockSize" yaml:"blockSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#id AwsDatasyncLocationHdfs#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#kerberos_keytab AwsDatasyncLocationHdfs#kerberos_keytab}.
	// Experimental.
	KerberosKeytab *string `field:"optional" json:"kerberosKeytab" yaml:"kerberosKeytab"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#kerberos_keytab_base64 AwsDatasyncLocationHdfs#kerberos_keytab_base64}.
	// Experimental.
	KerberosKeytabBase64 *string `field:"optional" json:"kerberosKeytabBase64" yaml:"kerberosKeytabBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#kerberos_krb5_conf AwsDatasyncLocationHdfs#kerberos_krb5_conf}.
	// Experimental.
	KerberosKrb5Conf *string `field:"optional" json:"kerberosKrb5Conf" yaml:"kerberosKrb5Conf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#kerberos_krb5_conf_base64 AwsDatasyncLocationHdfs#kerberos_krb5_conf_base64}.
	// Experimental.
	KerberosKrb5ConfBase64 *string `field:"optional" json:"kerberosKrb5ConfBase64" yaml:"kerberosKrb5ConfBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#kerberos_principal AwsDatasyncLocationHdfs#kerberos_principal}.
	// Experimental.
	KerberosPrincipal *string `field:"optional" json:"kerberosPrincipal" yaml:"kerberosPrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#kms_key_provider_uri AwsDatasyncLocationHdfs#kms_key_provider_uri}.
	// Experimental.
	KmsKeyProviderUri *string `field:"optional" json:"kmsKeyProviderUri" yaml:"kmsKeyProviderUri"`
	// qop_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#qop_configuration AwsDatasyncLocationHdfs#qop_configuration}
	// Experimental.
	QopConfiguration *AwsDatasyncLocationHdfs_QopConfigurationProperty `field:"optional" json:"qopConfiguration" yaml:"qopConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#region AwsDatasyncLocationHdfs#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#replication_factor AwsDatasyncLocationHdfs#replication_factor}.
	// Experimental.
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#simple_user AwsDatasyncLocationHdfs#simple_user}.
	// Experimental.
	SimpleUser *string `field:"optional" json:"simpleUser" yaml:"simpleUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#subdirectory AwsDatasyncLocationHdfs#subdirectory}.
	// Experimental.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#tags AwsDatasyncLocationHdfs#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#tags_all AwsDatasyncLocationHdfs#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

