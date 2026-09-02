package awscloudwatchrum


// Experimental.
type TfAppMonitor_AppMonitorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#allow_cookies TfAppMonitor#allow_cookies}.
	// Experimental.
	AllowCookies interface{} `field:"optional" json:"allowCookies" yaml:"allowCookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#enable_xray TfAppMonitor#enable_xray}.
	// Experimental.
	EnableXray interface{} `field:"optional" json:"enableXray" yaml:"enableXray"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#excluded_pages TfAppMonitor#excluded_pages}.
	// Experimental.
	ExcludedPages *[]*string `field:"optional" json:"excludedPages" yaml:"excludedPages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#favorite_pages TfAppMonitor#favorite_pages}.
	// Experimental.
	FavoritePages *[]*string `field:"optional" json:"favoritePages" yaml:"favoritePages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#guest_role_arn TfAppMonitor#guest_role_arn}.
	// Experimental.
	GuestRoleArn *string `field:"optional" json:"guestRoleArn" yaml:"guestRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#identity_pool_id TfAppMonitor#identity_pool_id}.
	// Experimental.
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#included_pages TfAppMonitor#included_pages}.
	// Experimental.
	IncludedPages *[]*string `field:"optional" json:"includedPages" yaml:"includedPages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#session_sample_rate TfAppMonitor#session_sample_rate}.
	// Experimental.
	SessionSampleRate *float64 `field:"optional" json:"sessionSampleRate" yaml:"sessionSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rum_app_monitor#telemetries TfAppMonitor#telemetries}.
	// Experimental.
	Telemetries *[]*string `field:"optional" json:"telemetries" yaml:"telemetries"`
}

