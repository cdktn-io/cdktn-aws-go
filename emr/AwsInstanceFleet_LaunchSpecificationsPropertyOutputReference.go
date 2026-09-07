package emr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference interface {
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
	InternalValue() *AwsInstanceFleet_LaunchSpecificationsProperty
	// Experimental.
	SetInternalValue(val *AwsInstanceFleet_LaunchSpecificationsProperty)
	// Experimental.
	OnDemandSpecification() AwsInstanceFleet_OnDemandSpecificationPropertyList
	// Experimental.
	OnDemandSpecificationInput() interface{}
	// Experimental.
	SpotSpecification() AwsInstanceFleet_SpotSpecificationPropertyList
	// Experimental.
	SpotSpecificationInput() interface{}
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
	PutOnDemandSpecification(value interface{})
	// Experimental.
	PutSpotSpecification(value interface{})
	// Experimental.
	ResetOnDemandSpecification()
	// Experimental.
	ResetSpotSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference
type jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) InternalValue() *AwsInstanceFleet_LaunchSpecificationsProperty {
	var returns *AwsInstanceFleet_LaunchSpecificationsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) OnDemandSpecification() AwsInstanceFleet_OnDemandSpecificationPropertyList {
	var returns AwsInstanceFleet_OnDemandSpecificationPropertyList
	_jsii_.Get(
		j,
		"onDemandSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) OnDemandSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) SpotSpecification() AwsInstanceFleet_SpotSpecificationPropertyList {
	var returns AwsInstanceFleet_SpotSpecificationPropertyList
	_jsii_.Get(
		j,
		"spotSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) SpotSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsInstanceFleet_LaunchSpecificationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsInstanceFleet_LaunchSpecificationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsInstanceFleet.LaunchSpecificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsInstanceFleet_LaunchSpecificationsPropertyOutputReference_Override(a AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsInstanceFleet.LaunchSpecificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference)SetInternalValue(val *AwsInstanceFleet_LaunchSpecificationsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) PutOnDemandSpecification(value interface{}) {
	if err := a.validatePutOnDemandSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOnDemandSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) PutSpotSpecification(value interface{}) {
	if err := a.validatePutSpotSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpotSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) ResetOnDemandSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) ResetSpotSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsInstanceFleet_LaunchSpecificationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

