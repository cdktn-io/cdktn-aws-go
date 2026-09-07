package batch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/batch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/batch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference interface {
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
	InternalValue() *AwsSchedulingPolicy_FairSharePolicyProperty
	// Experimental.
	SetInternalValue(val *AwsSchedulingPolicy_FairSharePolicyProperty)
	// Experimental.
	ShareDecaySeconds() *float64
	// Experimental.
	SetShareDecaySeconds(val *float64)
	// Experimental.
	ShareDecaySecondsInput() *float64
	// Experimental.
	ShareDistribution() AwsSchedulingPolicy_ShareDistributionPropertyList
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

// The jsii proxy struct for AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference
type jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComputeReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComputeReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) InternalValue() *AwsSchedulingPolicy_FairSharePolicyProperty {
	var returns *AwsSchedulingPolicy_FairSharePolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDecaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDecaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDistribution() AwsSchedulingPolicy_ShareDistributionPropertyList {
	var returns AwsSchedulingPolicy_ShareDistributionPropertyList
	_jsii_.Get(
		j,
		"shareDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ShareDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSchedulingPolicy_FairSharePolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSchedulingPolicy_FairSharePolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.AwsSchedulingPolicy.FairSharePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSchedulingPolicy_FairSharePolicyPropertyOutputReference_Override(a AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.AwsSchedulingPolicy.FairSharePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetComputeReservation(val *float64) {
	if err := j.validateSetComputeReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeReservation",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetInternalValue(val *AwsSchedulingPolicy_FairSharePolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetShareDecaySeconds(val *float64) {
	if err := j.validateSetShareDecaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDecaySeconds",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) PutShareDistribution(value interface{}) {
	if err := a.validatePutShareDistributionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putShareDistribution",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ResetComputeReservation() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeReservation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ResetShareDecaySeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetShareDecaySeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ResetShareDistribution() {
	_jsii_.InvokeVoid(
		a,
		"resetShareDistribution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSchedulingPolicy_FairSharePolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

