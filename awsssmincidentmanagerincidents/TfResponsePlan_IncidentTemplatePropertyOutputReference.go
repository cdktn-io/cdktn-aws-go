package awsssmincidentmanagerincidents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssmincidentmanagerincidents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssmincidentmanagerincidents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResponsePlan_IncidentTemplatePropertyOutputReference interface {
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
	DedupeString() *string
	// Experimental.
	SetDedupeString(val *string)
	// Experimental.
	DedupeStringInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	Impact() *float64
	// Experimental.
	SetImpact(val *float64)
	// Experimental.
	ImpactInput() *float64
	// Experimental.
	IncidentTags() *map[string]*string
	// Experimental.
	SetIncidentTags(val *map[string]*string)
	// Experimental.
	IncidentTagsInput() *map[string]*string
	// Experimental.
	InternalValue() *TfResponsePlan_IncidentTemplateProperty
	// Experimental.
	SetInternalValue(val *TfResponsePlan_IncidentTemplateProperty)
	// Experimental.
	NotificationTarget() TfResponsePlan_NotificationTargetPropertyList
	// Experimental.
	NotificationTargetInput() interface{}
	// Experimental.
	Summary() *string
	// Experimental.
	SetSummary(val *string)
	// Experimental.
	SummaryInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Title() *string
	// Experimental.
	SetTitle(val *string)
	// Experimental.
	TitleInput() *string
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
	PutNotificationTarget(value interface{})
	// Experimental.
	ResetDedupeString()
	// Experimental.
	ResetIncidentTags()
	// Experimental.
	ResetNotificationTarget()
	// Experimental.
	ResetSummary()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfResponsePlan_IncidentTemplatePropertyOutputReference
type jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) DedupeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedupeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) DedupeStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedupeStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) Impact() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ImpactInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) IncidentTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"incidentTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) IncidentTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"incidentTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) InternalValue() *TfResponsePlan_IncidentTemplateProperty {
	var returns *TfResponsePlan_IncidentTemplateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) NotificationTarget() TfResponsePlan_NotificationTargetPropertyList {
	var returns TfResponsePlan_NotificationTargetPropertyList
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) NotificationTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfResponsePlan_IncidentTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfResponsePlan_IncidentTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfResponsePlan_IncidentTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.TfResponsePlan.IncidentTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfResponsePlan_IncidentTemplatePropertyOutputReference_Override(t TfResponsePlan_IncidentTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.TfResponsePlan.IncidentTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetDedupeString(val *string) {
	if err := j.validateSetDedupeStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedupeString",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetImpact(val *float64) {
	if err := j.validateSetImpactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"impact",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetIncidentTags(val *map[string]*string) {
	if err := j.validateSetIncidentTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incidentTags",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetInternalValue(val *TfResponsePlan_IncidentTemplateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) PutNotificationTarget(value interface{}) {
	if err := t.validatePutNotificationTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotificationTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ResetDedupeString() {
	_jsii_.InvokeVoid(
		t,
		"resetDedupeString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ResetIncidentTags() {
	_jsii_.InvokeVoid(
		t,
		"resetIncidentTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ResetNotificationTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetNotificationTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ResetSummary() {
	_jsii_.InvokeVoid(
		t,
		"resetSummary",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfResponsePlan_IncidentTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

