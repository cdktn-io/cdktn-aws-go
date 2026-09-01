package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearchserverless/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearchserverless/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxIndexingCapacityInOcu() *float64
	// Experimental.
	SetMaxIndexingCapacityInOcu(val *float64)
	// Experimental.
	MaxIndexingCapacityInOcuInput() *float64
	// Experimental.
	MaxSearchCapacityInOcu() *float64
	// Experimental.
	SetMaxSearchCapacityInOcu(val *float64)
	// Experimental.
	MaxSearchCapacityInOcuInput() *float64
	// Experimental.
	MinIndexingCapacityInOcu() *float64
	// Experimental.
	SetMinIndexingCapacityInOcu(val *float64)
	// Experimental.
	MinIndexingCapacityInOcuInput() *float64
	// Experimental.
	MinSearchCapacityInOcu() *float64
	// Experimental.
	SetMinSearchCapacityInOcu(val *float64)
	// Experimental.
	MinSearchCapacityInOcuInput() *float64
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
	ResetMaxIndexingCapacityInOcu()
	// Experimental.
	ResetMaxSearchCapacityInOcu()
	// Experimental.
	ResetMinIndexingCapacityInOcu()
	// Experimental.
	ResetMinSearchCapacityInOcu()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference
type jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MaxIndexingCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIndexingCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MaxIndexingCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIndexingCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MaxSearchCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSearchCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MaxSearchCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSearchCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MinIndexingCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIndexingCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MinIndexingCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIndexingCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MinSearchCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSearchCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) MinSearchCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSearchCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch-serverless.AwsOpensearchserverlessCollectionGroup.CapacityLimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference_Override(a AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch-serverless.AwsOpensearchserverlessCollectionGroup.CapacityLimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetMaxIndexingCapacityInOcu(val *float64) {
	if err := j.validateSetMaxIndexingCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIndexingCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetMaxSearchCapacityInOcu(val *float64) {
	if err := j.validateSetMaxSearchCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSearchCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetMinIndexingCapacityInOcu(val *float64) {
	if err := j.validateSetMinIndexingCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIndexingCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetMinSearchCapacityInOcu(val *float64) {
	if err := j.validateSetMinSearchCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSearchCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMaxIndexingCapacityInOcu() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxIndexingCapacityInOcu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMaxSearchCapacityInOcu() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxSearchCapacityInOcu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMinIndexingCapacityInOcu() {
	_jsii_.InvokeVoid(
		a,
		"resetMinIndexingCapacityInOcu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMinSearchCapacityInOcu() {
	_jsii_.InvokeVoid(
		a,
		"resetMinSearchCapacityInOcu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOpensearchserverlessCollectionGroup_CapacityLimitsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

