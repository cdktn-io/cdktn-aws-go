package rds


// Experimental.
type AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#connection_borrow_timeout AwsDbProxyDefaultTargetGroup#connection_borrow_timeout}.
	// Experimental.
	ConnectionBorrowTimeout *float64 `field:"optional" json:"connectionBorrowTimeout" yaml:"connectionBorrowTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#init_query AwsDbProxyDefaultTargetGroup#init_query}.
	// Experimental.
	InitQuery *string `field:"optional" json:"initQuery" yaml:"initQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#max_connections_percent AwsDbProxyDefaultTargetGroup#max_connections_percent}.
	// Experimental.
	MaxConnectionsPercent *float64 `field:"optional" json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#max_idle_connections_percent AwsDbProxyDefaultTargetGroup#max_idle_connections_percent}.
	// Experimental.
	MaxIdleConnectionsPercent *float64 `field:"optional" json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#session_pinning_filters AwsDbProxyDefaultTargetGroup#session_pinning_filters}.
	// Experimental.
	SessionPinningFilters *[]*string `field:"optional" json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
}

