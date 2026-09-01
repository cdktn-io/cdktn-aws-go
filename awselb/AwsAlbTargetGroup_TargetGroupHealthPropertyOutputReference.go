package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference interface {
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
	DnsFailover() AwsAlbTargetGroup_DnsFailoverPropertyOutputReference
	// Experimental.
	DnsFailoverInput() *AwsAlbTargetGroup_DnsFailoverProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsAlbTargetGroup_TargetGroupHealthProperty
	// Experimental.
	SetInternalValue(val *AwsAlbTargetGroup_TargetGroupHealthProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnhealthyStateRouting() AwsAlbTargetGroup_UnhealthyStateRoutingPropertyOutputReference
	// Experimental.
	UnhealthyStateRoutingInput() *AwsAlbTargetGroup_UnhealthyStateRoutingProperty
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
	PutDnsFailover(value *AwsAlbTargetGroup_DnsFailoverProperty)
	// Experimental.
	PutUnhealthyStateRouting(value *AwsAlbTargetGroup_UnhealthyStateRoutingProperty)
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

// The jsii proxy struct for AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference
type jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) DnsFailover() AwsAlbTargetGroup_DnsFailoverPropertyOutputReference {
	var returns AwsAlbTargetGroup_DnsFailoverPropertyOutputReference
	_jsii_.Get(
		j,
		"dnsFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) DnsFailoverInput() *AwsAlbTargetGroup_DnsFailoverProperty {
	var returns *AwsAlbTargetGroup_DnsFailoverProperty
	_jsii_.Get(
		j,
		"dnsFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) InternalValue() *AwsAlbTargetGroup_TargetGroupHealthProperty {
	var returns *AwsAlbTargetGroup_TargetGroupHealthProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) UnhealthyStateRouting() AwsAlbTargetGroup_UnhealthyStateRoutingPropertyOutputReference {
	var returns AwsAlbTargetGroup_UnhealthyStateRoutingPropertyOutputReference
	_jsii_.Get(
		j,
		"unhealthyStateRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) UnhealthyStateRoutingInput() *AwsAlbTargetGroup_UnhealthyStateRoutingProperty {
	var returns *AwsAlbTargetGroup_UnhealthyStateRoutingProperty
	_jsii_.Get(
		j,
		"unhealthyStateRoutingInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAlbTargetGroup_TargetGroupHealthPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbTargetGroup.TargetGroupHealthPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference_Override(a AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbTargetGroup.TargetGroupHealthPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference)SetInternalValue(val *AwsAlbTargetGroup_TargetGroupHealthProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) PutDnsFailover(value *AwsAlbTargetGroup_DnsFailoverProperty) {
	if err := a.validatePutDnsFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDnsFailover",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) PutUnhealthyStateRouting(value *AwsAlbTargetGroup_UnhealthyStateRoutingProperty) {
	if err := a.validatePutUnhealthyStateRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUnhealthyStateRouting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) ResetDnsFailover() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsFailover",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) ResetUnhealthyStateRouting() {
	_jsii_.InvokeVoid(
		a,
		"resetUnhealthyStateRouting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAlbTargetGroup_TargetGroupHealthPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

