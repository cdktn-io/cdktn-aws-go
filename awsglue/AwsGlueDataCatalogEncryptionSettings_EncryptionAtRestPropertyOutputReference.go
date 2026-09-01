package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CatalogEncryptionMode() *string
	// Experimental.
	SetCatalogEncryptionMode(val *string)
	// Experimental.
	CatalogEncryptionModeInput() *string
	// Experimental.
	CatalogEncryptionServiceRole() *string
	// Experimental.
	SetCatalogEncryptionServiceRole(val *string)
	// Experimental.
	CatalogEncryptionServiceRoleInput() *string
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
	InternalValue() *AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestProperty
	// Experimental.
	SetInternalValue(val *AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestProperty)
	// Experimental.
	SseAwsKmsKeyId() *string
	// Experimental.
	SetSseAwsKmsKeyId(val *string)
	// Experimental.
	SseAwsKmsKeyIdInput() *string
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
	ResetCatalogEncryptionServiceRole()
	// Experimental.
	ResetSseAwsKmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference
type jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionServiceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionServiceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) InternalValue() *AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestProperty {
	var returns *AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) SseAwsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAwsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) SseAwsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAwsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueDataCatalogEncryptionSettings.EncryptionAtRestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference_Override(a AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueDataCatalogEncryptionSettings.EncryptionAtRestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetCatalogEncryptionMode(val *string) {
	if err := j.validateSetCatalogEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetCatalogEncryptionServiceRole(val *string) {
	if err := j.validateSetCatalogEncryptionServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogEncryptionServiceRole",
		val,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetInternalValue(val *AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetSseAwsKmsKeyId(val *string) {
	if err := j.validateSetSseAwsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sseAwsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ResetCatalogEncryptionServiceRole() {
	_jsii_.InvokeVoid(
		a,
		"resetCatalogEncryptionServiceRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ResetSseAwsKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetSseAwsKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

