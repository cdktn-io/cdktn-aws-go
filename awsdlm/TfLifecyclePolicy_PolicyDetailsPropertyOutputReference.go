package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLifecyclePolicy_PolicyDetailsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() TfLifecyclePolicy_ActionPropertyOutputReference
	// Experimental.
	ActionInput() *TfLifecyclePolicy_ActionProperty
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
	CreateInterval() *float64
	// Experimental.
	SetCreateInterval(val *float64)
	// Experimental.
	CreateIntervalInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EventSource() TfLifecyclePolicy_EventSourcePropertyOutputReference
	// Experimental.
	EventSourceInput() *TfLifecyclePolicy_EventSourceProperty
	// Experimental.
	Exclusions() TfLifecyclePolicy_ExclusionsPropertyOutputReference
	// Experimental.
	ExclusionsInput() *TfLifecyclePolicy_ExclusionsProperty
	// Experimental.
	ExtendDeletion() interface{}
	// Experimental.
	SetExtendDeletion(val interface{})
	// Experimental.
	ExtendDeletionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfLifecyclePolicy_PolicyDetailsProperty
	// Experimental.
	SetInternalValue(val *TfLifecyclePolicy_PolicyDetailsProperty)
	// Experimental.
	Parameters() TfLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference
	// Experimental.
	ParametersInput() *TfLifecyclePolicy_PolicyDetailsParametersProperty
	// Experimental.
	PolicyLanguage() *string
	// Experimental.
	SetPolicyLanguage(val *string)
	// Experimental.
	PolicyLanguageInput() *string
	// Experimental.
	PolicyType() *string
	// Experimental.
	SetPolicyType(val *string)
	// Experimental.
	PolicyTypeInput() *string
	// Experimental.
	ResourceLocations() *[]*string
	// Experimental.
	SetResourceLocations(val *[]*string)
	// Experimental.
	ResourceLocationsInput() *[]*string
	// Experimental.
	ResourceType() *string
	// Experimental.
	SetResourceType(val *string)
	// Experimental.
	ResourceTypeInput() *string
	// Experimental.
	ResourceTypes() *[]*string
	// Experimental.
	SetResourceTypes(val *[]*string)
	// Experimental.
	ResourceTypesInput() *[]*string
	// Experimental.
	RetainInterval() *float64
	// Experimental.
	SetRetainInterval(val *float64)
	// Experimental.
	RetainIntervalInput() *float64
	// Experimental.
	Schedule() TfLifecyclePolicy_SchedulePropertyList
	// Experimental.
	ScheduleInput() interface{}
	// Experimental.
	TargetTags() *map[string]*string
	// Experimental.
	SetTargetTags(val *map[string]*string)
	// Experimental.
	TargetTagsInput() *map[string]*string
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
	PutAction(value *TfLifecyclePolicy_ActionProperty)
	// Experimental.
	PutEventSource(value *TfLifecyclePolicy_EventSourceProperty)
	// Experimental.
	PutExclusions(value *TfLifecyclePolicy_ExclusionsProperty)
	// Experimental.
	PutParameters(value *TfLifecyclePolicy_PolicyDetailsParametersProperty)
	// Experimental.
	PutSchedule(value interface{})
	// Experimental.
	ResetAction()
	// Experimental.
	ResetCopyTags()
	// Experimental.
	ResetCreateInterval()
	// Experimental.
	ResetEventSource()
	// Experimental.
	ResetExclusions()
	// Experimental.
	ResetExtendDeletion()
	// Experimental.
	ResetParameters()
	// Experimental.
	ResetPolicyLanguage()
	// Experimental.
	ResetPolicyType()
	// Experimental.
	ResetResourceLocations()
	// Experimental.
	ResetResourceType()
	// Experimental.
	ResetResourceTypes()
	// Experimental.
	ResetRetainInterval()
	// Experimental.
	ResetSchedule()
	// Experimental.
	ResetTargetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLifecyclePolicy_PolicyDetailsPropertyOutputReference
type jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) Action() TfLifecyclePolicy_ActionPropertyOutputReference {
	var returns TfLifecyclePolicy_ActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ActionInput() *TfLifecyclePolicy_ActionProperty {
	var returns *TfLifecyclePolicy_ActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreateInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreateIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) EventSource() TfLifecyclePolicy_EventSourcePropertyOutputReference {
	var returns TfLifecyclePolicy_EventSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) EventSourceInput() *TfLifecyclePolicy_EventSourceProperty {
	var returns *TfLifecyclePolicy_EventSourceProperty
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) Exclusions() TfLifecyclePolicy_ExclusionsPropertyOutputReference {
	var returns TfLifecyclePolicy_ExclusionsPropertyOutputReference
	_jsii_.Get(
		j,
		"exclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExclusionsInput() *TfLifecyclePolicy_ExclusionsProperty {
	var returns *TfLifecyclePolicy_ExclusionsProperty
	_jsii_.Get(
		j,
		"exclusionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExtendDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExtendDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) InternalValue() *TfLifecyclePolicy_PolicyDetailsProperty {
	var returns *TfLifecyclePolicy_PolicyDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) Parameters() TfLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference {
	var returns TfLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ParametersInput() *TfLifecyclePolicy_PolicyDetailsParametersProperty {
	var returns *TfLifecyclePolicy_PolicyDetailsParametersProperty
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) RetainInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) RetainIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) Schedule() TfLifecyclePolicy_SchedulePropertyList {
	var returns TfLifecyclePolicy_SchedulePropertyList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) TargetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) TargetTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLifecyclePolicy_PolicyDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLifecyclePolicy_PolicyDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLifecyclePolicy_PolicyDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.PolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLifecyclePolicy_PolicyDetailsPropertyOutputReference_Override(t TfLifecyclePolicy_PolicyDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.PolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetCreateInterval(val *float64) {
	if err := j.validateSetCreateIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createInterval",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetExtendDeletion(val interface{}) {
	if err := j.validateSetExtendDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendDeletion",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetInternalValue(val *TfLifecyclePolicy_PolicyDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetPolicyLanguage(val *string) {
	if err := j.validateSetPolicyLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyLanguage",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceLocations(val *[]*string) {
	if err := j.validateSetResourceLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLocations",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetRetainInterval(val *float64) {
	if err := j.validateSetRetainIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainInterval",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTargetTags(val *map[string]*string) {
	if err := j.validateSetTargetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTags",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutAction(value *TfLifecyclePolicy_ActionProperty) {
	if err := t.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutEventSource(value *TfLifecyclePolicy_EventSourceProperty) {
	if err := t.validatePutEventSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEventSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutExclusions(value *TfLifecyclePolicy_ExclusionsProperty) {
	if err := t.validatePutExclusionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExclusions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutParameters(value *TfLifecyclePolicy_PolicyDetailsParametersProperty) {
	if err := t.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutSchedule(value interface{}) {
	if err := t.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchedule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		t,
		"resetAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetCreateInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetEventSource() {
	_jsii_.InvokeVoid(
		t,
		"resetEventSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetExclusions() {
	_jsii_.InvokeVoid(
		t,
		"resetExclusions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetExtendDeletion() {
	_jsii_.InvokeVoid(
		t,
		"resetExtendDeletion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetPolicyLanguage() {
	_jsii_.InvokeVoid(
		t,
		"resetPolicyLanguage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetPolicyType() {
	_jsii_.InvokeVoid(
		t,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceLocations() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceLocations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetRetainInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetRetainInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetTargetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_PolicyDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

