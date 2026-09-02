//go:build no_runtime_type_checking

package awsvpc

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfEndpoint_DnsEntryPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfEndpoint_DnsEntryPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfEndpoint_DnsEntryPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfEndpoint_DnsEntryPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfEndpoint_DnsEntryPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfEndpoint_DnsEntryPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfEndpoint_DnsEntryPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

