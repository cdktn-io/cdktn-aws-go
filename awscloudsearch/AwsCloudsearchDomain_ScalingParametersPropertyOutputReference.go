package awscloudsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudsearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudsearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudsearchDomain_ScalingParametersPropertyOutputReference interface {
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
	DesiredInstanceType() *string
	// Experimental.
	SetDesiredInstanceType(val *string)
	// Experimental.
	DesiredInstanceTypeInput() *string
	// Experimental.
	DesiredPartitionCount() *float64
	// Experimental.
	SetDesiredPartitionCount(val *float64)
	// Experimental.
	DesiredPartitionCountInput() *float64
	// Experimental.
	DesiredReplicationCount() *float64
	// Experimental.
	SetDesiredReplicationCount(val *float64)
	// Experimental.
	DesiredReplicationCountInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCloudsearchDomain_ScalingParametersProperty
	// Experimental.
	SetInternalValue(val *AwsCloudsearchDomain_ScalingParametersProperty)
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
	ResetDesiredInstanceType()
	// Experimental.
	ResetDesiredPartitionCount()
	// Experimental.
	ResetDesiredReplicationCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudsearchDomain_ScalingParametersPropertyOutputReference
type jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) DesiredInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) DesiredInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) DesiredPartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredPartitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) DesiredPartitionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredPartitionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) DesiredReplicationCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredReplicationCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) DesiredReplicationCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredReplicationCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) InternalValue() *AwsCloudsearchDomain_ScalingParametersProperty {
	var returns *AwsCloudsearchDomain_ScalingParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudsearchDomain_ScalingParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCloudsearchDomain_ScalingParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudsearchDomain_ScalingParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudsearch.AwsCloudsearchDomain.ScalingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudsearchDomain_ScalingParametersPropertyOutputReference_Override(a AwsCloudsearchDomain_ScalingParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudsearch.AwsCloudsearchDomain.ScalingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetDesiredInstanceType(val *string) {
	if err := j.validateSetDesiredInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetDesiredPartitionCount(val *float64) {
	if err := j.validateSetDesiredPartitionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredPartitionCount",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetDesiredReplicationCount(val *float64) {
	if err := j.validateSetDesiredReplicationCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredReplicationCount",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetInternalValue(val *AwsCloudsearchDomain_ScalingParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) ResetDesiredInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) ResetDesiredPartitionCount() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredPartitionCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) ResetDesiredReplicationCount() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredReplicationCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudsearchDomain_ScalingParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

