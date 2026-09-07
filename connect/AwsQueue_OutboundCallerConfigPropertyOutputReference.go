package connect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/connect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/connect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQueue_OutboundCallerConfigPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsQueue_OutboundCallerConfigProperty
	// Experimental.
	SetInternalValue(val *AwsQueue_OutboundCallerConfigProperty)
	// Experimental.
	OutboundCallerIdName() *string
	// Experimental.
	SetOutboundCallerIdName(val *string)
	// Experimental.
	OutboundCallerIdNameInput() *string
	// Experimental.
	OutboundCallerIdNumberId() *string
	// Experimental.
	SetOutboundCallerIdNumberId(val *string)
	// Experimental.
	OutboundCallerIdNumberIdInput() *string
	// Experimental.
	OutboundFlowId() *string
	// Experimental.
	SetOutboundFlowId(val *string)
	// Experimental.
	OutboundFlowIdInput() *string
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
	ResetOutboundCallerIdName()
	// Experimental.
	ResetOutboundCallerIdNumberId()
	// Experimental.
	ResetOutboundFlowId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsQueue_OutboundCallerConfigPropertyOutputReference
type jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) InternalValue() *AwsQueue_OutboundCallerConfigProperty {
	var returns *AwsQueue_OutboundCallerConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) OutboundCallerIdName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) OutboundCallerIdNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) OutboundCallerIdNumberId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNumberId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) OutboundCallerIdNumberIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNumberIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) OutboundFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) OutboundFlowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundFlowIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQueue_OutboundCallerConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQueue_OutboundCallerConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQueue_OutboundCallerConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.AwsQueue.OutboundCallerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQueue_OutboundCallerConfigPropertyOutputReference_Override(a AwsQueue_OutboundCallerConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.AwsQueue.OutboundCallerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetInternalValue(val *AwsQueue_OutboundCallerConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetOutboundCallerIdName(val *string) {
	if err := j.validateSetOutboundCallerIdNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallerIdName",
		val,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetOutboundCallerIdNumberId(val *string) {
	if err := j.validateSetOutboundCallerIdNumberIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallerIdNumberId",
		val,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetOutboundFlowId(val *string) {
	if err := j.validateSetOutboundFlowIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundFlowId",
		val,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) ResetOutboundCallerIdName() {
	_jsii_.InvokeVoid(
		a,
		"resetOutboundCallerIdName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) ResetOutboundCallerIdNumberId() {
	_jsii_.InvokeVoid(
		a,
		"resetOutboundCallerIdNumberId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) ResetOutboundFlowId() {
	_jsii_.InvokeVoid(
		a,
		"resetOutboundFlowId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsQueue_OutboundCallerConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

