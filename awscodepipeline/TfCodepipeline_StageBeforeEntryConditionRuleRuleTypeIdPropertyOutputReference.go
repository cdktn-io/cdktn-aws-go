package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Category() *string
	// Experimental.
	SetCategory(val *string)
	// Experimental.
	CategoryInput() *string
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
	InternalValue() *TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdProperty
	// Experimental.
	SetInternalValue(val *TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdProperty)
	// Experimental.
	Owner() *string
	// Experimental.
	SetOwner(val *string)
	// Experimental.
	OwnerInput() *string
	// Experimental.
	Provider() *string
	// Experimental.
	SetProvider(val *string)
	// Experimental.
	ProviderInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
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
	ResetOwner()
	// Experimental.
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference
type jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) InternalValue() *TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdProperty {
	var returns *TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference_Override(t TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetCategory(val *string) {
	if err := j.validateSetCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetInternalValue(val *TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		t,
		"resetOwner",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

