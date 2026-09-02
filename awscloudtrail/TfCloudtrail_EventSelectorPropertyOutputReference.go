package awscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudtrail/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudtrail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCloudtrail_EventSelectorPropertyOutputReference interface {
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
	DataResource() TfCloudtrail_DataResourcePropertyList
	// Experimental.
	DataResourceInput() interface{}
	// Experimental.
	ExcludeManagementEventSources() *[]*string
	// Experimental.
	SetExcludeManagementEventSources(val *[]*string)
	// Experimental.
	ExcludeManagementEventSourcesInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeManagementEvents() interface{}
	// Experimental.
	SetIncludeManagementEvents(val interface{})
	// Experimental.
	IncludeManagementEventsInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ReadWriteType() *string
	// Experimental.
	SetReadWriteType(val *string)
	// Experimental.
	ReadWriteTypeInput() *string
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
	PutDataResource(value interface{})
	// Experimental.
	ResetDataResource()
	// Experimental.
	ResetExcludeManagementEventSources()
	// Experimental.
	ResetIncludeManagementEvents()
	// Experimental.
	ResetReadWriteType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCloudtrail_EventSelectorPropertyOutputReference
type jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) DataResource() TfCloudtrail_DataResourcePropertyList {
	var returns TfCloudtrail_DataResourcePropertyList
	_jsii_.Get(
		j,
		"dataResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) DataResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ExcludeManagementEventSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeManagementEventSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ExcludeManagementEventSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeManagementEventSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) IncludeManagementEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeManagementEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) IncludeManagementEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeManagementEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ReadWriteType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readWriteType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ReadWriteTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readWriteTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCloudtrail_EventSelectorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCloudtrail_EventSelectorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCloudtrail_EventSelectorPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudtrail.TfCloudtrail.EventSelectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCloudtrail_EventSelectorPropertyOutputReference_Override(t TfCloudtrail_EventSelectorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudtrail.TfCloudtrail.EventSelectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetExcludeManagementEventSources(val *[]*string) {
	if err := j.validateSetExcludeManagementEventSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeManagementEventSources",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetIncludeManagementEvents(val interface{}) {
	if err := j.validateSetIncludeManagementEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeManagementEvents",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetReadWriteType(val *string) {
	if err := j.validateSetReadWriteTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readWriteType",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) PutDataResource(value interface{}) {
	if err := t.validatePutDataResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ResetDataResource() {
	_jsii_.InvokeVoid(
		t,
		"resetDataResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ResetExcludeManagementEventSources() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludeManagementEventSources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ResetIncludeManagementEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeManagementEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ResetReadWriteType() {
	_jsii_.InvokeVoid(
		t,
		"resetReadWriteType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCloudtrail_EventSelectorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

