package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSchedulingPolicy_FairSharePolicyPropertyOutputReference interface {
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
	// Experimental.
	ComputeReservation() *float64
	// Experimental.
	SetComputeReservation(val *float64)
	// Experimental.
	ComputeReservationInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfSchedulingPolicy_FairSharePolicyProperty
	// Experimental.
	SetInternalValue(val *TfSchedulingPolicy_FairSharePolicyProperty)
	// Experimental.
	ShareDecaySeconds() *float64
	// Experimental.
	SetShareDecaySeconds(val *float64)
	// Experimental.
	ShareDecaySecondsInput() *float64
	// Experimental.
	ShareDistribution() TfSchedulingPolicy_ShareDistributionPropertyList
	// Experimental.
	ShareDistributionInput() interface{}
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
	PutShareDistribution(value interface{})
	// Experimental.
	ResetComputeReservation()
	// Experimental.
	ResetShareDecaySeconds()
	// Experimental.
	ResetShareDistribution()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSchedulingPolicy_FairSharePolicyPropertyOutputReference
type jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComputeReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComputeReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) InternalValue() *TfSchedulingPolicy_FairSharePolicyProperty {
	var returns *TfSchedulingPolicy_FairSharePolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDecaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDecaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDistribution() TfSchedulingPolicy_ShareDistributionPropertyList {
	var returns TfSchedulingPolicy_ShareDistributionPropertyList
	_jsii_.Get(
		j,
		"shareDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSchedulingPolicy_FairSharePolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfSchedulingPolicy_FairSharePolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSchedulingPolicy_FairSharePolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.TfSchedulingPolicy.FairSharePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSchedulingPolicy_FairSharePolicyPropertyOutputReference_Override(t TfSchedulingPolicy_FairSharePolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.TfSchedulingPolicy.FairSharePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetComputeReservation(val *float64) {
	if err := j.validateSetComputeReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeReservation",
		val,
	)
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetInternalValue(val *TfSchedulingPolicy_FairSharePolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetShareDecaySeconds(val *float64) {
	if err := j.validateSetShareDecaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDecaySeconds",
		val,
	)
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) PutShareDistribution(value interface{}) {
	if err := t.validatePutShareDistributionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putShareDistribution",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ResetComputeReservation() {
	_jsii_.InvokeVoid(
		t,
		"resetComputeReservation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ResetShareDecaySeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetShareDecaySeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ResetShareDistribution() {
	_jsii_.InvokeVoid(
		t,
		"resetShareDistribution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSchedulingPolicy_FairSharePolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

