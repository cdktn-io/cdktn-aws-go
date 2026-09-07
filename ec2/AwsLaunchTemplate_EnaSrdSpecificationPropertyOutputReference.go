package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference interface {
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
	EnaSrdEnabled() interface{}
	// Experimental.
	SetEnaSrdEnabled(val interface{})
	// Experimental.
	EnaSrdEnabledInput() interface{}
	// Experimental.
	EnaSrdUdpSpecification() AwsLaunchTemplate_EnaSrdUdpSpecificationPropertyOutputReference
	// Experimental.
	EnaSrdUdpSpecificationInput() *AwsLaunchTemplate_EnaSrdUdpSpecificationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLaunchTemplate_EnaSrdSpecificationProperty
	// Experimental.
	SetInternalValue(val *AwsLaunchTemplate_EnaSrdSpecificationProperty)
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
	PutEnaSrdUdpSpecification(value *AwsLaunchTemplate_EnaSrdUdpSpecificationProperty)
	// Experimental.
	ResetEnaSrdEnabled()
	// Experimental.
	ResetEnaSrdUdpSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference
type jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) EnaSrdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSrdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) EnaSrdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSrdEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) EnaSrdUdpSpecification() AwsLaunchTemplate_EnaSrdUdpSpecificationPropertyOutputReference {
	var returns AwsLaunchTemplate_EnaSrdUdpSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"enaSrdUdpSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) EnaSrdUdpSpecificationInput() *AwsLaunchTemplate_EnaSrdUdpSpecificationProperty {
	var returns *AwsLaunchTemplate_EnaSrdUdpSpecificationProperty
	_jsii_.Get(
		j,
		"enaSrdUdpSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) InternalValue() *AwsLaunchTemplate_EnaSrdSpecificationProperty {
	var returns *AwsLaunchTemplate_EnaSrdSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.EnaSrdSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference_Override(a AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.EnaSrdSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference)SetEnaSrdEnabled(val interface{}) {
	if err := j.validateSetEnaSrdEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enaSrdEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference)SetInternalValue(val *AwsLaunchTemplate_EnaSrdSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) PutEnaSrdUdpSpecification(value *AwsLaunchTemplate_EnaSrdUdpSpecificationProperty) {
	if err := a.validatePutEnaSrdUdpSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnaSrdUdpSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) ResetEnaSrdEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnaSrdEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) ResetEnaSrdUdpSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetEnaSrdUdpSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

