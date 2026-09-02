package awsdataexchange

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdataexchange/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdataexchange/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bucket() *string
	// Experimental.
	SetBucket(val *string)
	// Experimental.
	BucketInput() *string
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KeyPrefixes() *[]*string
	// Experimental.
	SetKeyPrefixes(val *[]*string)
	// Experimental.
	KeyPrefixesInput() *[]*string
	// Experimental.
	Keys() *[]*string
	// Experimental.
	SetKeys(val *[]*string)
	// Experimental.
	KeysInput() *[]*string
	// Experimental.
	KmsKeysToGrant() TfRevisionAssets_KmsKeysToGrantPropertyList
	// Experimental.
	KmsKeysToGrantInput() interface{}
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
	PutKmsKeysToGrant(value interface{})
	// Experimental.
	ResetKeyPrefixes()
	// Experimental.
	ResetKeys()
	// Experimental.
	ResetKmsKeysToGrant()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference
type jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) KeyPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) KeyPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) Keys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) KeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) KmsKeysToGrant() TfRevisionAssets_KmsKeysToGrantPropertyList {
	var returns TfRevisionAssets_KmsKeysToGrantPropertyList
	_jsii_.Get(
		j,
		"kmsKeysToGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) KmsKeysToGrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsKeysToGrantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-data-exchange.TfRevisionAssets.AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference_Override(t TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-data-exchange.TfRevisionAssets.AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetKeyPrefixes(val *[]*string) {
	if err := j.validateSetKeyPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPrefixes",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetKeys(val *[]*string) {
	if err := j.validateSetKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keys",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) PutKmsKeysToGrant(value interface{}) {
	if err := t.validatePutKmsKeysToGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKmsKeysToGrant",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) ResetKeyPrefixes() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyPrefixes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) ResetKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) ResetKmsKeysToGrant() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeysToGrant",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

