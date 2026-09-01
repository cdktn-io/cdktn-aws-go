package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmacie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmacie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference interface {
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
	SimpleScopeTerm() AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermPropertyOutputReference
	// Experimental.
	SimpleScopeTermInput() *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermProperty
	// Experimental.
	TagScopeTerm() AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference
	// Experimental.
	TagScopeTermInput() *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty
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
	PutSimpleScopeTerm(value *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermProperty)
	// Experimental.
	PutTagScopeTerm(value *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty)
	// Experimental.
	ResetSimpleScopeTerm()
	// Experimental.
	ResetTagScopeTerm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference
type jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) SimpleScopeTerm() AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermPropertyOutputReference {
	var returns AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermPropertyOutputReference
	_jsii_.Get(
		j,
		"simpleScopeTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) SimpleScopeTermInput() *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermProperty {
	var returns *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermProperty
	_jsii_.Get(
		j,
		"simpleScopeTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) TagScopeTerm() AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference {
	var returns AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermPropertyOutputReference
	_jsii_.Get(
		j,
		"tagScopeTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) TagScopeTermInput() *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty {
	var returns *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty
	_jsii_.Get(
		j,
		"tagScopeTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.AwsMacie2ClassificationJob.S3JobDefinitionScopingExcludesAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference_Override(a AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.AwsMacie2ClassificationJob.S3JobDefinitionScopingExcludesAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) PutSimpleScopeTerm(value *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermProperty) {
	if err := a.validatePutSimpleScopeTermParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSimpleScopeTerm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) PutTagScopeTerm(value *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty) {
	if err := a.validatePutTagScopeTermParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagScopeTerm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) ResetSimpleScopeTerm() {
	_jsii_.InvokeVoid(
		a,
		"resetSimpleScopeTerm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) ResetTagScopeTerm() {
	_jsii_.InvokeVoid(
		a,
		"resetTagScopeTerm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

