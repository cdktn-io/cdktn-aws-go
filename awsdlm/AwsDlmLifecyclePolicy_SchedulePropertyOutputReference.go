package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDlmLifecyclePolicy_SchedulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveRule() AwsDlmLifecyclePolicy_ArchiveRulePropertyOutputReference
	// Experimental.
	ArchiveRuleInput() *AwsDlmLifecyclePolicy_ArchiveRuleProperty
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
	CreateRule() AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference
	// Experimental.
	CreateRuleInput() *AwsDlmLifecyclePolicy_CreateRuleProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CrossRegionCopyRule() AwsDlmLifecyclePolicy_CrossRegionCopyRulePropertyList
	// Experimental.
	CrossRegionCopyRuleInput() interface{}
	// Experimental.
	DeprecateRule() AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRulePropertyOutputReference
	// Experimental.
	DeprecateRuleInput() *AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty
	// Experimental.
	FastRestoreRule() AwsDlmLifecyclePolicy_FastRestoreRulePropertyOutputReference
	// Experimental.
	FastRestoreRuleInput() *AwsDlmLifecyclePolicy_FastRestoreRuleProperty
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
	RetainRule() AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRulePropertyOutputReference
	// Experimental.
	RetainRuleInput() *AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty
	// Experimental.
	ShareRule() AwsDlmLifecyclePolicy_ShareRulePropertyOutputReference
	// Experimental.
	ShareRuleInput() *AwsDlmLifecyclePolicy_ShareRuleProperty
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
	PutArchiveRule(value *AwsDlmLifecyclePolicy_ArchiveRuleProperty)
	// Experimental.
	PutCreateRule(value *AwsDlmLifecyclePolicy_CreateRuleProperty)
	// Experimental.
	PutCrossRegionCopyRule(value interface{})
	// Experimental.
	PutDeprecateRule(value *AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty)
	// Experimental.
	PutFastRestoreRule(value *AwsDlmLifecyclePolicy_FastRestoreRuleProperty)
	// Experimental.
	PutRetainRule(value *AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty)
	// Experimental.
	PutShareRule(value *AwsDlmLifecyclePolicy_ShareRuleProperty)
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

// The jsii proxy struct for AwsDlmLifecyclePolicy_SchedulePropertyOutputReference
type jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ArchiveRule() AwsDlmLifecyclePolicy_ArchiveRulePropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_ArchiveRulePropertyOutputReference
	_jsii_.Get(
		j,
		"archiveRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ArchiveRuleInput() *AwsDlmLifecyclePolicy_ArchiveRuleProperty {
	var returns *AwsDlmLifecyclePolicy_ArchiveRuleProperty
	_jsii_.Get(
		j,
		"archiveRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) CreateRule() AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference
	_jsii_.Get(
		j,
		"createRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) CreateRuleInput() *AwsDlmLifecyclePolicy_CreateRuleProperty {
	var returns *AwsDlmLifecyclePolicy_CreateRuleProperty
	_jsii_.Get(
		j,
		"createRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) CrossRegionCopyRule() AwsDlmLifecyclePolicy_CrossRegionCopyRulePropertyList {
	var returns AwsDlmLifecyclePolicy_CrossRegionCopyRulePropertyList
	_jsii_.Get(
		j,
		"crossRegionCopyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) CrossRegionCopyRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionCopyRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) DeprecateRule() AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRulePropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRulePropertyOutputReference
	_jsii_.Get(
		j,
		"deprecateRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) DeprecateRuleInput() *AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty {
	var returns *AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty
	_jsii_.Get(
		j,
		"deprecateRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) FastRestoreRule() AwsDlmLifecyclePolicy_FastRestoreRulePropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_FastRestoreRulePropertyOutputReference
	_jsii_.Get(
		j,
		"fastRestoreRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) FastRestoreRuleInput() *AwsDlmLifecyclePolicy_FastRestoreRuleProperty {
	var returns *AwsDlmLifecyclePolicy_FastRestoreRuleProperty
	_jsii_.Get(
		j,
		"fastRestoreRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) RetainRule() AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRulePropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRulePropertyOutputReference
	_jsii_.Get(
		j,
		"retainRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) RetainRuleInput() *AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty {
	var returns *AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty
	_jsii_.Get(
		j,
		"retainRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ShareRule() AwsDlmLifecyclePolicy_ShareRulePropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_ShareRulePropertyOutputReference
	_jsii_.Get(
		j,
		"shareRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ShareRuleInput() *AwsDlmLifecyclePolicy_ShareRuleProperty {
	var returns *AwsDlmLifecyclePolicy_ShareRuleProperty
	_jsii_.Get(
		j,
		"shareRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) TagsToAdd() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) TagsToAddInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) VariableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) VariableTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableTagsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDlmLifecyclePolicy_SchedulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDlmLifecyclePolicy_SchedulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDlmLifecyclePolicy_SchedulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.SchedulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDlmLifecyclePolicy_SchedulePropertyOutputReference_Override(a AwsDlmLifecyclePolicy_SchedulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.SchedulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetTagsToAdd(val *map[string]*string) {
	if err := j.validateSetTagsToAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsToAdd",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference)SetVariableTags(val *map[string]*string) {
	if err := j.validateSetVariableTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variableTags",
		val,
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) PutArchiveRule(value *AwsDlmLifecyclePolicy_ArchiveRuleProperty) {
	if err := a.validatePutArchiveRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArchiveRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) PutCreateRule(value *AwsDlmLifecyclePolicy_CreateRuleProperty) {
	if err := a.validatePutCreateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreateRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) PutCrossRegionCopyRule(value interface{}) {
	if err := a.validatePutCrossRegionCopyRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCrossRegionCopyRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) PutDeprecateRule(value *AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty) {
	if err := a.validatePutDeprecateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeprecateRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) PutFastRestoreRule(value *AwsDlmLifecyclePolicy_FastRestoreRuleProperty) {
	if err := a.validatePutFastRestoreRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFastRestoreRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) PutRetainRule(value *AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty) {
	if err := a.validatePutRetainRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetainRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) PutShareRule(value *AwsDlmLifecyclePolicy_ShareRuleProperty) {
	if err := a.validatePutShareRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putShareRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetArchiveRule() {
	_jsii_.InvokeVoid(
		a,
		"resetArchiveRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetCrossRegionCopyRule() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossRegionCopyRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetDeprecateRule() {
	_jsii_.InvokeVoid(
		a,
		"resetDeprecateRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetFastRestoreRule() {
	_jsii_.InvokeVoid(
		a,
		"resetFastRestoreRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetShareRule() {
	_jsii_.InvokeVoid(
		a,
		"resetShareRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetTagsToAdd() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsToAdd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ResetVariableTags() {
	_jsii_.InvokeVoid(
		a,
		"resetVariableTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_SchedulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

