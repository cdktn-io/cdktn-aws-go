package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLifecyclePolicy_SchedulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveRule() TfLifecyclePolicy_ArchiveRulePropertyOutputReference
	// Experimental.
	ArchiveRuleInput() *TfLifecyclePolicy_ArchiveRuleProperty
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
	CopyTags() interface{}
	// Experimental.
	SetCopyTags(val interface{})
	// Experimental.
	CopyTagsInput() interface{}
	// Experimental.
	CreateRule() TfLifecyclePolicy_CreateRulePropertyOutputReference
	// Experimental.
	CreateRuleInput() *TfLifecyclePolicy_CreateRuleProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CrossRegionCopyRule() TfLifecyclePolicy_CrossRegionCopyRulePropertyList
	// Experimental.
	CrossRegionCopyRuleInput() interface{}
	// Experimental.
	DeprecateRule() TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRulePropertyOutputReference
	// Experimental.
	DeprecateRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty
	// Experimental.
	FastRestoreRule() TfLifecyclePolicy_FastRestoreRulePropertyOutputReference
	// Experimental.
	FastRestoreRuleInput() *TfLifecyclePolicy_FastRestoreRuleProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	RetainRule() TfLifecyclePolicy_PolicyDetailsScheduleRetainRulePropertyOutputReference
	// Experimental.
	RetainRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty
	// Experimental.
	ShareRule() TfLifecyclePolicy_ShareRulePropertyOutputReference
	// Experimental.
	ShareRuleInput() *TfLifecyclePolicy_ShareRuleProperty
	// Experimental.
	TagsToAdd() *map[string]*string
	// Experimental.
	SetTagsToAdd(val *map[string]*string)
	// Experimental.
	TagsToAddInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VariableTags() *map[string]*string
	// Experimental.
	SetVariableTags(val *map[string]*string)
	// Experimental.
	VariableTagsInput() *map[string]*string
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
	PutArchiveRule(value *TfLifecyclePolicy_ArchiveRuleProperty)
	// Experimental.
	PutCreateRule(value *TfLifecyclePolicy_CreateRuleProperty)
	// Experimental.
	PutCrossRegionCopyRule(value interface{})
	// Experimental.
	PutDeprecateRule(value *TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty)
	// Experimental.
	PutFastRestoreRule(value *TfLifecyclePolicy_FastRestoreRuleProperty)
	// Experimental.
	PutRetainRule(value *TfLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty)
	// Experimental.
	PutShareRule(value *TfLifecyclePolicy_ShareRuleProperty)
	// Experimental.
	ResetArchiveRule()
	// Experimental.
	ResetCopyTags()
	// Experimental.
	ResetCrossRegionCopyRule()
	// Experimental.
	ResetDeprecateRule()
	// Experimental.
	ResetFastRestoreRule()
	// Experimental.
	ResetShareRule()
	// Experimental.
	ResetTagsToAdd()
	// Experimental.
	ResetVariableTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLifecyclePolicy_SchedulePropertyOutputReference
type jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ArchiveRule() TfLifecyclePolicy_ArchiveRulePropertyOutputReference {
	var returns TfLifecyclePolicy_ArchiveRulePropertyOutputReference
	_jsii_.Get(
		j,
		"archiveRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ArchiveRuleInput() *TfLifecyclePolicy_ArchiveRuleProperty {
	var returns *TfLifecyclePolicy_ArchiveRuleProperty
	_jsii_.Get(
		j,
		"archiveRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) CreateRule() TfLifecyclePolicy_CreateRulePropertyOutputReference {
	var returns TfLifecyclePolicy_CreateRulePropertyOutputReference
	_jsii_.Get(
		j,
		"createRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) CreateRuleInput() *TfLifecyclePolicy_CreateRuleProperty {
	var returns *TfLifecyclePolicy_CreateRuleProperty
	_jsii_.Get(
		j,
		"createRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) CrossRegionCopyRule() TfLifecyclePolicy_CrossRegionCopyRulePropertyList {
	var returns TfLifecyclePolicy_CrossRegionCopyRulePropertyList
	_jsii_.Get(
		j,
		"crossRegionCopyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) CrossRegionCopyRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionCopyRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) DeprecateRule() TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRulePropertyOutputReference {
	var returns TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRulePropertyOutputReference
	_jsii_.Get(
		j,
		"deprecateRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) DeprecateRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty {
	var returns *TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty
	_jsii_.Get(
		j,
		"deprecateRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) FastRestoreRule() TfLifecyclePolicy_FastRestoreRulePropertyOutputReference {
	var returns TfLifecyclePolicy_FastRestoreRulePropertyOutputReference
	_jsii_.Get(
		j,
		"fastRestoreRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) FastRestoreRuleInput() *TfLifecyclePolicy_FastRestoreRuleProperty {
	var returns *TfLifecyclePolicy_FastRestoreRuleProperty
	_jsii_.Get(
		j,
		"fastRestoreRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) RetainRule() TfLifecyclePolicy_PolicyDetailsScheduleRetainRulePropertyOutputReference {
	var returns TfLifecyclePolicy_PolicyDetailsScheduleRetainRulePropertyOutputReference
	_jsii_.Get(
		j,
		"retainRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) RetainRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty {
	var returns *TfLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty
	_jsii_.Get(
		j,
		"retainRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ShareRule() TfLifecyclePolicy_ShareRulePropertyOutputReference {
	var returns TfLifecyclePolicy_ShareRulePropertyOutputReference
	_jsii_.Get(
		j,
		"shareRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ShareRuleInput() *TfLifecyclePolicy_ShareRuleProperty {
	var returns *TfLifecyclePolicy_ShareRuleProperty
	_jsii_.Get(
		j,
		"shareRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) TagsToAdd() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) TagsToAddInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) VariableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) VariableTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableTagsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLifecyclePolicy_SchedulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfLifecyclePolicy_SchedulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLifecyclePolicy_SchedulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.SchedulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLifecyclePolicy_SchedulePropertyOutputReference_Override(t TfLifecyclePolicy_SchedulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.SchedulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetTagsToAdd(val *map[string]*string) {
	if err := j.validateSetTagsToAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsToAdd",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference)SetVariableTags(val *map[string]*string) {
	if err := j.validateSetVariableTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variableTags",
		val,
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) PutArchiveRule(value *TfLifecyclePolicy_ArchiveRuleProperty) {
	if err := t.validatePutArchiveRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArchiveRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) PutCreateRule(value *TfLifecyclePolicy_CreateRuleProperty) {
	if err := t.validatePutCreateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreateRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) PutCrossRegionCopyRule(value interface{}) {
	if err := t.validatePutCrossRegionCopyRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCrossRegionCopyRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) PutDeprecateRule(value *TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty) {
	if err := t.validatePutDeprecateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeprecateRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) PutFastRestoreRule(value *TfLifecyclePolicy_FastRestoreRuleProperty) {
	if err := t.validatePutFastRestoreRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFastRestoreRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) PutRetainRule(value *TfLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty) {
	if err := t.validatePutRetainRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetainRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) PutShareRule(value *TfLifecyclePolicy_ShareRuleProperty) {
	if err := t.validatePutShareRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putShareRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetArchiveRule() {
	_jsii_.InvokeVoid(
		t,
		"resetArchiveRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetCrossRegionCopyRule() {
	_jsii_.InvokeVoid(
		t,
		"resetCrossRegionCopyRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetDeprecateRule() {
	_jsii_.InvokeVoid(
		t,
		"resetDeprecateRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetFastRestoreRule() {
	_jsii_.InvokeVoid(
		t,
		"resetFastRestoreRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetShareRule() {
	_jsii_.InvokeVoid(
		t,
		"resetShareRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetTagsToAdd() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsToAdd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ResetVariableTags() {
	_jsii_.InvokeVoid(
		t,
		"resetVariableTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_SchedulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

