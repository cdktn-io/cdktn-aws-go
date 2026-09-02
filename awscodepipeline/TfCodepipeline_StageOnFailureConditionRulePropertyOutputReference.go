package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference interface {
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
	RuleTypeId() TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdPropertyOutputReference
	// Experimental.
	RuleTypeIdInput() *TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdProperty
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
	PutRuleTypeId(value *TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdProperty)
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

// The jsii proxy struct for TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference
type jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) Commands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) CommandsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) Configuration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) InputArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) InputArtifactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) RuleTypeId() TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdPropertyOutputReference {
	var returns TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdPropertyOutputReference
	_jsii_.Get(
		j,
		"ruleTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) RuleTypeIdInput() *TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdProperty {
	var returns *TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdProperty
	_jsii_.Get(
		j,
		"ruleTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) TimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCodepipeline_StageOnFailureConditionRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCodepipeline_StageOnFailureConditionRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.StageOnFailureConditionRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCodepipeline_StageOnFailureConditionRulePropertyOutputReference_Override(t TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.StageOnFailureConditionRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetCommands(val *[]*string) {
	if err := j.validateSetCommandsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commands",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetConfiguration(val *map[string]*string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetInputArtifacts(val *[]*string) {
	if err := j.validateSetInputArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputArtifacts",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference)SetTimeoutInMinutes(val *float64) {
	if err := j.validateSetTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) PutRuleTypeId(value *TfCodepipeline_StageOnFailureConditionRuleRuleTypeIdProperty) {
	if err := t.validatePutRuleTypeIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRuleTypeId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ResetCommands() {
	_jsii_.InvokeVoid(
		t,
		"resetCommands",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ResetConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ResetInputArtifacts() {
	_jsii_.InvokeVoid(
		t,
		"resetInputArtifacts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ResetTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutInMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCodepipeline_StageOnFailureConditionRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

