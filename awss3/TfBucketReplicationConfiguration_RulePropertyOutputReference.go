package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucketReplicationConfiguration_RulePropertyOutputReference interface {
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
	DeleteMarkerReplication() TfBucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference
	// Experimental.
	DeleteMarkerReplicationInput() *TfBucketReplicationConfiguration_DeleteMarkerReplicationProperty
	// Experimental.
	Destination() TfBucketReplicationConfiguration_DestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *TfBucketReplicationConfiguration_DestinationProperty
	// Experimental.
	ExistingObjectReplication() TfBucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference
	// Experimental.
	ExistingObjectReplicationInput() *TfBucketReplicationConfiguration_ExistingObjectReplicationProperty
	// Experimental.
	Filter() TfBucketReplicationConfiguration_FilterPropertyOutputReference
	// Experimental.
	FilterInput() *TfBucketReplicationConfiguration_FilterProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
	// Experimental.
	SourceSelectionCriteria() TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
	// Experimental.
	SourceSelectionCriteriaInput() *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty
	// Experimental.
	Status() *string
	// Experimental.
	SetStatus(val *string)
	// Experimental.
	StatusInput() *string
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
	PutDeleteMarkerReplication(value *TfBucketReplicationConfiguration_DeleteMarkerReplicationProperty)
	// Experimental.
	PutDestination(value *TfBucketReplicationConfiguration_DestinationProperty)
	// Experimental.
	PutExistingObjectReplication(value *TfBucketReplicationConfiguration_ExistingObjectReplicationProperty)
	// Experimental.
	PutFilter(value *TfBucketReplicationConfiguration_FilterProperty)
	// Experimental.
	PutSourceSelectionCriteria(value *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty)
	// Experimental.
	ResetDeleteMarkerReplication()
	// Experimental.
	ResetExistingObjectReplication()
	// Experimental.
	ResetFilter()
	// Experimental.
	ResetId()
	// Experimental.
	ResetPrefix()
	// Experimental.
	ResetPriority()
	// Experimental.
	ResetSourceSelectionCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfBucketReplicationConfiguration_RulePropertyOutputReference
type jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) DeleteMarkerReplication() TfBucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"deleteMarkerReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) DeleteMarkerReplicationInput() *TfBucketReplicationConfiguration_DeleteMarkerReplicationProperty {
	var returns *TfBucketReplicationConfiguration_DeleteMarkerReplicationProperty
	_jsii_.Get(
		j,
		"deleteMarkerReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Destination() TfBucketReplicationConfiguration_DestinationPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_DestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) DestinationInput() *TfBucketReplicationConfiguration_DestinationProperty {
	var returns *TfBucketReplicationConfiguration_DestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ExistingObjectReplication() TfBucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"existingObjectReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ExistingObjectReplicationInput() *TfBucketReplicationConfiguration_ExistingObjectReplicationProperty {
	var returns *TfBucketReplicationConfiguration_ExistingObjectReplicationProperty
	_jsii_.Get(
		j,
		"existingObjectReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Filter() TfBucketReplicationConfiguration_FilterPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_FilterPropertyOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) FilterInput() *TfBucketReplicationConfiguration_FilterProperty {
	var returns *TfBucketReplicationConfiguration_FilterProperty
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) SourceSelectionCriteria() TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceSelectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) SourceSelectionCriteriaInput() *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty {
	var returns *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty
	_jsii_.Get(
		j,
		"sourceSelectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucketReplicationConfiguration_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfBucketReplicationConfiguration_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucketReplicationConfiguration_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketReplicationConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucketReplicationConfiguration_RulePropertyOutputReference_Override(t TfBucketReplicationConfiguration_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketReplicationConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) PutDeleteMarkerReplication(value *TfBucketReplicationConfiguration_DeleteMarkerReplicationProperty) {
	if err := t.validatePutDeleteMarkerReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeleteMarkerReplication",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) PutDestination(value *TfBucketReplicationConfiguration_DestinationProperty) {
	if err := t.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) PutExistingObjectReplication(value *TfBucketReplicationConfiguration_ExistingObjectReplicationProperty) {
	if err := t.validatePutExistingObjectReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExistingObjectReplication",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) PutFilter(value *TfBucketReplicationConfiguration_FilterProperty) {
	if err := t.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) PutSourceSelectionCriteria(value *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty) {
	if err := t.validatePutSourceSelectionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceSelectionCriteria",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ResetDeleteMarkerReplication() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteMarkerReplication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ResetExistingObjectReplication() {
	_jsii_.InvokeVoid(
		t,
		"resetExistingObjectReplication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		t,
		"resetPriority",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ResetSourceSelectionCriteria() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceSelectionCriteria",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

