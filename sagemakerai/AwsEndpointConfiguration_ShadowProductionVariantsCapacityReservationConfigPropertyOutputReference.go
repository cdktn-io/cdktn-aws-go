package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityReservationPreference() *string
	// Experimental.
	SetCapacityReservationPreference(val *string)
	// Experimental.
	CapacityReservationPreferenceInput() *string
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
	InternalValue() *AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty
	// Experimental.
	SetInternalValue(val *AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty)
	// Experimental.
	MlReservationArn() *string
	// Experimental.
	SetMlReservationArn(val *string)
	// Experimental.
	MlReservationArnInput() *string
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
	ResetCapacityReservationPreference()
	// Experimental.
	ResetMlReservationArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference
type jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) CapacityReservationPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) CapacityReservationPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) InternalValue() *AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty {
	var returns *AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) MlReservationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlReservationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) MlReservationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlReservationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpointConfiguration.ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference_Override(a AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpointConfiguration.ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetCapacityReservationPreference(val *string) {
	if err := j.validateSetCapacityReservationPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationPreference",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetInternalValue(val *AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetMlReservationArn(val *string) {
	if err := j.validateSetMlReservationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlReservationArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ResetCapacityReservationPreference() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityReservationPreference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ResetMlReservationArn() {
	_jsii_.InvokeVoid(
		a,
		"resetMlReservationArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

