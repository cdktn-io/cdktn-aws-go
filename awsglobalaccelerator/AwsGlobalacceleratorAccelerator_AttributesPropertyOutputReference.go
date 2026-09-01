package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglobalaccelerator/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglobalaccelerator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference interface {
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
	FlowLogsEnabled() interface{}
	// Experimental.
	SetFlowLogsEnabled(val interface{})
	// Experimental.
	FlowLogsEnabledInput() interface{}
	// Experimental.
	FlowLogsS3Bucket() *string
	// Experimental.
	SetFlowLogsS3Bucket(val *string)
	// Experimental.
	FlowLogsS3BucketInput() *string
	// Experimental.
	FlowLogsS3Prefix() *string
	// Experimental.
	SetFlowLogsS3Prefix(val *string)
	// Experimental.
	FlowLogsS3PrefixInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsGlobalacceleratorAccelerator_AttributesProperty
	// Experimental.
	SetInternalValue(val *AwsGlobalacceleratorAccelerator_AttributesProperty)
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
	ResetFlowLogsEnabled()
	// Experimental.
	ResetFlowLogsS3Bucket()
	// Experimental.
	ResetFlowLogsS3Prefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference
type jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) FlowLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) FlowLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) FlowLogsS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) FlowLogsS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) FlowLogsS3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) FlowLogsS3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) InternalValue() *AwsGlobalacceleratorAccelerator_AttributesProperty {
	var returns *AwsGlobalacceleratorAccelerator_AttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGlobalacceleratorAccelerator_AttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-global-accelerator.AwsGlobalacceleratorAccelerator.AttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference_Override(a AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-global-accelerator.AwsGlobalacceleratorAccelerator.AttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetFlowLogsEnabled(val interface{}) {
	if err := j.validateSetFlowLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetFlowLogsS3Bucket(val *string) {
	if err := j.validateSetFlowLogsS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsS3Bucket",
		val,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetFlowLogsS3Prefix(val *string) {
	if err := j.validateSetFlowLogsS3PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsS3Prefix",
		val,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetInternalValue(val *AwsGlobalacceleratorAccelerator_AttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) ResetFlowLogsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFlowLogsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) ResetFlowLogsS3Bucket() {
	_jsii_.InvokeVoid(
		a,
		"resetFlowLogsS3Bucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) ResetFlowLogsS3Prefix() {
	_jsii_.InvokeVoid(
		a,
		"resetFlowLogsS3Prefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGlobalacceleratorAccelerator_AttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

