package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBucketReplicationConfiguration_RulePropertyOutputReference interface {
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
	DeleteMarkerReplication() AwsBucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference
	// Experimental.
	DeleteMarkerReplicationInput() *AwsBucketReplicationConfiguration_DeleteMarkerReplicationProperty
	// Experimental.
	Destination() AwsBucketReplicationConfiguration_DestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *AwsBucketReplicationConfiguration_DestinationProperty
	// Experimental.
	ExistingObjectReplication() AwsBucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference
	// Experimental.
	ExistingObjectReplicationInput() *AwsBucketReplicationConfiguration_ExistingObjectReplicationProperty
	// Experimental.
	Filter() AwsBucketReplicationConfiguration_FilterPropertyOutputReference
	// Experimental.
	FilterInput() *AwsBucketReplicationConfiguration_FilterProperty
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
	SourceSelectionCriteria() AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
	// Experimental.
	SourceSelectionCriteriaInput() *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty
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
	PutDeleteMarkerReplication(value *AwsBucketReplicationConfiguration_DeleteMarkerReplicationProperty)
	// Experimental.
	PutDestination(value *AwsBucketReplicationConfiguration_DestinationProperty)
	// Experimental.
	PutExistingObjectReplication(value *AwsBucketReplicationConfiguration_ExistingObjectReplicationProperty)
	// Experimental.
	PutFilter(value *AwsBucketReplicationConfiguration_FilterProperty)
	// Experimental.
	PutSourceSelectionCriteria(value *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty)
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

// The jsii proxy struct for AwsBucketReplicationConfiguration_RulePropertyOutputReference
type jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) DeleteMarkerReplication() AwsBucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference {
	var returns AwsBucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"deleteMarkerReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) DeleteMarkerReplicationInput() *AwsBucketReplicationConfiguration_DeleteMarkerReplicationProperty {
	var returns *AwsBucketReplicationConfiguration_DeleteMarkerReplicationProperty
	_jsii_.Get(
		j,
		"deleteMarkerReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Destination() AwsBucketReplicationConfiguration_DestinationPropertyOutputReference {
	var returns AwsBucketReplicationConfiguration_DestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) DestinationInput() *AwsBucketReplicationConfiguration_DestinationProperty {
	var returns *AwsBucketReplicationConfiguration_DestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ExistingObjectReplication() AwsBucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference {
	var returns AwsBucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"existingObjectReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ExistingObjectReplicationInput() *AwsBucketReplicationConfiguration_ExistingObjectReplicationProperty {
	var returns *AwsBucketReplicationConfiguration_ExistingObjectReplicationProperty
	_jsii_.Get(
		j,
		"existingObjectReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Filter() AwsBucketReplicationConfiguration_FilterPropertyOutputReference {
	var returns AwsBucketReplicationConfiguration_FilterPropertyOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) FilterInput() *AwsBucketReplicationConfiguration_FilterProperty {
	var returns *AwsBucketReplicationConfiguration_FilterProperty
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) SourceSelectionCriteria() AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference {
	var returns AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceSelectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) SourceSelectionCriteriaInput() *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty {
	var returns *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty
	_jsii_.Get(
		j,
		"sourceSelectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBucketReplicationConfiguration_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBucketReplicationConfiguration_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBucketReplicationConfiguration_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketReplicationConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBucketReplicationConfiguration_RulePropertyOutputReference_Override(a AwsBucketReplicationConfiguration_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketReplicationConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) PutDeleteMarkerReplication(value *AwsBucketReplicationConfiguration_DeleteMarkerReplicationProperty) {
	if err := a.validatePutDeleteMarkerReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeleteMarkerReplication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) PutDestination(value *AwsBucketReplicationConfiguration_DestinationProperty) {
	if err := a.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) PutExistingObjectReplication(value *AwsBucketReplicationConfiguration_ExistingObjectReplicationProperty) {
	if err := a.validatePutExistingObjectReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExistingObjectReplication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) PutFilter(value *AwsBucketReplicationConfiguration_FilterProperty) {
	if err := a.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) PutSourceSelectionCriteria(value *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty) {
	if err := a.validatePutSourceSelectionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceSelectionCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ResetDeleteMarkerReplication() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteMarkerReplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ResetExistingObjectReplication() {
	_jsii_.InvokeVoid(
		a,
		"resetExistingObjectReplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ResetSourceSelectionCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceSelectionCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

