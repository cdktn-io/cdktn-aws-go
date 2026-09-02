package awsrds


// Experimental.
type TfDbProxyDefaultTargetGroup_ConnectionPoolConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#connection_borrow_timeout TfDbProxyDefaultTargetGroup#connection_borrow_timeout}.
	// Experimental.
	ConnectionBorrowTimeout *float64 `field:"optional" json:"connectionBorrowTimeout" yaml:"connectionBorrowTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#init_query TfDbProxyDefaultTargetGroup#init_query}.
	// Experimental.
	InitQuery *string `field:"optional" json:"initQuery" yaml:"initQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#max_connections_percent TfDbProxyDefaultTargetGroup#max_connections_percent}.
	// Experimental.
	MaxConnectionsPercent *float64 `field:"optional" json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#max_idle_connections_percent TfDbProxyDefaultTargetGroup#max_idle_connections_percent}.
	// Experimental.
	MaxIdleConnectionsPercent *float64 `field:"optional" json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy_default_target_group#session_pinning_filters TfDbProxyDefaultTargetGroup#session_pinning_filters}.
	// Experimental.
	SessionPinningFilters *[]*string `field:"optional" json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
}

