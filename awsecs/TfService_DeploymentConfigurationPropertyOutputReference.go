package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfService_DeploymentConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BakeTimeInMinutes() *string
	// Experimental.
	SetBakeTimeInMinutes(val *string)
	// Experimental.
	BakeTimeInMinutesInput() *string
	// Experimental.
	CanaryConfiguration() TfService_CanaryConfigurationPropertyOutputReference
	// Experimental.
	CanaryConfigurationInput() *TfService_CanaryConfigurationProperty
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfService_DeploymentConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfService_DeploymentConfigurationProperty)
	// Experimental.
	LifecycleHook() TfService_LifecycleHookPropertyList
	// Experimental.
	LifecycleHookInput() interface{}
	// Experimental.
	LinearConfiguration() TfService_LinearConfigurationPropertyOutputReference
	// Experimental.
	LinearConfigurationInput() *TfService_LinearConfigurationProperty
	// Experimental.
	Strategy() *string
	// Experimental.
	SetStrategy(val *string)
	// Experimental.
	StrategyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutCanaryConfiguration(value *TfService_CanaryConfigurationProperty)
	// Experimental.
	PutLifecycleHook(value interface{})
	// Experimental.
	PutLinearConfiguration(value *TfService_LinearConfigurationProperty)
	// Experimental.
	ResetBakeTimeInMinutes()
	// Experimental.
	ResetCanaryConfiguration()
	// Experimental.
	ResetLifecycleHook()
	// Experimental.
	ResetLinearConfiguration()
	// Experimental.
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfService_DeploymentConfigurationPropertyOutputReference
type jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) BakeTimeInMinutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bakeTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) BakeTimeInMinutesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bakeTimeInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) CanaryConfiguration() TfService_CanaryConfigurationPropertyOutputReference {
	var returns TfService_CanaryConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"canaryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) CanaryConfigurationInput() *TfService_CanaryConfigurationProperty {
	var returns *TfService_CanaryConfigurationProperty
	_jsii_.Get(
		j,
		"canaryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) InternalValue() *TfService_DeploymentConfigurationProperty {
	var returns *TfService_DeploymentConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) LifecycleHook() TfService_LifecycleHookPropertyList {
	var returns TfService_LifecycleHookPropertyList
	_jsii_.Get(
		j,
		"lifecycleHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) LifecycleHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) LinearConfiguration() TfService_LinearConfigurationPropertyOutputReference {
	var returns TfService_LinearConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"linearConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) LinearConfigurationInput() *TfService_LinearConfigurationProperty {
	var returns *TfService_LinearConfigurationProperty
	_jsii_.Get(
		j,
		"linearConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfService_DeploymentConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfService_DeploymentConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfService_DeploymentConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfService.DeploymentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfService_DeploymentConfigurationPropertyOutputReference_Override(t TfService_DeploymentConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfService.DeploymentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference)SetBakeTimeInMinutes(val *string) {
	if err := j.validateSetBakeTimeInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bakeTimeInMinutes",
		val,
	)
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference)SetInternalValue(val *TfService_DeploymentConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) PutCanaryConfiguration(value *TfService_CanaryConfigurationProperty) {
	if err := t.validatePutCanaryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCanaryConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) PutLifecycleHook(value interface{}) {
	if err := t.validatePutLifecycleHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLifecycleHook",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) PutLinearConfiguration(value *TfService_LinearConfigurationProperty) {
	if err := t.validatePutLinearConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLinearConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ResetBakeTimeInMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetBakeTimeInMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ResetCanaryConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCanaryConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ResetLifecycleHook() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleHook",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ResetLinearConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLinearConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfService_DeploymentConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

