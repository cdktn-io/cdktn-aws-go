package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_DeploymentConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoRollbackConfiguration() AwsEndpoint_AutoRollbackConfigurationPropertyOutputReference
	// Experimental.
	AutoRollbackConfigurationInput() *AwsEndpoint_AutoRollbackConfigurationProperty
	// Experimental.
	BlueGreenUpdatePolicy() AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference
	// Experimental.
	BlueGreenUpdatePolicyInput() *AwsEndpoint_BlueGreenUpdatePolicyProperty
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
	InternalValue() *AwsEndpoint_DeploymentConfigProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_DeploymentConfigProperty)
	// Experimental.
	RollingUpdatePolicy() AwsEndpoint_RollingUpdatePolicyPropertyOutputReference
	// Experimental.
	RollingUpdatePolicyInput() *AwsEndpoint_RollingUpdatePolicyProperty
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
	PutAutoRollbackConfiguration(value *AwsEndpoint_AutoRollbackConfigurationProperty)
	// Experimental.
	PutBlueGreenUpdatePolicy(value *AwsEndpoint_BlueGreenUpdatePolicyProperty)
	// Experimental.
	PutRollingUpdatePolicy(value *AwsEndpoint_RollingUpdatePolicyProperty)
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

// The jsii proxy struct for AwsEndpoint_DeploymentConfigPropertyOutputReference
type jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) AutoRollbackConfiguration() AwsEndpoint_AutoRollbackConfigurationPropertyOutputReference {
	var returns AwsEndpoint_AutoRollbackConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) AutoRollbackConfigurationInput() *AwsEndpoint_AutoRollbackConfigurationProperty {
	var returns *AwsEndpoint_AutoRollbackConfigurationProperty
	_jsii_.Get(
		j,
		"autoRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) BlueGreenUpdatePolicy() AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference {
	var returns AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"blueGreenUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) BlueGreenUpdatePolicyInput() *AwsEndpoint_BlueGreenUpdatePolicyProperty {
	var returns *AwsEndpoint_BlueGreenUpdatePolicyProperty
	_jsii_.Get(
		j,
		"blueGreenUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) InternalValue() *AwsEndpoint_DeploymentConfigProperty {
	var returns *AwsEndpoint_DeploymentConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) RollingUpdatePolicy() AwsEndpoint_RollingUpdatePolicyPropertyOutputReference {
	var returns AwsEndpoint_RollingUpdatePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"rollingUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) RollingUpdatePolicyInput() *AwsEndpoint_RollingUpdatePolicyProperty {
	var returns *AwsEndpoint_RollingUpdatePolicyProperty
	_jsii_.Get(
		j,
		"rollingUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_DeploymentConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_DeploymentConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_DeploymentConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpoint.DeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_DeploymentConfigPropertyOutputReference_Override(a AwsEndpoint_DeploymentConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpoint.DeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference)SetInternalValue(val *AwsEndpoint_DeploymentConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) PutAutoRollbackConfiguration(value *AwsEndpoint_AutoRollbackConfigurationProperty) {
	if err := a.validatePutAutoRollbackConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoRollbackConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) PutBlueGreenUpdatePolicy(value *AwsEndpoint_BlueGreenUpdatePolicyProperty) {
	if err := a.validatePutBlueGreenUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlueGreenUpdatePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) PutRollingUpdatePolicy(value *AwsEndpoint_RollingUpdatePolicyProperty) {
	if err := a.validatePutRollingUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRollingUpdatePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) ResetAutoRollbackConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoRollbackConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) ResetBlueGreenUpdatePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetBlueGreenUpdatePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) ResetRollingUpdatePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRollingUpdatePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_DeploymentConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

