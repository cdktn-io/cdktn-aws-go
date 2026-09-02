package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty
	// Experimental.
	SetInternalValue(val *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty)
	// Experimental.
	ReplicaModifications() TfBucketReplicationConfiguration_ReplicaModificationsPropertyOutputReference
	// Experimental.
	ReplicaModificationsInput() *TfBucketReplicationConfiguration_ReplicaModificationsProperty
	// Experimental.
	SseKmsEncryptedObjects() TfBucketReplicationConfiguration_SseKmsEncryptedObjectsPropertyOutputReference
	// Experimental.
	SseKmsEncryptedObjectsInput() *TfBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty
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
	PutReplicaModifications(value *TfBucketReplicationConfiguration_ReplicaModificationsProperty)
	// Experimental.
	PutSseKmsEncryptedObjects(value *TfBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty)
	// Experimental.
	ResetReplicaModifications()
	// Experimental.
	ResetSseKmsEncryptedObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
type jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) InternalValue() *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty {
	var returns *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ReplicaModifications() TfBucketReplicationConfiguration_ReplicaModificationsPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_ReplicaModificationsPropertyOutputReference
	_jsii_.Get(
		j,
		"replicaModifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ReplicaModificationsInput() *TfBucketReplicationConfiguration_ReplicaModificationsProperty {
	var returns *TfBucketReplicationConfiguration_ReplicaModificationsProperty
	_jsii_.Get(
		j,
		"replicaModificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) SseKmsEncryptedObjects() TfBucketReplicationConfiguration_SseKmsEncryptedObjectsPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_SseKmsEncryptedObjectsPropertyOutputReference
	_jsii_.Get(
		j,
		"sseKmsEncryptedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) SseKmsEncryptedObjectsInput() *TfBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty {
	var returns *TfBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty
	_jsii_.Get(
		j,
		"sseKmsEncryptedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketReplicationConfiguration.SourceSelectionCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference_Override(t TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketReplicationConfiguration.SourceSelectionCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetInternalValue(val *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) PutReplicaModifications(value *TfBucketReplicationConfiguration_ReplicaModificationsProperty) {
	if err := t.validatePutReplicaModificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReplicaModifications",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) PutSseKmsEncryptedObjects(value *TfBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty) {
	if err := t.validatePutSseKmsEncryptedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSseKmsEncryptedObjects",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ResetReplicaModifications() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicaModifications",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ResetSseKmsEncryptedObjects() {
	_jsii_.InvokeVoid(
		t,
		"resetSseKmsEncryptedObjects",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

