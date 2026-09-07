package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference interface {
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
	DestinationAddress() *string
	// Experimental.
	SetDestinationAddress(val *string)
	// Experimental.
	DestinationAddressInput() *string
	// Experimental.
	DestinationPortRange() AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangePropertyOutputReference
	// Experimental.
	DestinationPortRangeInput() *AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEc2NetworkInsightsPath_FilterAtDestinationProperty
	// Experimental.
	SetInternalValue(val *AwsEc2NetworkInsightsPath_FilterAtDestinationProperty)
	// Experimental.
	SourceAddress() *string
	// Experimental.
	SetSourceAddress(val *string)
	// Experimental.
	SourceAddressInput() *string
	// Experimental.
	SourcePortRange() AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangePropertyOutputReference
	// Experimental.
	SourcePortRangeInput() *AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty
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
	PutDestinationPortRange(value *AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty)
	// Experimental.
	PutSourcePortRange(value *AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty)
	// Experimental.
	ResetDestinationAddress()
	// Experimental.
	ResetDestinationPortRange()
	// Experimental.
	ResetSourceAddress()
	// Experimental.
	ResetSourcePortRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference
type jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationPortRange() AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangePropertyOutputReference {
	var returns AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationPortRangeInput() *AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty {
	var returns *AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) InternalValue() *AwsEc2NetworkInsightsPath_FilterAtDestinationProperty {
	var returns *AwsEc2NetworkInsightsPath_FilterAtDestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourcePortRange() AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangePropertyOutputReference {
	var returns AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourcePortRangeInput() *AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty {
	var returns *AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsPath.FilterAtDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference_Override(a AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsPath.FilterAtDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetInternalValue(val *AwsEc2NetworkInsightsPath_FilterAtDestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetSourceAddress(val *string) {
	if err := j.validateSetSourceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddress",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) PutDestinationPortRange(value *AwsEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty) {
	if err := a.validatePutDestinationPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestinationPortRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) PutSourcePortRange(value *AwsEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty) {
	if err := a.validatePutSourcePortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourcePortRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetSourceAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

