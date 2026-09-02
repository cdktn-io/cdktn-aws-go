package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTargetGroup_TargetGroupHealthPropertyOutputReference interface {
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
	DnsFailover() TfTargetGroup_DnsFailoverPropertyOutputReference
	// Experimental.
	DnsFailoverInput() *TfTargetGroup_DnsFailoverProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTargetGroup_TargetGroupHealthProperty
	// Experimental.
	SetInternalValue(val *TfTargetGroup_TargetGroupHealthProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnhealthyStateRouting() TfTargetGroup_UnhealthyStateRoutingPropertyOutputReference
	// Experimental.
	UnhealthyStateRoutingInput() *TfTargetGroup_UnhealthyStateRoutingProperty
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
	PutDnsFailover(value *TfTargetGroup_DnsFailoverProperty)
	// Experimental.
	PutUnhealthyStateRouting(value *TfTargetGroup_UnhealthyStateRoutingProperty)
	// Experimental.
	ResetDnsFailover()
	// Experimental.
	ResetUnhealthyStateRouting()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTargetGroup_TargetGroupHealthPropertyOutputReference
type jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) DnsFailover() TfTargetGroup_DnsFailoverPropertyOutputReference {
	var returns TfTargetGroup_DnsFailoverPropertyOutputReference
	_jsii_.Get(
		j,
		"dnsFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) DnsFailoverInput() *TfTargetGroup_DnsFailoverProperty {
	var returns *TfTargetGroup_DnsFailoverProperty
	_jsii_.Get(
		j,
		"dnsFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) InternalValue() *TfTargetGroup_TargetGroupHealthProperty {
	var returns *TfTargetGroup_TargetGroupHealthProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) UnhealthyStateRouting() TfTargetGroup_UnhealthyStateRoutingPropertyOutputReference {
	var returns TfTargetGroup_UnhealthyStateRoutingPropertyOutputReference
	_jsii_.Get(
		j,
		"unhealthyStateRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) UnhealthyStateRoutingInput() *TfTargetGroup_UnhealthyStateRoutingProperty {
	var returns *TfTargetGroup_UnhealthyStateRoutingProperty
	_jsii_.Get(
		j,
		"unhealthyStateRoutingInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTargetGroup_TargetGroupHealthPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTargetGroup_TargetGroupHealthPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTargetGroup_TargetGroupHealthPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfTargetGroup.TargetGroupHealthPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTargetGroup_TargetGroupHealthPropertyOutputReference_Override(t TfTargetGroup_TargetGroupHealthPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfTargetGroup.TargetGroupHealthPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference)SetInternalValue(val *TfTargetGroup_TargetGroupHealthProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) PutDnsFailover(value *TfTargetGroup_DnsFailoverProperty) {
	if err := t.validatePutDnsFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDnsFailover",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) PutUnhealthyStateRouting(value *TfTargetGroup_UnhealthyStateRoutingProperty) {
	if err := t.validatePutUnhealthyStateRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUnhealthyStateRouting",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) ResetDnsFailover() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsFailover",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) ResetUnhealthyStateRouting() {
	_jsii_.InvokeVoid(
		t,
		"resetUnhealthyStateRouting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTargetGroup_TargetGroupHealthPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

