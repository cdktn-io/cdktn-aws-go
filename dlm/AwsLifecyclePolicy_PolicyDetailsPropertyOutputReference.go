package dlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() AwsLifecyclePolicy_ActionPropertyOutputReference
	// Experimental.
	ActionInput() *AwsLifecyclePolicy_ActionProperty
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
	EventSource() AwsLifecyclePolicy_EventSourcePropertyOutputReference
	// Experimental.
	EventSourceInput() *AwsLifecyclePolicy_EventSourceProperty
	// Experimental.
	Exclusions() AwsLifecyclePolicy_ExclusionsPropertyOutputReference
	// Experimental.
	ExclusionsInput() *AwsLifecyclePolicy_ExclusionsProperty
	// Experimental.
	ExtendDeletion() interface{}
	// Experimental.
	SetExtendDeletion(val interface{})
	// Experimental.
	ExtendDeletionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLifecyclePolicy_PolicyDetailsProperty
	// Experimental.
	SetInternalValue(val *AwsLifecyclePolicy_PolicyDetailsProperty)
	// Experimental.
	Parameters() AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference
	// Experimental.
	ParametersInput() *AwsLifecyclePolicy_PolicyDetailsParametersProperty
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
	Schedule() AwsLifecyclePolicy_SchedulePropertyList
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
	PutAction(value *AwsLifecyclePolicy_ActionProperty)
	// Experimental.
	PutEventSource(value *AwsLifecyclePolicy_EventSourceProperty)
	// Experimental.
	PutExclusions(value *AwsLifecyclePolicy_ExclusionsProperty)
	// Experimental.
	PutParameters(value *AwsLifecyclePolicy_PolicyDetailsParametersProperty)
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

// The jsii proxy struct for AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference
type jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) Action() AwsLifecyclePolicy_ActionPropertyOutputReference {
	var returns AwsLifecyclePolicy_ActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ActionInput() *AwsLifecyclePolicy_ActionProperty {
	var returns *AwsLifecyclePolicy_ActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreateInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreateIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) EventSource() AwsLifecyclePolicy_EventSourcePropertyOutputReference {
	var returns AwsLifecyclePolicy_EventSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) EventSourceInput() *AwsLifecyclePolicy_EventSourceProperty {
	var returns *AwsLifecyclePolicy_EventSourceProperty
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) Exclusions() AwsLifecyclePolicy_ExclusionsPropertyOutputReference {
	var returns AwsLifecyclePolicy_ExclusionsPropertyOutputReference
	_jsii_.Get(
		j,
		"exclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExclusionsInput() *AwsLifecyclePolicy_ExclusionsProperty {
	var returns *AwsLifecyclePolicy_ExclusionsProperty
	_jsii_.Get(
		j,
		"exclusionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExtendDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ExtendDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) InternalValue() *AwsLifecyclePolicy_PolicyDetailsProperty {
	var returns *AwsLifecyclePolicy_PolicyDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) Parameters() AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference {
	var returns AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ParametersInput() *AwsLifecyclePolicy_PolicyDetailsParametersProperty {
	var returns *AwsLifecyclePolicy_PolicyDetailsParametersProperty
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) RetainInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) RetainIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) Schedule() AwsLifecyclePolicy_SchedulePropertyList {
	var returns AwsLifecyclePolicy_SchedulePropertyList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) TargetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) TargetTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLifecyclePolicy_PolicyDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLifecyclePolicy_PolicyDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.PolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLifecyclePolicy_PolicyDetailsPropertyOutputReference_Override(a AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.PolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetCreateInterval(val *float64) {
	if err := j.validateSetCreateIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createInterval",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetExtendDeletion(val interface{}) {
	if err := j.validateSetExtendDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendDeletion",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetInternalValue(val *AwsLifecyclePolicy_PolicyDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetPolicyLanguage(val *string) {
	if err := j.validateSetPolicyLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyLanguage",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceLocations(val *[]*string) {
	if err := j.validateSetResourceLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLocations",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetRetainInterval(val *float64) {
	if err := j.validateSetRetainIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainInterval",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTargetTags(val *map[string]*string) {
	if err := j.validateSetTargetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTags",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutAction(value *AwsLifecyclePolicy_ActionProperty) {
	if err := a.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutEventSource(value *AwsLifecyclePolicy_EventSourceProperty) {
	if err := a.validatePutEventSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutExclusions(value *AwsLifecyclePolicy_ExclusionsProperty) {
	if err := a.validatePutExclusionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclusions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutParameters(value *AwsLifecyclePolicy_PolicyDetailsParametersProperty) {
	if err := a.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) PutSchedule(value interface{}) {
	if err := a.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		a,
		"resetAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetCreateInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetCreateInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetEventSource() {
	_jsii_.InvokeVoid(
		a,
		"resetEventSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetExclusions() {
	_jsii_.InvokeVoid(
		a,
		"resetExclusions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetExtendDeletion() {
	_jsii_.InvokeVoid(
		a,
		"resetExtendDeletion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetPolicyLanguage() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyLanguage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetPolicyType() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceLocations() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceLocations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetRetainInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetRetainInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ResetTargetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

