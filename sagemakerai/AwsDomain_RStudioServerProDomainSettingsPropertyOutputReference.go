package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference interface {
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
	DefaultResourceSpec() AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty
	// Experimental.
	DomainExecutionRoleArn() *string
	// Experimental.
	SetDomainExecutionRoleArn(val *string)
	// Experimental.
	DomainExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDomain_RStudioServerProDomainSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_RStudioServerProDomainSettingsProperty)
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
	PutDefaultResourceSpec(value *AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty)
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

// The jsii proxy struct for AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference
type jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) DefaultResourceSpec() AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecPropertyOutputReference {
	var returns AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) DefaultResourceSpecInput() *AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty {
	var returns *AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) DomainExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) DomainExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) InternalValue() *AwsDomain_RStudioServerProDomainSettingsProperty {
	var returns *AwsDomain_RStudioServerProDomainSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioConnectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioConnectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioConnectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioConnectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioPackageManagerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioPackageManagerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) RStudioPackageManagerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioPackageManagerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_RStudioServerProDomainSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_RStudioServerProDomainSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.RStudioServerProDomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_RStudioServerProDomainSettingsPropertyOutputReference_Override(a AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.RStudioServerProDomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetDomainExecutionRoleArn(val *string) {
	if err := j.validateSetDomainExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetInternalValue(val *AwsDomain_RStudioServerProDomainSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetRStudioConnectUrl(val *string) {
	if err := j.validateSetRStudioConnectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rStudioConnectUrl",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetRStudioPackageManagerUrl(val *string) {
	if err := j.validateSetRStudioPackageManagerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rStudioPackageManagerUrl",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) PutDefaultResourceSpec(value *AwsDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty) {
	if err := a.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) ResetRStudioConnectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRStudioConnectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) ResetRStudioPackageManagerUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRStudioPackageManagerUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_RStudioServerProDomainSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

