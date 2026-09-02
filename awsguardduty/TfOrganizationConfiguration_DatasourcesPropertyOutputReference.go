package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsguardduty/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsguardduty/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOrganizationConfiguration_DatasourcesPropertyOutputReference interface {
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
	InternalValue() *TfOrganizationConfiguration_DatasourcesProperty
	// Experimental.
	SetInternalValue(val *TfOrganizationConfiguration_DatasourcesProperty)
	// Experimental.
	Kubernetes() TfOrganizationConfiguration_KubernetesPropertyOutputReference
	// Experimental.
	KubernetesInput() *TfOrganizationConfiguration_KubernetesProperty
	// Experimental.
	MalwareProtection() TfOrganizationConfiguration_MalwareProtectionPropertyOutputReference
	// Experimental.
	MalwareProtectionInput() *TfOrganizationConfiguration_MalwareProtectionProperty
	// Experimental.
	S3Logs() TfOrganizationConfiguration_S3LogsPropertyOutputReference
	// Experimental.
	S3LogsInput() *TfOrganizationConfiguration_S3LogsProperty
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
	PutKubernetes(value *TfOrganizationConfiguration_KubernetesProperty)
	// Experimental.
	PutMalwareProtection(value *TfOrganizationConfiguration_MalwareProtectionProperty)
	// Experimental.
	PutS3Logs(value *TfOrganizationConfiguration_S3LogsProperty)
	// Experimental.
	ResetKubernetes()
	// Experimental.
	ResetMalwareProtection()
	// Experimental.
	ResetS3Logs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfOrganizationConfiguration_DatasourcesPropertyOutputReference
type jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) InternalValue() *TfOrganizationConfiguration_DatasourcesProperty {
	var returns *TfOrganizationConfiguration_DatasourcesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) Kubernetes() TfOrganizationConfiguration_KubernetesPropertyOutputReference {
	var returns TfOrganizationConfiguration_KubernetesPropertyOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) KubernetesInput() *TfOrganizationConfiguration_KubernetesProperty {
	var returns *TfOrganizationConfiguration_KubernetesProperty
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) MalwareProtection() TfOrganizationConfiguration_MalwareProtectionPropertyOutputReference {
	var returns TfOrganizationConfiguration_MalwareProtectionPropertyOutputReference
	_jsii_.Get(
		j,
		"malwareProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) MalwareProtectionInput() *TfOrganizationConfiguration_MalwareProtectionProperty {
	var returns *TfOrganizationConfiguration_MalwareProtectionProperty
	_jsii_.Get(
		j,
		"malwareProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) S3Logs() TfOrganizationConfiguration_S3LogsPropertyOutputReference {
	var returns TfOrganizationConfiguration_S3LogsPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) S3LogsInput() *TfOrganizationConfiguration_S3LogsProperty {
	var returns *TfOrganizationConfiguration_S3LogsProperty
	_jsii_.Get(
		j,
		"s3LogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOrganizationConfiguration_DatasourcesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfOrganizationConfiguration_DatasourcesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOrganizationConfiguration_DatasourcesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-guardduty.TfOrganizationConfiguration.DatasourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOrganizationConfiguration_DatasourcesPropertyOutputReference_Override(t TfOrganizationConfiguration_DatasourcesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-guardduty.TfOrganizationConfiguration.DatasourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference)SetInternalValue(val *TfOrganizationConfiguration_DatasourcesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) PutKubernetes(value *TfOrganizationConfiguration_KubernetesProperty) {
	if err := t.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) PutMalwareProtection(value *TfOrganizationConfiguration_MalwareProtectionProperty) {
	if err := t.validatePutMalwareProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMalwareProtection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) PutS3Logs(value *TfOrganizationConfiguration_S3LogsProperty) {
	if err := t.validatePutS3LogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Logs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) ResetKubernetes() {
	_jsii_.InvokeVoid(
		t,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) ResetMalwareProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetMalwareProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) ResetS3Logs() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Logs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOrganizationConfiguration_DatasourcesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

