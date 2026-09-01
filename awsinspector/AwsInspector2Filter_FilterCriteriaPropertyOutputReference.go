package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsinspector/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsinspector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsInspector2Filter_FilterCriteriaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() AwsInspector2Filter_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	CodeRepositoryProjectName() AwsInspector2Filter_CodeRepositoryProjectNamePropertyList
	// Experimental.
	CodeRepositoryProjectNameInput() interface{}
	// Experimental.
	CodeRepositoryProviderType() AwsInspector2Filter_CodeRepositoryProviderTypePropertyList
	// Experimental.
	CodeRepositoryProviderTypeInput() interface{}
	// Experimental.
	CodeVulnerabilityDetectorName() AwsInspector2Filter_CodeVulnerabilityDetectorNamePropertyList
	// Experimental.
	CodeVulnerabilityDetectorNameInput() interface{}
	// Experimental.
	CodeVulnerabilityDetectorTags() AwsInspector2Filter_CodeVulnerabilityDetectorTagsPropertyList
	// Experimental.
	CodeVulnerabilityDetectorTagsInput() interface{}
	// Experimental.
	CodeVulnerabilityFilePath() AwsInspector2Filter_CodeVulnerabilityFilePathPropertyList
	// Experimental.
	CodeVulnerabilityFilePathInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	ComponentId() AwsInspector2Filter_ComponentIdPropertyList
	// Experimental.
	ComponentIdInput() interface{}
	// Experimental.
	ComponentType() AwsInspector2Filter_ComponentTypePropertyList
	// Experimental.
	ComponentTypeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Ec2InstanceImageId() AwsInspector2Filter_Ec2InstanceImageIdPropertyList
	// Experimental.
	Ec2InstanceImageIdInput() interface{}
	// Experimental.
	Ec2InstanceSubnetId() AwsInspector2Filter_Ec2InstanceSubnetIdPropertyList
	// Experimental.
	Ec2InstanceSubnetIdInput() interface{}
	// Experimental.
	Ec2InstanceVpcId() AwsInspector2Filter_Ec2InstanceVpcIdPropertyList
	// Experimental.
	Ec2InstanceVpcIdInput() interface{}
	// Experimental.
	EcrImageArchitecture() AwsInspector2Filter_EcrImageArchitecturePropertyList
	// Experimental.
	EcrImageArchitectureInput() interface{}
	// Experimental.
	EcrImageHash() AwsInspector2Filter_EcrImageHashPropertyList
	// Experimental.
	EcrImageHashInput() interface{}
	// Experimental.
	EcrImageInUseCount() AwsInspector2Filter_EcrImageInUseCountPropertyList
	// Experimental.
	EcrImageInUseCountInput() interface{}
	// Experimental.
	EcrImageLastInUseAt() AwsInspector2Filter_EcrImageLastInUseAtPropertyList
	// Experimental.
	EcrImageLastInUseAtInput() interface{}
	// Experimental.
	EcrImagePushedAt() AwsInspector2Filter_EcrImagePushedAtPropertyList
	// Experimental.
	EcrImagePushedAtInput() interface{}
	// Experimental.
	EcrImageRegistry() AwsInspector2Filter_EcrImageRegistryPropertyList
	// Experimental.
	EcrImageRegistryInput() interface{}
	// Experimental.
	EcrImageRepositoryName() AwsInspector2Filter_EcrImageRepositoryNamePropertyList
	// Experimental.
	EcrImageRepositoryNameInput() interface{}
	// Experimental.
	EcrImageTags() AwsInspector2Filter_EcrImageTagsPropertyList
	// Experimental.
	EcrImageTagsInput() interface{}
	// Experimental.
	EpssScore() AwsInspector2Filter_EpssScorePropertyList
	// Experimental.
	EpssScoreInput() interface{}
	// Experimental.
	ExploitAvailable() AwsInspector2Filter_ExploitAvailablePropertyList
	// Experimental.
	ExploitAvailableInput() interface{}
	// Experimental.
	FindingArn() AwsInspector2Filter_FindingArnPropertyList
	// Experimental.
	FindingArnInput() interface{}
	// Experimental.
	FindingStatus() AwsInspector2Filter_FindingStatusPropertyList
	// Experimental.
	FindingStatusInput() interface{}
	// Experimental.
	FindingType() AwsInspector2Filter_FindingTypePropertyList
	// Experimental.
	FindingTypeInput() interface{}
	// Experimental.
	FirstObservedAt() AwsInspector2Filter_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	FixAvailable() AwsInspector2Filter_FixAvailablePropertyList
	// Experimental.
	FixAvailableInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InspectorScore() AwsInspector2Filter_InspectorScorePropertyList
	// Experimental.
	InspectorScoreInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaFunctionExecutionRoleArn() AwsInspector2Filter_LambdaFunctionExecutionRoleArnPropertyList
	// Experimental.
	LambdaFunctionExecutionRoleArnInput() interface{}
	// Experimental.
	LambdaFunctionLastModifiedAt() AwsInspector2Filter_LambdaFunctionLastModifiedAtPropertyList
	// Experimental.
	LambdaFunctionLastModifiedAtInput() interface{}
	// Experimental.
	LambdaFunctionLayers() AwsInspector2Filter_LambdaFunctionLayersPropertyList
	// Experimental.
	LambdaFunctionLayersInput() interface{}
	// Experimental.
	LambdaFunctionName() AwsInspector2Filter_LambdaFunctionNamePropertyList
	// Experimental.
	LambdaFunctionNameInput() interface{}
	// Experimental.
	LambdaFunctionRuntime() AwsInspector2Filter_LambdaFunctionRuntimePropertyList
	// Experimental.
	LambdaFunctionRuntimeInput() interface{}
	// Experimental.
	LastObservedAt() AwsInspector2Filter_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	NetworkProtocol() AwsInspector2Filter_NetworkProtocolPropertyList
	// Experimental.
	NetworkProtocolInput() interface{}
	// Experimental.
	PortRange() AwsInspector2Filter_PortRangePropertyList
	// Experimental.
	PortRangeInput() interface{}
	// Experimental.
	RelatedVulnerabilities() AwsInspector2Filter_RelatedVulnerabilitiesPropertyList
	// Experimental.
	RelatedVulnerabilitiesInput() interface{}
	// Experimental.
	ResourceId() AwsInspector2Filter_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourceTags() AwsInspector2Filter_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() AwsInspector2Filter_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	Severity() AwsInspector2Filter_SeverityPropertyList
	// Experimental.
	SeverityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Title() AwsInspector2Filter_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	UpdatedAt() AwsInspector2Filter_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	VendorSeverity() AwsInspector2Filter_VendorSeverityPropertyList
	// Experimental.
	VendorSeverityInput() interface{}
	// Experimental.
	VulnerabilityId() AwsInspector2Filter_VulnerabilityIdPropertyList
	// Experimental.
	VulnerabilityIdInput() interface{}
	// Experimental.
	VulnerabilitySource() AwsInspector2Filter_VulnerabilitySourcePropertyList
	// Experimental.
	VulnerabilitySourceInput() interface{}
	// Experimental.
	VulnerablePackages() AwsInspector2Filter_VulnerablePackagesPropertyList
	// Experimental.
	VulnerablePackagesInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAwsAccountId(value interface{})
	// Experimental.
	PutCodeRepositoryProjectName(value interface{})
	// Experimental.
	PutCodeRepositoryProviderType(value interface{})
	// Experimental.
	PutCodeVulnerabilityDetectorName(value interface{})
	// Experimental.
	PutCodeVulnerabilityDetectorTags(value interface{})
	// Experimental.
	PutCodeVulnerabilityFilePath(value interface{})
	// Experimental.
	PutComponentId(value interface{})
	// Experimental.
	PutComponentType(value interface{})
	// Experimental.
	PutEc2InstanceImageId(value interface{})
	// Experimental.
	PutEc2InstanceSubnetId(value interface{})
	// Experimental.
	PutEc2InstanceVpcId(value interface{})
	// Experimental.
	PutEcrImageArchitecture(value interface{})
	// Experimental.
	PutEcrImageHash(value interface{})
	// Experimental.
	PutEcrImageInUseCount(value interface{})
	// Experimental.
	PutEcrImageLastInUseAt(value interface{})
	// Experimental.
	PutEcrImagePushedAt(value interface{})
	// Experimental.
	PutEcrImageRegistry(value interface{})
	// Experimental.
	PutEcrImageRepositoryName(value interface{})
	// Experimental.
	PutEcrImageTags(value interface{})
	// Experimental.
	PutEpssScore(value interface{})
	// Experimental.
	PutExploitAvailable(value interface{})
	// Experimental.
	PutFindingArn(value interface{})
	// Experimental.
	PutFindingStatus(value interface{})
	// Experimental.
	PutFindingType(value interface{})
	// Experimental.
	PutFirstObservedAt(value interface{})
	// Experimental.
	PutFixAvailable(value interface{})
	// Experimental.
	PutInspectorScore(value interface{})
	// Experimental.
	PutLambdaFunctionExecutionRoleArn(value interface{})
	// Experimental.
	PutLambdaFunctionLastModifiedAt(value interface{})
	// Experimental.
	PutLambdaFunctionLayers(value interface{})
	// Experimental.
	PutLambdaFunctionName(value interface{})
	// Experimental.
	PutLambdaFunctionRuntime(value interface{})
	// Experimental.
	PutLastObservedAt(value interface{})
	// Experimental.
	PutNetworkProtocol(value interface{})
	// Experimental.
	PutPortRange(value interface{})
	// Experimental.
	PutRelatedVulnerabilities(value interface{})
	// Experimental.
	PutResourceId(value interface{})
	// Experimental.
	PutResourceTags(value interface{})
	// Experimental.
	PutResourceType(value interface{})
	// Experimental.
	PutSeverity(value interface{})
	// Experimental.
	PutTitle(value interface{})
	// Experimental.
	PutUpdatedAt(value interface{})
	// Experimental.
	PutVendorSeverity(value interface{})
	// Experimental.
	PutVulnerabilityId(value interface{})
	// Experimental.
	PutVulnerabilitySource(value interface{})
	// Experimental.
	PutVulnerablePackages(value interface{})
	// Experimental.
	ResetAwsAccountId()
	// Experimental.
	ResetCodeRepositoryProjectName()
	// Experimental.
	ResetCodeRepositoryProviderType()
	// Experimental.
	ResetCodeVulnerabilityDetectorName()
	// Experimental.
	ResetCodeVulnerabilityDetectorTags()
	// Experimental.
	ResetCodeVulnerabilityFilePath()
	// Experimental.
	ResetComponentId()
	// Experimental.
	ResetComponentType()
	// Experimental.
	ResetEc2InstanceImageId()
	// Experimental.
	ResetEc2InstanceSubnetId()
	// Experimental.
	ResetEc2InstanceVpcId()
	// Experimental.
	ResetEcrImageArchitecture()
	// Experimental.
	ResetEcrImageHash()
	// Experimental.
	ResetEcrImageInUseCount()
	// Experimental.
	ResetEcrImageLastInUseAt()
	// Experimental.
	ResetEcrImagePushedAt()
	// Experimental.
	ResetEcrImageRegistry()
	// Experimental.
	ResetEcrImageRepositoryName()
	// Experimental.
	ResetEcrImageTags()
	// Experimental.
	ResetEpssScore()
	// Experimental.
	ResetExploitAvailable()
	// Experimental.
	ResetFindingArn()
	// Experimental.
	ResetFindingStatus()
	// Experimental.
	ResetFindingType()
	// Experimental.
	ResetFirstObservedAt()
	// Experimental.
	ResetFixAvailable()
	// Experimental.
	ResetInspectorScore()
	// Experimental.
	ResetLambdaFunctionExecutionRoleArn()
	// Experimental.
	ResetLambdaFunctionLastModifiedAt()
	// Experimental.
	ResetLambdaFunctionLayers()
	// Experimental.
	ResetLambdaFunctionName()
	// Experimental.
	ResetLambdaFunctionRuntime()
	// Experimental.
	ResetLastObservedAt()
	// Experimental.
	ResetNetworkProtocol()
	// Experimental.
	ResetPortRange()
	// Experimental.
	ResetRelatedVulnerabilities()
	// Experimental.
	ResetResourceId()
	// Experimental.
	ResetResourceTags()
	// Experimental.
	ResetResourceType()
	// Experimental.
	ResetSeverity()
	// Experimental.
	ResetTitle()
	// Experimental.
	ResetUpdatedAt()
	// Experimental.
	ResetVendorSeverity()
	// Experimental.
	ResetVulnerabilityId()
	// Experimental.
	ResetVulnerabilitySource()
	// Experimental.
	ResetVulnerablePackages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsInspector2Filter_FilterCriteriaPropertyOutputReference
type jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) AwsAccountId() AwsInspector2Filter_AwsAccountIdPropertyList {
	var returns AwsInspector2Filter_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeRepositoryProjectName() AwsInspector2Filter_CodeRepositoryProjectNamePropertyList {
	var returns AwsInspector2Filter_CodeRepositoryProjectNamePropertyList
	_jsii_.Get(
		j,
		"codeRepositoryProjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeRepositoryProjectNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeRepositoryProviderType() AwsInspector2Filter_CodeRepositoryProviderTypePropertyList {
	var returns AwsInspector2Filter_CodeRepositoryProviderTypePropertyList
	_jsii_.Get(
		j,
		"codeRepositoryProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeRepositoryProviderTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorName() AwsInspector2Filter_CodeVulnerabilityDetectorNamePropertyList {
	var returns AwsInspector2Filter_CodeVulnerabilityDetectorNamePropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorTags() AwsInspector2Filter_CodeVulnerabilityDetectorTagsPropertyList {
	var returns AwsInspector2Filter_CodeVulnerabilityDetectorTagsPropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityFilePath() AwsInspector2Filter_CodeVulnerabilityFilePathPropertyList {
	var returns AwsInspector2Filter_CodeVulnerabilityFilePathPropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityFilePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ComponentId() AwsInspector2Filter_ComponentIdPropertyList {
	var returns AwsInspector2Filter_ComponentIdPropertyList
	_jsii_.Get(
		j,
		"componentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ComponentIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ComponentType() AwsInspector2Filter_ComponentTypePropertyList {
	var returns AwsInspector2Filter_ComponentTypePropertyList
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ComponentTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Ec2InstanceImageId() AwsInspector2Filter_Ec2InstanceImageIdPropertyList {
	var returns AwsInspector2Filter_Ec2InstanceImageIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Ec2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Ec2InstanceSubnetId() AwsInspector2Filter_Ec2InstanceSubnetIdPropertyList {
	var returns AwsInspector2Filter_Ec2InstanceSubnetIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Ec2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Ec2InstanceVpcId() AwsInspector2Filter_Ec2InstanceVpcIdPropertyList {
	var returns AwsInspector2Filter_Ec2InstanceVpcIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Ec2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageArchitecture() AwsInspector2Filter_EcrImageArchitecturePropertyList {
	var returns AwsInspector2Filter_EcrImageArchitecturePropertyList
	_jsii_.Get(
		j,
		"ecrImageArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageArchitectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageHash() AwsInspector2Filter_EcrImageHashPropertyList {
	var returns AwsInspector2Filter_EcrImageHashPropertyList
	_jsii_.Get(
		j,
		"ecrImageHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageInUseCount() AwsInspector2Filter_EcrImageInUseCountPropertyList {
	var returns AwsInspector2Filter_EcrImageInUseCountPropertyList
	_jsii_.Get(
		j,
		"ecrImageInUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageInUseCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageInUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageLastInUseAt() AwsInspector2Filter_EcrImageLastInUseAtPropertyList {
	var returns AwsInspector2Filter_EcrImageLastInUseAtPropertyList
	_jsii_.Get(
		j,
		"ecrImageLastInUseAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageLastInUseAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageLastInUseAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImagePushedAt() AwsInspector2Filter_EcrImagePushedAtPropertyList {
	var returns AwsInspector2Filter_EcrImagePushedAtPropertyList
	_jsii_.Get(
		j,
		"ecrImagePushedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImagePushedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImagePushedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageRegistry() AwsInspector2Filter_EcrImageRegistryPropertyList {
	var returns AwsInspector2Filter_EcrImageRegistryPropertyList
	_jsii_.Get(
		j,
		"ecrImageRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageRepositoryName() AwsInspector2Filter_EcrImageRepositoryNamePropertyList {
	var returns AwsInspector2Filter_EcrImageRepositoryNamePropertyList
	_jsii_.Get(
		j,
		"ecrImageRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageRepositoryNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRepositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageTags() AwsInspector2Filter_EcrImageTagsPropertyList {
	var returns AwsInspector2Filter_EcrImageTagsPropertyList
	_jsii_.Get(
		j,
		"ecrImageTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EcrImageTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EpssScore() AwsInspector2Filter_EpssScorePropertyList {
	var returns AwsInspector2Filter_EpssScorePropertyList
	_jsii_.Get(
		j,
		"epssScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) EpssScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epssScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ExploitAvailable() AwsInspector2Filter_ExploitAvailablePropertyList {
	var returns AwsInspector2Filter_ExploitAvailablePropertyList
	_jsii_.Get(
		j,
		"exploitAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ExploitAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exploitAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FindingArn() AwsInspector2Filter_FindingArnPropertyList {
	var returns AwsInspector2Filter_FindingArnPropertyList
	_jsii_.Get(
		j,
		"findingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FindingArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FindingStatus() AwsInspector2Filter_FindingStatusPropertyList {
	var returns AwsInspector2Filter_FindingStatusPropertyList
	_jsii_.Get(
		j,
		"findingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FindingStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FindingType() AwsInspector2Filter_FindingTypePropertyList {
	var returns AwsInspector2Filter_FindingTypePropertyList
	_jsii_.Get(
		j,
		"findingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FindingTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FirstObservedAt() AwsInspector2Filter_FirstObservedAtPropertyList {
	var returns AwsInspector2Filter_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FixAvailable() AwsInspector2Filter_FixAvailablePropertyList {
	var returns AwsInspector2Filter_FixAvailablePropertyList
	_jsii_.Get(
		j,
		"fixAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) FixAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) InspectorScore() AwsInspector2Filter_InspectorScorePropertyList {
	var returns AwsInspector2Filter_InspectorScorePropertyList
	_jsii_.Get(
		j,
		"inspectorScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) InspectorScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectorScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionExecutionRoleArn() AwsInspector2Filter_LambdaFunctionExecutionRoleArnPropertyList {
	var returns AwsInspector2Filter_LambdaFunctionExecutionRoleArnPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionExecutionRoleArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionLastModifiedAt() AwsInspector2Filter_LambdaFunctionLastModifiedAtPropertyList {
	var returns AwsInspector2Filter_LambdaFunctionLastModifiedAtPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionLastModifiedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionLayers() AwsInspector2Filter_LambdaFunctionLayersPropertyList {
	var returns AwsInspector2Filter_LambdaFunctionLayersPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionLayers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionLayersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLayersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionName() AwsInspector2Filter_LambdaFunctionNamePropertyList {
	var returns AwsInspector2Filter_LambdaFunctionNamePropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionRuntime() AwsInspector2Filter_LambdaFunctionRuntimePropertyList {
	var returns AwsInspector2Filter_LambdaFunctionRuntimePropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LambdaFunctionRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LastObservedAt() AwsInspector2Filter_LastObservedAtPropertyList {
	var returns AwsInspector2Filter_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) NetworkProtocol() AwsInspector2Filter_NetworkProtocolPropertyList {
	var returns AwsInspector2Filter_NetworkProtocolPropertyList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PortRange() AwsInspector2Filter_PortRangePropertyList {
	var returns AwsInspector2Filter_PortRangePropertyList
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) RelatedVulnerabilities() AwsInspector2Filter_RelatedVulnerabilitiesPropertyList {
	var returns AwsInspector2Filter_RelatedVulnerabilitiesPropertyList
	_jsii_.Get(
		j,
		"relatedVulnerabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) RelatedVulnerabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedVulnerabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResourceId() AwsInspector2Filter_ResourceIdPropertyList {
	var returns AwsInspector2Filter_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResourceTags() AwsInspector2Filter_ResourceTagsPropertyList {
	var returns AwsInspector2Filter_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResourceType() AwsInspector2Filter_ResourceTypePropertyList {
	var returns AwsInspector2Filter_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Severity() AwsInspector2Filter_SeverityPropertyList {
	var returns AwsInspector2Filter_SeverityPropertyList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Title() AwsInspector2Filter_TitlePropertyList {
	var returns AwsInspector2Filter_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) UpdatedAt() AwsInspector2Filter_UpdatedAtPropertyList {
	var returns AwsInspector2Filter_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VendorSeverity() AwsInspector2Filter_VendorSeverityPropertyList {
	var returns AwsInspector2Filter_VendorSeverityPropertyList
	_jsii_.Get(
		j,
		"vendorSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VendorSeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vendorSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VulnerabilityId() AwsInspector2Filter_VulnerabilityIdPropertyList {
	var returns AwsInspector2Filter_VulnerabilityIdPropertyList
	_jsii_.Get(
		j,
		"vulnerabilityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VulnerabilityIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VulnerabilitySource() AwsInspector2Filter_VulnerabilitySourcePropertyList {
	var returns AwsInspector2Filter_VulnerabilitySourcePropertyList
	_jsii_.Get(
		j,
		"vulnerabilitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VulnerabilitySourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilitySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VulnerablePackages() AwsInspector2Filter_VulnerablePackagesPropertyList {
	var returns AwsInspector2Filter_VulnerablePackagesPropertyList
	_jsii_.Get(
		j,
		"vulnerablePackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) VulnerablePackagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerablePackagesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsInspector2Filter_FilterCriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsInspector2Filter_FilterCriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsInspector2Filter_FilterCriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsInspector2Filter.FilterCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsInspector2Filter_FilterCriteriaPropertyOutputReference_Override(a AwsInspector2Filter_FilterCriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsInspector2Filter.FilterCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := a.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutCodeRepositoryProjectName(value interface{}) {
	if err := a.validatePutCodeRepositoryProjectNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepositoryProjectName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutCodeRepositoryProviderType(value interface{}) {
	if err := a.validatePutCodeRepositoryProviderTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepositoryProviderType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityDetectorName(value interface{}) {
	if err := a.validatePutCodeVulnerabilityDetectorNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeVulnerabilityDetectorName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityDetectorTags(value interface{}) {
	if err := a.validatePutCodeVulnerabilityDetectorTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeVulnerabilityDetectorTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityFilePath(value interface{}) {
	if err := a.validatePutCodeVulnerabilityFilePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeVulnerabilityFilePath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutComponentId(value interface{}) {
	if err := a.validatePutComponentIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComponentId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutComponentType(value interface{}) {
	if err := a.validatePutComponentTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComponentType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEc2InstanceImageId(value interface{}) {
	if err := a.validatePutEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEc2InstanceSubnetId(value interface{}) {
	if err := a.validatePutEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEc2InstanceVpcId(value interface{}) {
	if err := a.validatePutEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImageArchitecture(value interface{}) {
	if err := a.validatePutEcrImageArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageArchitecture",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImageHash(value interface{}) {
	if err := a.validatePutEcrImageHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageHash",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImageInUseCount(value interface{}) {
	if err := a.validatePutEcrImageInUseCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageInUseCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImageLastInUseAt(value interface{}) {
	if err := a.validatePutEcrImageLastInUseAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageLastInUseAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImagePushedAt(value interface{}) {
	if err := a.validatePutEcrImagePushedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImagePushedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImageRegistry(value interface{}) {
	if err := a.validatePutEcrImageRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageRegistry",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImageRepositoryName(value interface{}) {
	if err := a.validatePutEcrImageRepositoryNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageRepositoryName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEcrImageTags(value interface{}) {
	if err := a.validatePutEcrImageTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutEpssScore(value interface{}) {
	if err := a.validatePutEpssScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEpssScore",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutExploitAvailable(value interface{}) {
	if err := a.validatePutExploitAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExploitAvailable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutFindingArn(value interface{}) {
	if err := a.validatePutFindingArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutFindingStatus(value interface{}) {
	if err := a.validatePutFindingStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutFindingType(value interface{}) {
	if err := a.validatePutFindingTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := a.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutFixAvailable(value interface{}) {
	if err := a.validatePutFixAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFixAvailable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutInspectorScore(value interface{}) {
	if err := a.validatePutInspectorScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInspectorScore",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionExecutionRoleArn(value interface{}) {
	if err := a.validatePutLambdaFunctionExecutionRoleArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionExecutionRoleArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionLastModifiedAt(value interface{}) {
	if err := a.validatePutLambdaFunctionLastModifiedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionLastModifiedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionLayers(value interface{}) {
	if err := a.validatePutLambdaFunctionLayersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionLayers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionName(value interface{}) {
	if err := a.validatePutLambdaFunctionNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionRuntime(value interface{}) {
	if err := a.validatePutLambdaFunctionRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := a.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutNetworkProtocol(value interface{}) {
	if err := a.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutPortRange(value interface{}) {
	if err := a.validatePutPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPortRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutRelatedVulnerabilities(value interface{}) {
	if err := a.validatePutRelatedVulnerabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedVulnerabilities",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutResourceId(value interface{}) {
	if err := a.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := a.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutResourceType(value interface{}) {
	if err := a.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutSeverity(value interface{}) {
	if err := a.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutTitle(value interface{}) {
	if err := a.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTitle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := a.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutVendorSeverity(value interface{}) {
	if err := a.validatePutVendorSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVendorSeverity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutVulnerabilityId(value interface{}) {
	if err := a.validatePutVulnerabilityIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVulnerabilityId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutVulnerabilitySource(value interface{}) {
	if err := a.validatePutVulnerabilitySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVulnerabilitySource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) PutVulnerablePackages(value interface{}) {
	if err := a.validatePutVulnerablePackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVulnerablePackages",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetCodeRepositoryProjectName() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepositoryProjectName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetCodeRepositoryProviderType() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepositoryProviderType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityDetectorName() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeVulnerabilityDetectorName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityDetectorTags() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeVulnerabilityDetectorTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeVulnerabilityFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetComponentId() {
	_jsii_.InvokeVoid(
		a,
		"resetComponentId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetComponentType() {
	_jsii_.InvokeVoid(
		a,
		"resetComponentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2InstanceImageId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImageArchitecture() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageArchitecture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImageHash() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageHash",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImageInUseCount() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageInUseCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImageLastInUseAt() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageLastInUseAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImagePushedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImagePushedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImageRegistry() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageRegistry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImageRepositoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageRepositoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEcrImageTags() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetEpssScore() {
	_jsii_.InvokeVoid(
		a,
		"resetEpssScore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetExploitAvailable() {
	_jsii_.InvokeVoid(
		a,
		"resetExploitAvailable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetFindingArn() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetFindingStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetFindingType() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetFixAvailable() {
	_jsii_.InvokeVoid(
		a,
		"resetFixAvailable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetInspectorScore() {
	_jsii_.InvokeVoid(
		a,
		"resetInspectorScore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionExecutionRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionExecutionRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionLastModifiedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionLastModifiedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionLayers() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionLayers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionName() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		a,
		"resetPortRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetRelatedVulnerabilities() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedVulnerabilities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		a,
		"resetSeverity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetVendorSeverity() {
	_jsii_.InvokeVoid(
		a,
		"resetVendorSeverity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetVulnerabilityId() {
	_jsii_.InvokeVoid(
		a,
		"resetVulnerabilityId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetVulnerabilitySource() {
	_jsii_.InvokeVoid(
		a,
		"resetVulnerabilitySource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ResetVulnerablePackages() {
	_jsii_.InvokeVoid(
		a,
		"resetVulnerablePackages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2Filter_FilterCriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

