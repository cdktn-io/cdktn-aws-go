package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_DeploymentConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoRollbackConfiguration() TfEndpoint_AutoRollbackConfigurationPropertyOutputReference
	// Experimental.
	AutoRollbackConfigurationInput() *TfEndpoint_AutoRollbackConfigurationProperty
	// Experimental.
	BlueGreenUpdatePolicy() TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference
	// Experimental.
	BlueGreenUpdatePolicyInput() *TfEndpoint_BlueGreenUpdatePolicyProperty
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
	InternalValue() *TfEndpoint_DeploymentConfigProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_DeploymentConfigProperty)
	// Experimental.
	RollingUpdatePolicy() TfEndpoint_RollingUpdatePolicyPropertyOutputReference
	// Experimental.
	RollingUpdatePolicyInput() *TfEndpoint_RollingUpdatePolicyProperty
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
	PutAutoRollbackConfiguration(value *TfEndpoint_AutoRollbackConfigurationProperty)
	// Experimental.
	PutBlueGreenUpdatePolicy(value *TfEndpoint_BlueGreenUpdatePolicyProperty)
	// Experimental.
	PutRollingUpdatePolicy(value *TfEndpoint_RollingUpdatePolicyProperty)
	// Experimental.
	ResetAutoRollbackConfiguration()
	// Experimental.
	ResetBlueGreenUpdatePolicy()
	// Experimental.
	ResetRollingUpdatePolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_DeploymentConfigPropertyOutputReference
type jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) AutoRollbackConfiguration() TfEndpoint_AutoRollbackConfigurationPropertyOutputReference {
	var returns TfEndpoint_AutoRollbackConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) AutoRollbackConfigurationInput() *TfEndpoint_AutoRollbackConfigurationProperty {
	var returns *TfEndpoint_AutoRollbackConfigurationProperty
	_jsii_.Get(
		j,
		"autoRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) BlueGreenUpdatePolicy() TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference {
	var returns TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"blueGreenUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) BlueGreenUpdatePolicyInput() *TfEndpoint_BlueGreenUpdatePolicyProperty {
	var returns *TfEndpoint_BlueGreenUpdatePolicyProperty
	_jsii_.Get(
		j,
		"blueGreenUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) InternalValue() *TfEndpoint_DeploymentConfigProperty {
	var returns *TfEndpoint_DeploymentConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) RollingUpdatePolicy() TfEndpoint_RollingUpdatePolicyPropertyOutputReference {
	var returns TfEndpoint_RollingUpdatePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"rollingUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) RollingUpdatePolicyInput() *TfEndpoint_RollingUpdatePolicyProperty {
	var returns *TfEndpoint_RollingUpdatePolicyProperty
	_jsii_.Get(
		j,
		"rollingUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_DeploymentConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_DeploymentConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_DeploymentConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpoint.DeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_DeploymentConfigPropertyOutputReference_Override(t TfEndpoint_DeploymentConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpoint.DeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference)SetInternalValue(val *TfEndpoint_DeploymentConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) PutAutoRollbackConfiguration(value *TfEndpoint_AutoRollbackConfigurationProperty) {
	if err := t.validatePutAutoRollbackConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutoRollbackConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) PutBlueGreenUpdatePolicy(value *TfEndpoint_BlueGreenUpdatePolicyProperty) {
	if err := t.validatePutBlueGreenUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlueGreenUpdatePolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) PutRollingUpdatePolicy(value *TfEndpoint_RollingUpdatePolicyProperty) {
	if err := t.validatePutRollingUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRollingUpdatePolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) ResetAutoRollbackConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoRollbackConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) ResetBlueGreenUpdatePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetBlueGreenUpdatePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) ResetRollingUpdatePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetRollingUpdatePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_DeploymentConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

