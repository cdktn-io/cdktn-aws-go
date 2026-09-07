package route53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/route53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/route53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsRegion() *string
	// Experimental.
	SetAwsRegion(val *string)
	// Experimental.
	AwsRegionInput() *string
	// Experimental.
	Bias() *float64
	// Experimental.
	SetBias(val *float64)
	// Experimental.
	BiasInput() *float64
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
	Coordinates() AwsRecord_CoordinatesPropertyList
	// Experimental.
	CoordinatesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsRecord_GeoproximityRoutingPolicyProperty
	// Experimental.
	SetInternalValue(val *AwsRecord_GeoproximityRoutingPolicyProperty)
	// Experimental.
	LocalZoneGroup() *string
	// Experimental.
	SetLocalZoneGroup(val *string)
	// Experimental.
	LocalZoneGroupInput() *string
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
	PutCoordinates(value interface{})
	// Experimental.
	ResetAwsRegion()
	// Experimental.
	ResetBias()
	// Experimental.
	ResetCoordinates()
	// Experimental.
	ResetLocalZoneGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference
type jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) Bias() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) BiasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"biasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) Coordinates() AwsRecord_CoordinatesPropertyList {
	var returns AwsRecord_CoordinatesPropertyList
	_jsii_.Get(
		j,
		"coordinates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) CoordinatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coordinatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) InternalValue() *AwsRecord_GeoproximityRoutingPolicyProperty {
	var returns *AwsRecord_GeoproximityRoutingPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) LocalZoneGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localZoneGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) LocalZoneGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localZoneGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRecord_GeoproximityRoutingPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRecord_GeoproximityRoutingPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.AwsRecord.GeoproximityRoutingPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRecord_GeoproximityRoutingPolicyPropertyOutputReference_Override(a AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.AwsRecord.GeoproximityRoutingPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetBias(val *float64) {
	if err := j.validateSetBiasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bias",
		val,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetInternalValue(val *AwsRecord_GeoproximityRoutingPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetLocalZoneGroup(val *string) {
	if err := j.validateSetLocalZoneGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localZoneGroup",
		val,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) PutCoordinates(value interface{}) {
	if err := a.validatePutCoordinatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCoordinates",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ResetBias() {
	_jsii_.InvokeVoid(
		a,
		"resetBias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ResetCoordinates() {
	_jsii_.InvokeVoid(
		a,
		"resetCoordinates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ResetLocalZoneGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalZoneGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

