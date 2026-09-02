package awsinspector


// Experimental.
type TfFilter_VulnerablePackagesProperty struct {
	// architecture block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#architecture TfFilter#architecture}
	// Experimental.
	Architecture interface{} `field:"optional" json:"architecture" yaml:"architecture"`
	// epoch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#epoch TfFilter#epoch}
	// Experimental.
	Epoch interface{} `field:"optional" json:"epoch" yaml:"epoch"`
	// file_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#file_path TfFilter#file_path}
	// Experimental.
	FilePath interface{} `field:"optional" json:"filePath" yaml:"filePath"`
	// name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#name TfFilter#name}
	// Experimental.
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// release block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#release TfFilter#release}
	// Experimental.
	Release interface{} `field:"optional" json:"release" yaml:"release"`
	// source_lambda_layer_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#source_lambda_layer_arn TfFilter#source_lambda_layer_arn}
	// Experimental.
	SourceLambdaLayerArn interface{} `field:"optional" json:"sourceLambdaLayerArn" yaml:"sourceLambdaLayerArn"`
	// source_layer_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#source_layer_hash TfFilter#source_layer_hash}
	// Experimental.
	SourceLayerHash interface{} `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#version TfFilter#version}
	// Experimental.
	Version interface{} `field:"optional" json:"version" yaml:"version"`
}

