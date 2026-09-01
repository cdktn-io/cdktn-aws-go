package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodepipeline_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Category() *string
	// Experimental.
	SetCategory(val *string)
	// Experimental.
	CategoryInput() *string
	// Experimental.
	Commands() *[]*string
	// Experimental.
	SetCommands(val *[]*string)
	// Experimental.
	CommandsInput() *[]*string
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
	Configuration() *map[string]*string
	// Experimental.
	SetConfiguration(val *map[string]*string)
	// Experimental.
	ConfigurationInput() *map[string]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InputArtifacts() *[]*string
	// Experimental.
	SetInputArtifacts(val *[]*string)
	// Experimental.
	InputArtifactsInput() *[]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Namespace() *string
	// Experimental.
	SetNamespace(val *string)
	// Experimental.
	NamespaceInput() *string
	// Experimental.
	OutputArtifacts() *[]*string
	// Experimental.
	SetOutputArtifacts(val *[]*string)
	// Experimental.
	OutputArtifactsForComputeAction() AwsCodepipeline_OutputArtifactsForComputeActionPropertyList
	// Experimental.
	OutputArtifactsForComputeActionInput() interface{}
	// Experimental.
	OutputArtifactsInput() *[]*string
	// Experimental.
	OutputVariables() *[]*string
	// Experimental.
	SetOutputVariables(val *[]*string)
	// Experimental.
	OutputVariablesInput() *[]*string
	// Experimental.
	Owner() *string
	// Experimental.
	SetOwner(val *string)
	// Experimental.
	OwnerInput() *string
	// Experimental.
	Provider() *string
	// Experimental.
	SetProvider(val *string)
	// Experimental.
	ProviderInput() *string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	RunOrder() *float64
	// Experimental.
	SetRunOrder(val *float64)
	// Experimental.
	RunOrderInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutInMinutes() *float64
	// Experimental.
	SetTimeoutInMinutes(val *float64)
	// Experimental.
	TimeoutInMinutesInput() *float64
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
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
	PutOutputArtifactsForComputeAction(value interface{})
	// Experimental.
	ResetCommands()
	// Experimental.
	ResetConfiguration()
	// Experimental.
	ResetInputArtifacts()
	// Experimental.
	ResetNamespace()
	// Experimental.
	ResetOutputArtifacts()
	// Experimental.
	ResetOutputArtifactsForComputeAction()
	// Experimental.
	ResetOutputVariables()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRoleArn()
	// Experimental.
	ResetRunOrder()
	// Experimental.
	ResetTimeoutInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCodepipeline_ActionPropertyOutputReference
type jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Commands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) CommandsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Configuration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) InputArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) InputArtifactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) OutputArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) OutputArtifactsForComputeAction() AwsCodepipeline_OutputArtifactsForComputeActionPropertyList {
	var returns AwsCodepipeline_OutputArtifactsForComputeActionPropertyList
	_jsii_.Get(
		j,
		"outputArtifactsForComputeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) OutputArtifactsForComputeActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputArtifactsForComputeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) OutputArtifactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) OutputVariables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) OutputVariablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) RunOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) RunOrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) TimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCodepipeline_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCodepipeline_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCodepipeline_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCodepipeline_ActionPropertyOutputReference_Override(a AwsCodepipeline_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetCategory(val *string) {
	if err := j.validateSetCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetCommands(val *[]*string) {
	if err := j.validateSetCommandsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commands",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetConfiguration(val *map[string]*string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetInputArtifacts(val *[]*string) {
	if err := j.validateSetInputArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputArtifacts",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetOutputArtifacts(val *[]*string) {
	if err := j.validateSetOutputArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputArtifacts",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetOutputVariables(val *[]*string) {
	if err := j.validateSetOutputVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputVariables",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetRunOrder(val *float64) {
	if err := j.validateSetRunOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runOrder",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetTimeoutInMinutes(val *float64) {
	if err := j.validateSetTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) PutOutputArtifactsForComputeAction(value interface{}) {
	if err := a.validatePutOutputArtifactsForComputeActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputArtifactsForComputeAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetCommands() {
	_jsii_.InvokeVoid(
		a,
		"resetCommands",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetInputArtifacts() {
	_jsii_.InvokeVoid(
		a,
		"resetInputArtifacts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetOutputArtifacts() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputArtifacts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetOutputArtifactsForComputeAction() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputArtifactsForComputeAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetOutputVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetRunOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetRunOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ResetTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

