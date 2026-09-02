package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_RStudioServerProDomainSettingsPropertyOutputReference interface {
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
	DefaultResourceSpec() TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty
	// Experimental.
	DomainExecutionRoleArn() *string
	// Experimental.
	SetDomainExecutionRoleArn(val *string)
	// Experimental.
	DomainExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_RStudioServerProDomainSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_RStudioServerProDomainSettingsProperty)
	// Experimental.
	RStudioConnectUrl() *string
	// Experimental.
	SetRStudioConnectUrl(val *string)
	// Experimental.
	RStudioConnectUrlInput() *string
	// Experimental.
	RStudioPackageManagerUrl() *string
	// Experimental.
	SetRStudioPackageManagerUrl(val *string)
	// Experimental.
	RStudioPackageManagerUrlInput() *string
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
	PutDefaultResourceSpec(value *TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty)
	// Experimental.
	ResetDefaultResourceSpec()
	// Experimental.
	ResetRStudioConnectUrl()
	// Experimental.
	ResetRStudioPackageManagerUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_RStudioServerProDomainSettingsPropertyOutputReference
type jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) DefaultResourceSpec() TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecPropertyOutputReference {
	var returns TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) DefaultResourceSpecInput() *TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty {
	var returns *TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) DomainExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) DomainExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) InternalValue() *TfDomain_RStudioServerProDomainSettingsProperty {
	var returns *TfDomain_RStudioServerProDomainSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioConnectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioConnectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioConnectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioConnectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioPackageManagerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioPackageManagerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioPackageManagerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioPackageManagerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_RStudioServerProDomainSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_RStudioServerProDomainSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_RStudioServerProDomainSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.RStudioServerProDomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_RStudioServerProDomainSettingsPropertyOutputReference_Override(t TfDomain_RStudioServerProDomainSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.RStudioServerProDomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetDomainExecutionRoleArn(val *string) {
	if err := j.validateSetDomainExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_RStudioServerProDomainSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetRStudioConnectUrl(val *string) {
	if err := j.validateSetRStudioConnectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rStudioConnectUrl",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetRStudioPackageManagerUrl(val *string) {
	if err := j.validateSetRStudioPackageManagerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rStudioPackageManagerUrl",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) PutDefaultResourceSpec(value *TfDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty) {
	if err := t.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) ResetRStudioConnectUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetRStudioConnectUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) ResetRStudioPackageManagerUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetRStudioPackageManagerUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_RStudioServerProDomainSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

