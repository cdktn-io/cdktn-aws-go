package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicator_ConsumerGroupReplicationPropertyOutputReference interface {
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
	// Experimental.
	ConsumerGroupOffsetSyncMode() *string
	// Experimental.
	SetConsumerGroupOffsetSyncMode(val *string)
	// Experimental.
	ConsumerGroupOffsetSyncModeInput() *string
	// Experimental.
	ConsumerGroupsToExclude() *[]*string
	// Experimental.
	SetConsumerGroupsToExclude(val *[]*string)
	// Experimental.
	ConsumerGroupsToExcludeInput() *[]*string
	// Experimental.
	ConsumerGroupsToReplicate() *[]*string
	// Experimental.
	SetConsumerGroupsToReplicate(val *[]*string)
	// Experimental.
	ConsumerGroupsToReplicateInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DetectAndCopyNewConsumerGroups() interface{}
	// Experimental.
	SetDetectAndCopyNewConsumerGroups(val interface{})
	// Experimental.
	DetectAndCopyNewConsumerGroupsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SynchroniseConsumerGroupOffsets() interface{}
	// Experimental.
	SetSynchroniseConsumerGroupOffsets(val interface{})
	// Experimental.
	SynchroniseConsumerGroupOffsetsInput() interface{}
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
	ResetConsumerGroupOffsetSyncMode()
	// Experimental.
	ResetConsumerGroupsToExclude()
	// Experimental.
	ResetDetectAndCopyNewConsumerGroups()
	// Experimental.
	ResetSynchroniseConsumerGroupOffsets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfReplicator_ConsumerGroupReplicationPropertyOutputReference
type jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupOffsetSyncMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupOffsetSyncMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupOffsetSyncModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupOffsetSyncModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToReplicate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToReplicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToReplicateInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToReplicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) DetectAndCopyNewConsumerGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewConsumerGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) DetectAndCopyNewConsumerGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewConsumerGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) SynchroniseConsumerGroupOffsets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synchroniseConsumerGroupOffsets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) SynchroniseConsumerGroupOffsetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synchroniseConsumerGroupOffsetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfReplicator_ConsumerGroupReplicationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfReplicator_ConsumerGroupReplicationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfReplicator_ConsumerGroupReplicationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.TfReplicator.ConsumerGroupReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfReplicator_ConsumerGroupReplicationPropertyOutputReference_Override(t TfReplicator_ConsumerGroupReplicationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.TfReplicator.ConsumerGroupReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetConsumerGroupOffsetSyncMode(val *string) {
	if err := j.validateSetConsumerGroupOffsetSyncModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupOffsetSyncMode",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetConsumerGroupsToExclude(val *[]*string) {
	if err := j.validateSetConsumerGroupsToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupsToExclude",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetConsumerGroupsToReplicate(val *[]*string) {
	if err := j.validateSetConsumerGroupsToReplicateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupsToReplicate",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetDetectAndCopyNewConsumerGroups(val interface{}) {
	if err := j.validateSetDetectAndCopyNewConsumerGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectAndCopyNewConsumerGroups",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetSynchroniseConsumerGroupOffsets(val interface{}) {
	if err := j.validateSetSynchroniseConsumerGroupOffsetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synchroniseConsumerGroupOffsets",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetConsumerGroupOffsetSyncMode() {
	_jsii_.InvokeVoid(
		t,
		"resetConsumerGroupOffsetSyncMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetConsumerGroupsToExclude() {
	_jsii_.InvokeVoid(
		t,
		"resetConsumerGroupsToExclude",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetDetectAndCopyNewConsumerGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetDetectAndCopyNewConsumerGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetSynchroniseConsumerGroupOffsets() {
	_jsii_.InvokeVoid(
		t,
		"resetSynchroniseConsumerGroupOffsets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfReplicator_ConsumerGroupReplicationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

