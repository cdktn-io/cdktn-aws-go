package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslambda/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslambda/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityProviderArn() *string
	// Experimental.
	SetCapacityProviderArn(val *string)
	// Experimental.
	CapacityProviderArnInput() *string
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
	ExecutionEnvironmentMemoryGibPerVcpu() *float64
	// Experimental.
	SetExecutionEnvironmentMemoryGibPerVcpu(val *float64)
	// Experimental.
	ExecutionEnvironmentMemoryGibPerVcpuInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFunction_LambdaManagedInstancesCapacityProviderConfigProperty
	// Experimental.
	SetInternalValue(val *TfFunction_LambdaManagedInstancesCapacityProviderConfigProperty)
	// Experimental.
	PerExecutionEnvironmentMaxConcurrency() *float64
	// Experimental.
	SetPerExecutionEnvironmentMaxConcurrency(val *float64)
	// Experimental.
	PerExecutionEnvironmentMaxConcurrencyInput() *float64
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
	ResetExecutionEnvironmentMemoryGibPerVcpu()
	// Experimental.
	ResetPerExecutionEnvironmentMaxConcurrency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference
type jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) CapacityProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) CapacityProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ExecutionEnvironmentMemoryGibPerVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionEnvironmentMemoryGibPerVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ExecutionEnvironmentMemoryGibPerVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionEnvironmentMemoryGibPerVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) InternalValue() *TfFunction_LambdaManagedInstancesCapacityProviderConfigProperty {
	var returns *TfFunction_LambdaManagedInstancesCapacityProviderConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) PerExecutionEnvironmentMaxConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perExecutionEnvironmentMaxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) PerExecutionEnvironmentMaxConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perExecutionEnvironmentMaxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lambda.TfFunction.LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference_Override(t TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.TfFunction.LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetCapacityProviderArn(val *string) {
	if err := j.validateSetCapacityProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityProviderArn",
		val,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetExecutionEnvironmentMemoryGibPerVcpu(val *float64) {
	if err := j.validateSetExecutionEnvironmentMemoryGibPerVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionEnvironmentMemoryGibPerVcpu",
		val,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetInternalValue(val *TfFunction_LambdaManagedInstancesCapacityProviderConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetPerExecutionEnvironmentMaxConcurrency(val *float64) {
	if err := j.validateSetPerExecutionEnvironmentMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perExecutionEnvironmentMaxConcurrency",
		val,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ResetExecutionEnvironmentMemoryGibPerVcpu() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionEnvironmentMemoryGibPerVcpu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ResetPerExecutionEnvironmentMaxConcurrency() {
	_jsii_.InvokeVoid(
		t,
		"resetPerExecutionEnvironmentMaxConcurrency",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFunction_LambdaManagedInstancesCapacityProviderConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

