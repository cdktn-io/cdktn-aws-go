package awsauditmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsauditmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsauditmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfControl_ControlMappingSourcesPropertyOutputReference interface {
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
	SourceDescription() *string
	// Experimental.
	SetSourceDescription(val *string)
	// Experimental.
	SourceDescriptionInput() *string
	// Experimental.
	SourceFrequency() *string
	// Experimental.
	SetSourceFrequency(val *string)
	// Experimental.
	SourceFrequencyInput() *string
	// Experimental.
	SourceId() *string
	// Experimental.
	SourceKeyword() TfControl_SourceKeywordPropertyList
	// Experimental.
	SourceKeywordInput() interface{}
	// Experimental.
	SourceName() *string
	// Experimental.
	SetSourceName(val *string)
	// Experimental.
	SourceNameInput() *string
	// Experimental.
	SourceSetUpOption() *string
	// Experimental.
	SetSourceSetUpOption(val *string)
	// Experimental.
	SourceSetUpOptionInput() *string
	// Experimental.
	SourceType() *string
	// Experimental.
	SetSourceType(val *string)
	// Experimental.
	SourceTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TroubleshootingText() *string
	// Experimental.
	SetTroubleshootingText(val *string)
	// Experimental.
	TroubleshootingTextInput() *string
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
	PutSourceKeyword(value interface{})
	// Experimental.
	ResetSourceDescription()
	// Experimental.
	ResetSourceFrequency()
	// Experimental.
	ResetSourceKeyword()
	// Experimental.
	ResetTroubleshootingText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfControl_ControlMappingSourcesPropertyOutputReference
type jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceKeyword() TfControl_SourceKeywordPropertyList {
	var returns TfControl_SourceKeywordPropertyList
	_jsii_.Get(
		j,
		"sourceKeyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceKeywordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceKeywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceSetUpOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSetUpOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceSetUpOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSetUpOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) SourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) TroubleshootingText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"troubleshootingText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) TroubleshootingTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"troubleshootingTextInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfControl_ControlMappingSourcesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfControl_ControlMappingSourcesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfControl_ControlMappingSourcesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-audit-manager.TfControl.ControlMappingSourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfControl_ControlMappingSourcesPropertyOutputReference_Override(t TfControl_ControlMappingSourcesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-audit-manager.TfControl.ControlMappingSourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetSourceDescription(val *string) {
	if err := j.validateSetSourceDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDescription",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetSourceFrequency(val *string) {
	if err := j.validateSetSourceFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFrequency",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetSourceName(val *string) {
	if err := j.validateSetSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceName",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetSourceSetUpOption(val *string) {
	if err := j.validateSetSourceSetUpOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSetUpOption",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetSourceType(val *string) {
	if err := j.validateSetSourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference)SetTroubleshootingText(val *string) {
	if err := j.validateSetTroubleshootingTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"troubleshootingText",
		val,
	)
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) PutSourceKeyword(value interface{}) {
	if err := t.validatePutSourceKeywordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceKeyword",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ResetSourceDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ResetSourceFrequency() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceFrequency",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ResetSourceKeyword() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceKeyword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ResetTroubleshootingText() {
	_jsii_.InvokeVoid(
		t,
		"resetTroubleshootingText",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfControl_ControlMappingSourcesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

