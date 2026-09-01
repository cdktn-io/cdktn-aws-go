package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference interface {
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
	InternalValue() *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty
	// Experimental.
	SetInternalValue(val *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty)
	// Experimental.
	TcpEstablishedTimeout() *float64
	// Experimental.
	SetTcpEstablishedTimeout(val *float64)
	// Experimental.
	TcpEstablishedTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UdpStreamTimeout() *float64
	// Experimental.
	SetUdpStreamTimeout(val *float64)
	// Experimental.
	UdpStreamTimeoutInput() *float64
	// Experimental.
	UdpTimeout() *float64
	// Experimental.
	SetUdpTimeout(val *float64)
	// Experimental.
	UdpTimeoutInput() *float64
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
	ResetTcpEstablishedTimeout()
	// Experimental.
	ResetUdpStreamTimeout()
	// Experimental.
	ResetUdpTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference
type jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) InternalValue() *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty {
	var returns *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) TcpEstablishedTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) TcpEstablishedTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) UdpStreamTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpStreamTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) UdpStreamTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpStreamTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) UdpTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) UdpTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpTimeoutInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.ConnectionTrackingSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference_Override(a AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.ConnectionTrackingSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetInternalValue(val *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetTcpEstablishedTimeout(val *float64) {
	if err := j.validateSetTcpEstablishedTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpEstablishedTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetUdpStreamTimeout(val *float64) {
	if err := j.validateSetUdpStreamTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpStreamTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference)SetUdpTimeout(val *float64) {
	if err := j.validateSetUdpTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpTimeout",
		val,
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) ResetTcpEstablishedTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpEstablishedTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) ResetUdpStreamTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetUdpStreamTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) ResetUdpTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetUdpTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

