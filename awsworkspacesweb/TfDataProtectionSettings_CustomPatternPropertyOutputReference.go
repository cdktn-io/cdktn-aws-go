package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspacesweb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsworkspacesweb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataProtectionSettings_CustomPatternPropertyOutputReference interface {
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
	KeywordRegex() *string
	// Experimental.
	SetKeywordRegex(val *string)
	// Experimental.
	KeywordRegexInput() *string
	// Experimental.
	PatternDescription() *string
	// Experimental.
	SetPatternDescription(val *string)
	// Experimental.
	PatternDescriptionInput() *string
	// Experimental.
	PatternName() *string
	// Experimental.
	SetPatternName(val *string)
	// Experimental.
	PatternNameInput() *string
	// Experimental.
	PatternRegex() *string
	// Experimental.
	SetPatternRegex(val *string)
	// Experimental.
	PatternRegexInput() *string
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
	ResetKeywordRegex()
	// Experimental.
	ResetPatternDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataProtectionSettings_CustomPatternPropertyOutputReference
type jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) KeywordRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keywordRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) KeywordRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keywordRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) PatternDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) PatternDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) PatternName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) PatternNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) PatternRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) PatternRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataProtectionSettings_CustomPatternPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDataProtectionSettings_CustomPatternPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataProtectionSettings_CustomPatternPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.TfDataProtectionSettings.CustomPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataProtectionSettings_CustomPatternPropertyOutputReference_Override(t TfDataProtectionSettings_CustomPatternPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.TfDataProtectionSettings.CustomPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetKeywordRegex(val *string) {
	if err := j.validateSetKeywordRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keywordRegex",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetPatternDescription(val *string) {
	if err := j.validateSetPatternDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternDescription",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetPatternName(val *string) {
	if err := j.validateSetPatternNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternName",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetPatternRegex(val *string) {
	if err := j.validateSetPatternRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternRegex",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) ResetKeywordRegex() {
	_jsii_.InvokeVoid(
		t,
		"resetKeywordRegex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) ResetPatternDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetPatternDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataProtectionSettings_CustomPatternPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

