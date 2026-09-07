package verifiedaccess


// Experimental.
type AwsEndpoint_RdsOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#port AwsEndpoint#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#protocol AwsEndpoint#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#rds_db_cluster_arn AwsEndpoint#rds_db_cluster_arn}.
	// Experimental.
	RdsDbClusterArn *string `field:"optional" json:"rdsDbClusterArn" yaml:"rdsDbClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#rds_db_instance_arn AwsEndpoint#rds_db_instance_arn}.
	// Experimental.
	RdsDbInstanceArn *string `field:"optional" json:"rdsDbInstanceArn" yaml:"rdsDbInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#rds_db_proxy_arn AwsEndpoint#rds_db_proxy_arn}.
	// Experimental.
	RdsDbProxyArn *string `field:"optional" json:"rdsDbProxyArn" yaml:"rdsDbProxyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#rds_endpoint AwsEndpoint#rds_endpoint}.
	// Experimental.
	RdsEndpoint *string `field:"optional" json:"rdsEndpoint" yaml:"rdsEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#subnet_ids AwsEndpoint#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

