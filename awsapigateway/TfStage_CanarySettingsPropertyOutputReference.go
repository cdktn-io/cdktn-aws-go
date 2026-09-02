package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapigateway/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStage_CanarySettingsPropertyOutputReference interface {
	cdktn.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DeploymentId() *string
	// Experimental.
	SetDeploymentId(val *string)
	// Experimental.
	DeploymentIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfStage_CanarySettingsProperty
	// Experimental.
	SetInternalValue(val *TfStage_CanarySettingsProperty)
	// Experimental.
	PercentTraffic() *float64
	// Experimental.
	SetPercentTraffic(val *float64)
	// Experimental.
	PercentTrafficInput() *float64
	// Experimental.
	StageVariableOverrides() *map[string]*string
	// Experimental.
	SetStageVariableOverrides(val *map[string]*string)
	// Experimental.
	StageVariableOverridesInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseStageCache() interface{}
	// Experimental.
	SetUseStageCache(val interface{})
	// Experimental.
	UseStageCacheInput() interface{}
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
	ResetPercentTraffic()
	// Experimental.
	ResetStageVariableOverrides()
	// Experimental.
	ResetUseStageCache()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfStage_CanarySettingsPropertyOutputReference
type jsiiProxy_TfStage_CanarySettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) DeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) InternalValue() *TfStage_CanarySettingsProperty {
	var returns *TfStage_CanarySettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) PercentTraffic() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) PercentTrafficInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) StageVariableOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariableOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) StageVariableOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariableOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) UseStageCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStageCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) UseStageCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStageCacheInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStage_CanarySettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStage_CanarySettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStage_CanarySettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStage_CanarySettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-api-gateway.TfStage.CanarySettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStage_CanarySettingsPropertyOutputReference_Override(t TfStage_CanarySettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-api-gateway.TfStage.CanarySettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetDeploymentId(val *string) {
	if err := j.validateSetDeploymentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetInternalValue(val *TfStage_CanarySettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetPercentTraffic(val *float64) {
	if err := j.validateSetPercentTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentTraffic",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetStageVariableOverrides(val *map[string]*string) {
	if err := j.validateSetStageVariableOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stageVariableOverrides",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference)SetUseStageCache(val interface{}) {
	if err := j.validateSetUseStageCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useStageCache",
		val,
	)
}

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) ResetPercentTraffic() {
	_jsii_.InvokeVoid(
		t,
		"resetPercentTraffic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) ResetStageVariableOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetStageVariableOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) ResetUseStageCache() {
	_jsii_.InvokeVoid(
		t,
		"resetUseStageCache",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStage_CanarySettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

