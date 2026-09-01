package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() AwsDlmLifecyclePolicy_ActionPropertyOutputReference
	// Experimental.
	ActionInput() *AwsDlmLifecyclePolicy_ActionProperty
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
	EventSource() AwsDlmLifecyclePolicy_EventSourcePropertyOutputReference
	// Experimental.
	EventSourceInput() *AwsDlmLifecyclePolicy_EventSourceProperty
	// Experimental.
	Exclusions() AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference
	// Experimental.
	ExclusionsInput() *AwsDlmLifecyclePolicy_ExclusionsProperty
	// Experimental.
	ExtendDeletion() interface{}
	// Experimental.
	SetExtendDeletion(val interface{})
	// Experimental.
	ExtendDeletionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDlmLifecyclePolicy_PolicyDetailsProperty
	// Experimental.
	SetInternalValue(val *AwsDlmLifecyclePolicy_PolicyDetailsProperty)
	// Experimental.
	Parameters() AwsDlmLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference
	// Experimental.
	ParametersInput() *AwsDlmLifecyclePolicy_PolicyDetailsParametersProperty
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
	Schedule() AwsDlmLifecyclePolicy_SchedulePropertyList
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
	PutAction(value *AwsDlmLifecyclePolicy_ActionProperty)
	// Experimental.
	PutEventSource(value *AwsDlmLifecyclePolicy_EventSourceProperty)
	// Experimental.
	PutExclusions(value *AwsDlmLifecyclePolicy_ExclusionsProperty)
	// Experimental.
	PutParameters(value *AwsDlmLifecyclePolicy_PolicyDetailsParametersProperty)
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

// The jsii proxy struct for AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference
type jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) Action() AwsDlmLifecyclePolicy_ActionPropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_ActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ActionInput() *AwsDlmLifecyclePolicy_ActionProperty {
	var returns *AwsDlmLifecyclePolicy_ActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreateInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreateIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) EventSource() AwsDlmLifecyclePolicy_EventSourcePropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_EventSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) EventSourceInput() *AwsDlmLifecyclePolicy_EventSourceProperty {
	var returns *AwsDlmLifecyclePolicy_EventSourceProperty
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) Exclusions() AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference
	_jsii_.Get(
		j,
		"exclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExclusionsInput() *AwsDlmLifecyclePolicy_ExclusionsProperty {
	var returns *AwsDlmLifecyclePolicy_ExclusionsProperty
	_jsii_.Get(
		j,
		"exclusionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExtendDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExtendDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) InternalValue() *AwsDlmLifecyclePolicy_PolicyDetailsProperty {
	var returns *AwsDlmLifecyclePolicy_PolicyDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) Parameters() AwsDlmLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ParametersInput() *AwsDlmLifecyclePolicy_PolicyDetailsParametersProperty {
	var returns *AwsDlmLifecyclePolicy_PolicyDetailsParametersProperty
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) RetainInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) RetainIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) Schedule() AwsDlmLifecyclePolicy_SchedulePropertyList {
	var returns AwsDlmLifecyclePolicy_SchedulePropertyList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) TargetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) TargetTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.PolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference_Override(a AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.PolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetCreateInterval(val *float64) {
	if err := j.validateSetCreateIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createInterval",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetExtendDeletion(val interface{}) {
	if err := j.validateSetExtendDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendDeletion",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetInternalValue(val *AwsDlmLifecyclePolicy_PolicyDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetPolicyLanguage(val *string) {
	if err := j.validateSetPolicyLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyLanguage",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceLocations(val *[]*string) {
	if err := j.validateSetResourceLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLocations",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetRetainInterval(val *float64) {
	if err := j.validateSetRetainIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainInterval",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTargetTags(val *map[string]*string) {
	if err := j.validateSetTargetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTags",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutAction(value *AwsDlmLifecyclePolicy_ActionProperty) {
	if err := a.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutEventSource(value *AwsDlmLifecyclePolicy_EventSourceProperty) {
	if err := a.validatePutEventSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutExclusions(value *AwsDlmLifecyclePolicy_ExclusionsProperty) {
	if err := a.validatePutExclusionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclusions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutParameters(value *AwsDlmLifecyclePolicy_PolicyDetailsParametersProperty) {
	if err := a.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutSchedule(value interface{}) {
	if err := a.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		a,
		"resetAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetCreateInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetCreateInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetEventSource() {
	_jsii_.InvokeVoid(
		a,
		"resetEventSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetExclusions() {
	_jsii_.InvokeVoid(
		a,
		"resetExclusions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetExtendDeletion() {
	_jsii_.InvokeVoid(
		a,
		"resetExtendDeletion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetPolicyLanguage() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyLanguage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetPolicyType() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceLocations() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceLocations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetRetainInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetRetainInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetTargetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_PolicyDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

