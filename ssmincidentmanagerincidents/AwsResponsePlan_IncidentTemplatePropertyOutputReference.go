package ssmincidentmanagerincidents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ssmincidentmanagerincidents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ssmincidentmanagerincidents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResponsePlan_IncidentTemplatePropertyOutputReference interface {
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
	InternalValue() *AwsResponsePlan_IncidentTemplateProperty
	// Experimental.
	SetInternalValue(val *AwsResponsePlan_IncidentTemplateProperty)
	// Experimental.
	NotificationTarget() AwsResponsePlan_NotificationTargetPropertyList
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

// The jsii proxy struct for AwsResponsePlan_IncidentTemplatePropertyOutputReference
type jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) DedupeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedupeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) DedupeStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedupeStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) Impact() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ImpactInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) IncidentTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"incidentTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) IncidentTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"incidentTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) InternalValue() *AwsResponsePlan_IncidentTemplateProperty {
	var returns *AwsResponsePlan_IncidentTemplateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) NotificationTarget() AwsResponsePlan_NotificationTargetPropertyList {
	var returns AwsResponsePlan_NotificationTargetPropertyList
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) NotificationTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResponsePlan_IncidentTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsResponsePlan_IncidentTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResponsePlan_IncidentTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.AwsResponsePlan.IncidentTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResponsePlan_IncidentTemplatePropertyOutputReference_Override(a AwsResponsePlan_IncidentTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.AwsResponsePlan.IncidentTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetDedupeString(val *string) {
	if err := j.validateSetDedupeStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedupeString",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetImpact(val *float64) {
	if err := j.validateSetImpactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"impact",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetIncidentTags(val *map[string]*string) {
	if err := j.validateSetIncidentTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incidentTags",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetInternalValue(val *AwsResponsePlan_IncidentTemplateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) PutNotificationTarget(value interface{}) {
	if err := a.validatePutNotificationTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotificationTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ResetDedupeString() {
	_jsii_.InvokeVoid(
		a,
		"resetDedupeString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ResetIncidentTags() {
	_jsii_.InvokeVoid(
		a,
		"resetIncidentTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ResetNotificationTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ResetSummary() {
	_jsii_.InvokeVoid(
		a,
		"resetSummary",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResponsePlan_IncidentTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

