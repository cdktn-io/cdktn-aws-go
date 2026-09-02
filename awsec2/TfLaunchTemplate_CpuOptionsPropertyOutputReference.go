package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLaunchTemplate_CpuOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmdSevSnp() *string
	// Experimental.
	SetAmdSevSnp(val *string)
	// Experimental.
	AmdSevSnpInput() *string
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
	CoreCount() *float64
	// Experimental.
	SetCoreCount(val *float64)
	// Experimental.
	CoreCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfLaunchTemplate_CpuOptionsProperty
	// Experimental.
	SetInternalValue(val *TfLaunchTemplate_CpuOptionsProperty)
	// Experimental.
	NestedVirtualization() *string
	// Experimental.
	SetNestedVirtualization(val *string)
	// Experimental.
	NestedVirtualizationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThreadsPerCore() *float64
	// Experimental.
	SetThreadsPerCore(val *float64)
	// Experimental.
	ThreadsPerCoreInput() *float64
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
	ResetAmdSevSnp()
	// Experimental.
	ResetCoreCount()
	// Experimental.
	ResetNestedVirtualization()
	// Experimental.
	ResetThreadsPerCore()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLaunchTemplate_CpuOptionsPropertyOutputReference
type jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) AmdSevSnp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amdSevSnp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) AmdSevSnpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amdSevSnpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) CoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) CoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) InternalValue() *TfLaunchTemplate_CpuOptionsProperty {
	var returns *TfLaunchTemplate_CpuOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) NestedVirtualization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestedVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) NestedVirtualizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestedVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLaunchTemplate_CpuOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLaunchTemplate_CpuOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLaunchTemplate_CpuOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.CpuOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLaunchTemplate_CpuOptionsPropertyOutputReference_Override(t TfLaunchTemplate_CpuOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.CpuOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetAmdSevSnp(val *string) {
	if err := j.validateSetAmdSevSnpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amdSevSnp",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetCoreCount(val *float64) {
	if err := j.validateSetCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreCount",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetInternalValue(val *TfLaunchTemplate_CpuOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetNestedVirtualization(val *string) {
	if err := j.validateSetNestedVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nestedVirtualization",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ResetAmdSevSnp() {
	_jsii_.InvokeVoid(
		t,
		"resetAmdSevSnp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ResetCoreCount() {
	_jsii_.InvokeVoid(
		t,
		"resetCoreCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ResetNestedVirtualization() {
	_jsii_.InvokeVoid(
		t,
		"resetNestedVirtualization",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		t,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_CpuOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

