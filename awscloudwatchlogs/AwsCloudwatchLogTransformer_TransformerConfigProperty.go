package awscloudwatchlogs


// Experimental.
type AwsCloudwatchLogTransformer_TransformerConfigProperty struct {
	// add_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#add_keys AwsCloudwatchLogTransformer#add_keys}
	// Experimental.
	AddKeys interface{} `field:"optional" json:"addKeys" yaml:"addKeys"`
	// copy_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#copy_value AwsCloudwatchLogTransformer#copy_value}
	// Experimental.
	CopyValue interface{} `field:"optional" json:"copyValue" yaml:"copyValue"`
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#csv AwsCloudwatchLogTransformer#csv}
	// Experimental.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// date_time_converter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#date_time_converter AwsCloudwatchLogTransformer#date_time_converter}
	// Experimental.
	DateTimeConverter interface{} `field:"optional" json:"dateTimeConverter" yaml:"dateTimeConverter"`
	// delete_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#delete_keys AwsCloudwatchLogTransformer#delete_keys}
	// Experimental.
	DeleteKeys interface{} `field:"optional" json:"deleteKeys" yaml:"deleteKeys"`
	// grok block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#grok AwsCloudwatchLogTransformer#grok}
	// Experimental.
	Grok interface{} `field:"optional" json:"grok" yaml:"grok"`
	// list_to_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#list_to_map AwsCloudwatchLogTransformer#list_to_map}
	// Experimental.
	ListToMap interface{} `field:"optional" json:"listToMap" yaml:"listToMap"`
	// lower_case_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#lower_case_string AwsCloudwatchLogTransformer#lower_case_string}
	// Experimental.
	LowerCaseString interface{} `field:"optional" json:"lowerCaseString" yaml:"lowerCaseString"`
	// move_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#move_keys AwsCloudwatchLogTransformer#move_keys}
	// Experimental.
	MoveKeys interface{} `field:"optional" json:"moveKeys" yaml:"moveKeys"`
	// parse_cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_cloudfront AwsCloudwatchLogTransformer#parse_cloudfront}
	// Experimental.
	ParseCloudfront interface{} `field:"optional" json:"parseCloudfront" yaml:"parseCloudfront"`
	// parse_json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_json AwsCloudwatchLogTransformer#parse_json}
	// Experimental.
	ParseJson interface{} `field:"optional" json:"parseJson" yaml:"parseJson"`
	// parse_key_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_key_value AwsCloudwatchLogTransformer#parse_key_value}
	// Experimental.
	ParseKeyValue interface{} `field:"optional" json:"parseKeyValue" yaml:"parseKeyValue"`
	// parse_postgres block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_postgres AwsCloudwatchLogTransformer#parse_postgres}
	// Experimental.
	ParsePostgres interface{} `field:"optional" json:"parsePostgres" yaml:"parsePostgres"`
	// parse_route53 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_route53 AwsCloudwatchLogTransformer#parse_route53}
	// Experimental.
	ParseRoute53 interface{} `field:"optional" json:"parseRoute53" yaml:"parseRoute53"`
	// parse_to_ocsf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_to_ocsf AwsCloudwatchLogTransformer#parse_to_ocsf}
	// Experimental.
	ParseToOcsf interface{} `field:"optional" json:"parseToOcsf" yaml:"parseToOcsf"`
	// parse_vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_vpc AwsCloudwatchLogTransformer#parse_vpc}
	// Experimental.
	ParseVpc interface{} `field:"optional" json:"parseVpc" yaml:"parseVpc"`
	// parse_waf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#parse_waf AwsCloudwatchLogTransformer#parse_waf}
	// Experimental.
	ParseWaf interface{} `field:"optional" json:"parseWaf" yaml:"parseWaf"`
	// rename_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#rename_keys AwsCloudwatchLogTransformer#rename_keys}
	// Experimental.
	RenameKeys interface{} `field:"optional" json:"renameKeys" yaml:"renameKeys"`
	// split_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#split_string AwsCloudwatchLogTransformer#split_string}
	// Experimental.
	SplitString interface{} `field:"optional" json:"splitString" yaml:"splitString"`
	// substitute_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#substitute_string AwsCloudwatchLogTransformer#substitute_string}
	// Experimental.
	SubstituteString interface{} `field:"optional" json:"substituteString" yaml:"substituteString"`
	// trim_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#trim_string AwsCloudwatchLogTransformer#trim_string}
	// Experimental.
	TrimString interface{} `field:"optional" json:"trimString" yaml:"trimString"`
	// type_converter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#type_converter AwsCloudwatchLogTransformer#type_converter}
	// Experimental.
	TypeConverter interface{} `field:"optional" json:"typeConverter" yaml:"typeConverter"`
	// upper_case_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#upper_case_string AwsCloudwatchLogTransformer#upper_case_string}
	// Experimental.
	UpperCaseString interface{} `field:"optional" json:"upperCaseString" yaml:"upperCaseString"`
}

