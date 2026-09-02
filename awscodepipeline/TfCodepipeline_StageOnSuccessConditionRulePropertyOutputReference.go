package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference interface {
	cdktn.ComplexObject
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
	RuleTypeId() TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdPropertyOutputReference
	// Experimental.
	RuleTypeIdInput() *TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdProperty
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
	PutRuleTypeId(value *TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdProperty)
	// Experimental.
	ResetCommands()
	// Experimental.
	ResetConfiguration()
	// Experimental.
	ResetInputArtifacts()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRoleArn()
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

// The jsii proxy struct for TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference
type jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) Commands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) CommandsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) Configuration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) InputArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) InputArtifactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) RuleTypeId() TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdPropertyOutputReference {
	var returns TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdPropertyOutputReference
	_jsii_.Get(
		j,
		"ruleTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) RuleTypeIdInput() *TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdProperty {
	var returns *TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdProperty
	_jsii_.Get(
		j,
		"ruleTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) TimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCodepipeline_StageOnSuccessConditionRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.StageOnSuccessConditionRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference_Override(t TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.StageOnSuccessConditionRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetCommands(val *[]*string) {
	if err := j.validateSetCommandsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commands",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetConfiguration(val *map[string]*string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetInputArtifacts(val *[]*string) {
	if err := j.validateSetInputArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputArtifacts",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference)SetTimeoutInMinutes(val *float64) {
	if err := j.validateSetTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) PutRuleTypeId(value *TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdProperty) {
	if err := t.validatePutRuleTypeIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRuleTypeId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ResetCommands() {
	_jsii_.InvokeVoid(
		t,
		"resetCommands",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ResetConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ResetInputArtifacts() {
	_jsii_.InvokeVoid(
		t,
		"resetInputArtifacts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ResetTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutInMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnSuccessConditionRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

