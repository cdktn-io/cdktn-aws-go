package awsconnect


// Experimental.
type AwsConnectUserHierarchyStructure_HierarchyStructureProperty struct {
	// level_five block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_five AwsConnectUserHierarchyStructure#level_five}
	// Experimental.
	LevelFive *AwsConnectUserHierarchyStructure_LevelFiveProperty `field:"optional" json:"levelFive" yaml:"levelFive"`
	// level_four block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_four AwsConnectUserHierarchyStructure#level_four}
	// Experimental.
	LevelFour *AwsConnectUserHierarchyStructure_LevelFourProperty `field:"optional" json:"levelFour" yaml:"levelFour"`
	// level_one block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_one AwsConnectUserHierarchyStructure#level_one}
	// Experimental.
	LevelOne *AwsConnectUserHierarchyStructure_LevelOneProperty `field:"optional" json:"levelOne" yaml:"levelOne"`
	// level_three block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_three AwsConnectUserHierarchyStructure#level_three}
	// Experimental.
	LevelThree *AwsConnectUserHierarchyStructure_LevelThreeProperty `field:"optional" json:"levelThree" yaml:"levelThree"`
	// level_two block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_two AwsConnectUserHierarchyStructure#level_two}
	// Experimental.
	LevelTwo *AwsConnectUserHierarchyStructure_LevelTwoProperty `field:"optional" json:"levelTwo" yaml:"levelTwo"`
}

