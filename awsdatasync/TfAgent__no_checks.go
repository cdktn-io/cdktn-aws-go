//go:build no_runtime_type_checking

package awsdatasync

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfAgent) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateMoveToIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validatePutTimeoutsParameters(value *TfAgent_TimeoutsProperty) error {
	return nil
}

func (t *jsiiProxy_TfAgent) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateTfAgent_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateTfAgent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTfAgent_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTfAgent_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetActivationKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetIpAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetPrivateLinkEndpointParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetRegionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetSecurityGroupArnsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetSubnetArnsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetTagsAllParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_TfAgent) validateSetVpcEndpointIdParameters(val *string) error {
	return nil
}

func validateNewTfAgentParameters(scope constructs.Construct, id *string, config *TfAgentConfig) error {
	return nil
}

