package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3BucketReplicationConfiguration_RulePropertyOutputReference interface {
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
	DeleteMarkerReplication() AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference
	// Experimental.
	DeleteMarkerReplicationInput() *AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationProperty
	// Experimental.
	Destination() AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *AwsS3BucketReplicationConfiguration_DestinationProperty
	// Experimental.
	ExistingObjectReplication() AwsS3BucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference
	// Experimental.
	ExistingObjectReplicationInput() *AwsS3BucketReplicationConfiguration_ExistingObjectReplicationProperty
	// Experimental.
	Filter() AwsS3BucketReplicationConfiguration_FilterPropertyOutputReference
	// Experimental.
	FilterInput() *AwsS3BucketReplicationConfiguration_FilterProperty
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
	SourceSelectionCriteria() AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
	// Experimental.
	SourceSelectionCriteriaInput() *AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaProperty
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
	PutDeleteMarkerReplication(value *AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationProperty)
	// Experimental.
	PutDestination(value *AwsS3BucketReplicationConfiguration_DestinationProperty)
	// Experimental.
	PutExistingObjectReplication(value *AwsS3BucketReplicationConfiguration_ExistingObjectReplicationProperty)
	// Experimental.
	PutFilter(value *AwsS3BucketReplicationConfiguration_FilterProperty)
	// Experimental.
	PutSourceSelectionCriteria(value *AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaProperty)
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

// The jsii proxy struct for AwsS3BucketReplicationConfiguration_RulePropertyOutputReference
type jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) DeleteMarkerReplication() AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"deleteMarkerReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) DeleteMarkerReplicationInput() *AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationProperty {
	var returns *AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationProperty
	_jsii_.Get(
		j,
		"deleteMarkerReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Destination() AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) DestinationInput() *AwsS3BucketReplicationConfiguration_DestinationProperty {
	var returns *AwsS3BucketReplicationConfiguration_DestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ExistingObjectReplication() AwsS3BucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_ExistingObjectReplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"existingObjectReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ExistingObjectReplicationInput() *AwsS3BucketReplicationConfiguration_ExistingObjectReplicationProperty {
	var returns *AwsS3BucketReplicationConfiguration_ExistingObjectReplicationProperty
	_jsii_.Get(
		j,
		"existingObjectReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Filter() AwsS3BucketReplicationConfiguration_FilterPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_FilterPropertyOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) FilterInput() *AwsS3BucketReplicationConfiguration_FilterProperty {
	var returns *AwsS3BucketReplicationConfiguration_FilterProperty
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) SourceSelectionCriteria() AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceSelectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) SourceSelectionCriteriaInput() *AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaProperty {
	var returns *AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaProperty
	_jsii_.Get(
		j,
		"sourceSelectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3BucketReplicationConfiguration_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsS3BucketReplicationConfiguration_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3BucketReplicationConfiguration_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3BucketReplicationConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3BucketReplicationConfiguration_RulePropertyOutputReference_Override(a AwsS3BucketReplicationConfiguration_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3BucketReplicationConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) PutDeleteMarkerReplication(value *AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationProperty) {
	if err := a.validatePutDeleteMarkerReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeleteMarkerReplication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) PutDestination(value *AwsS3BucketReplicationConfiguration_DestinationProperty) {
	if err := a.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) PutExistingObjectReplication(value *AwsS3BucketReplicationConfiguration_ExistingObjectReplicationProperty) {
	if err := a.validatePutExistingObjectReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExistingObjectReplication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) PutFilter(value *AwsS3BucketReplicationConfiguration_FilterProperty) {
	if err := a.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) PutSourceSelectionCriteria(value *AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaProperty) {
	if err := a.validatePutSourceSelectionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceSelectionCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ResetDeleteMarkerReplication() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteMarkerReplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ResetExistingObjectReplication() {
	_jsii_.InvokeVoid(
		a,
		"resetExistingObjectReplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ResetSourceSelectionCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceSelectionCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

