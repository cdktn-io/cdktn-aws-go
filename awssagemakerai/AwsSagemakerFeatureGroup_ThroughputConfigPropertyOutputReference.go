package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference interface {
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
	InternalValue() *AwsSagemakerFeatureGroup_ThroughputConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerFeatureGroup_ThroughputConfigProperty)
	// Experimental.
	ProvisionedReadCapacityUnits() *float64
	// Experimental.
	SetProvisionedReadCapacityUnits(val *float64)
	// Experimental.
	ProvisionedReadCapacityUnitsInput() *float64
	// Experimental.
	ProvisionedWriteCapacityUnits() *float64
	// Experimental.
	SetProvisionedWriteCapacityUnits(val *float64)
	// Experimental.
	ProvisionedWriteCapacityUnitsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThroughputMode() *string
	// Experimental.
	SetThroughputMode(val *string)
	// Experimental.
	ThroughputModeInput() *string
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
	ResetProvisionedReadCapacityUnits()
	// Experimental.
	ResetProvisionedWriteCapacityUnits()
	// Experimental.
	ResetThroughputMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) InternalValue() *AwsSagemakerFeatureGroup_ThroughputConfigProperty {
	var returns *AwsSagemakerFeatureGroup_ThroughputConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ProvisionedReadCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedReadCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ProvisionedReadCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedReadCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ProvisionedWriteCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedWriteCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ProvisionedWriteCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedWriteCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ThroughputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throughputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ThroughputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throughputModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerFeatureGroup.ThroughputConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference_Override(a AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerFeatureGroup.ThroughputConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetInternalValue(val *AwsSagemakerFeatureGroup_ThroughputConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetProvisionedReadCapacityUnits(val *float64) {
	if err := j.validateSetProvisionedReadCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedReadCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetProvisionedWriteCapacityUnits(val *float64) {
	if err := j.validateSetProvisionedWriteCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedWriteCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference)SetThroughputMode(val *string) {
	if err := j.validateSetThroughputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputMode",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ResetProvisionedReadCapacityUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetProvisionedReadCapacityUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ResetProvisionedWriteCapacityUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetProvisionedWriteCapacityUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ResetThroughputMode() {
	_jsii_.InvokeVoid(
		a,
		"resetThroughputMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_ThroughputConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

