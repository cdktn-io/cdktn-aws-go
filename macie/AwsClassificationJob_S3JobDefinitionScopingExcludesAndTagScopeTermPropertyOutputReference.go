package macie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/macie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/macie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Comparator() *string
	// Experimental.
	SetComparator(val *string)
	// Experimental.
	ComparatorInput() *string
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
	InternalValue() *AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty
	// Experimental.
	SetInternalValue(val *AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty)
	// Experimental.
	Key() *string
	// Experimental.
	SetKey(val *string)
	// Experimental.
	KeyInput() *string
	// Experimental.
	TagValues() AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermTagValuesPropertyList
	// Experimental.
	TagValuesInput() interface{}
	// Experimental.
	Target() *string
	// Experimental.
	SetTarget(val *string)
	// Experimental.
	TargetInput() *string
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
	PutTagValues(value interface{})
	// Experimental.
	ResetComparator()
	// Experimental.
	ResetKey()
	// Experimental.
	ResetTagValues()
	// Experimental.
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference
type jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) InternalValue() *AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty {
	var returns *AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) TagValues() AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermTagValuesPropertyList {
	var returns AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermTagValuesPropertyList
	_jsii_.Get(
		j,
		"tagValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) TagValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.AwsClassificationJob.S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference_Override(a AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.AwsClassificationJob.S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetComparator(val *string) {
	if err := j.validateSetComparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetInternalValue(val *AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) PutTagValues(value interface{}) {
	if err := a.validatePutTagValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagValues",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		a,
		"resetComparator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		a,
		"resetKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ResetTagValues() {
	_jsii_.InvokeVoid(
		a,
		"resetTagValues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

