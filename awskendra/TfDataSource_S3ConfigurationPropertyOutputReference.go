package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_S3ConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlListConfiguration() TfDataSource_AccessControlListConfigurationPropertyOutputReference
	// Experimental.
	AccessControlListConfigurationInput() *TfDataSource_AccessControlListConfigurationProperty
	// Experimental.
	BucketName() *string
	// Experimental.
	SetBucketName(val *string)
	// Experimental.
	BucketNameInput() *string
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
	DocumentsMetadataConfiguration() TfDataSource_DocumentsMetadataConfigurationPropertyOutputReference
	// Experimental.
	DocumentsMetadataConfigurationInput() *TfDataSource_DocumentsMetadataConfigurationProperty
	// Experimental.
	ExclusionPatterns() *[]*string
	// Experimental.
	SetExclusionPatterns(val *[]*string)
	// Experimental.
	ExclusionPatternsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InclusionPatterns() *[]*string
	// Experimental.
	SetInclusionPatterns(val *[]*string)
	// Experimental.
	InclusionPatternsInput() *[]*string
	// Experimental.
	InclusionPrefixes() *[]*string
	// Experimental.
	SetInclusionPrefixes(val *[]*string)
	// Experimental.
	InclusionPrefixesInput() *[]*string
	// Experimental.
	InternalValue() *TfDataSource_S3ConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDataSource_S3ConfigurationProperty)
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
	PutAccessControlListConfiguration(value *TfDataSource_AccessControlListConfigurationProperty)
	// Experimental.
	PutDocumentsMetadataConfiguration(value *TfDataSource_DocumentsMetadataConfigurationProperty)
	// Experimental.
	ResetAccessControlListConfiguration()
	// Experimental.
	ResetDocumentsMetadataConfiguration()
	// Experimental.
	ResetExclusionPatterns()
	// Experimental.
	ResetInclusionPatterns()
	// Experimental.
	ResetInclusionPrefixes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_S3ConfigurationPropertyOutputReference
type jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) AccessControlListConfiguration() TfDataSource_AccessControlListConfigurationPropertyOutputReference {
	var returns TfDataSource_AccessControlListConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlListConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) AccessControlListConfigurationInput() *TfDataSource_AccessControlListConfigurationProperty {
	var returns *TfDataSource_AccessControlListConfigurationProperty
	_jsii_.Get(
		j,
		"accessControlListConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) DocumentsMetadataConfiguration() TfDataSource_DocumentsMetadataConfigurationPropertyOutputReference {
	var returns TfDataSource_DocumentsMetadataConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"documentsMetadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) DocumentsMetadataConfigurationInput() *TfDataSource_DocumentsMetadataConfigurationProperty {
	var returns *TfDataSource_DocumentsMetadataConfigurationProperty
	_jsii_.Get(
		j,
		"documentsMetadataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ExclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) InclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) InclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) InclusionPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) InclusionPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) InternalValue() *TfDataSource_S3ConfigurationProperty {
	var returns *TfDataSource_S3ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_S3ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataSource_S3ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_S3ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.S3ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_S3ConfigurationPropertyOutputReference_Override(t TfDataSource_S3ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.S3ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetExclusionPatterns(val *[]*string) {
	if err := j.validateSetExclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetInclusionPatterns(val *[]*string) {
	if err := j.validateSetInclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetInclusionPrefixes(val *[]*string) {
	if err := j.validateSetInclusionPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionPrefixes",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetInternalValue(val *TfDataSource_S3ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) PutAccessControlListConfiguration(value *TfDataSource_AccessControlListConfigurationProperty) {
	if err := t.validatePutAccessControlListConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessControlListConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) PutDocumentsMetadataConfiguration(value *TfDataSource_DocumentsMetadataConfigurationProperty) {
	if err := t.validatePutDocumentsMetadataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDocumentsMetadataConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ResetAccessControlListConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessControlListConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ResetDocumentsMetadataConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentsMetadataConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ResetExclusionPatterns() {
	_jsii_.InvokeVoid(
		t,
		"resetExclusionPatterns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ResetInclusionPatterns() {
	_jsii_.InvokeVoid(
		t,
		"resetInclusionPatterns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ResetInclusionPrefixes() {
	_jsii_.InvokeVoid(
		t,
		"resetInclusionPrefixes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_S3ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

