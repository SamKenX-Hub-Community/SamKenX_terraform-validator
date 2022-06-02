// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func certManagerDefaultScopeDiffSuppress(_, old, new string, diff *schema.ResourceData) bool {
	if old == "" && new == "DEFAULT" || old == "DEFAULT" && new == "" {
		return true
	}
	return false
}

const CertificateManagerCertificateAssetType string = "certificatemanager.googleapis.com/Certificate"

func resourceConverterCertificateManagerCertificate() ResourceConverter {
	return ResourceConverter{
		AssetType: CertificateManagerCertificateAssetType,
		Convert:   GetCertificateManagerCertificateCaiObject,
	}
}

func GetCertificateManagerCertificateCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//certificatemanager.googleapis.com/projects/{{project}}/locations/global/certificates/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCertificateManagerCertificateApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: CertificateManagerCertificateAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/certificatemanager/v1/rest",
				DiscoveryName:        "Certificate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCertificateManagerCertificateApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerCertificateLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	scopeProp, err := expandCertificateManagerCertificateScope(d.Get("scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scope"); !isEmptyValue(reflect.ValueOf(scopeProp)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}
	selfManagedProp, err := expandCertificateManagerCertificateSelfManaged(d.Get("self_managed"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("self_managed"); !isEmptyValue(reflect.ValueOf(selfManagedProp)) && (ok || !reflect.DeepEqual(v, selfManagedProp)) {
		obj["selfManaged"] = selfManagedProp
	}
	managedProp, err := expandCertificateManagerCertificateManaged(d.Get("managed"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("managed"); !isEmptyValue(reflect.ValueOf(managedProp)) && (ok || !reflect.DeepEqual(v, managedProp)) {
		obj["managed"] = managedProp
	}

	return obj, nil
}

func expandCertificateManagerCertificateDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerCertificateScope(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManaged(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificatePem, err := expandCertificateManagerCertificateSelfManagedCertificatePem(original["certificate_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificatePem); val.IsValid() && !isEmptyValue(val) {
		transformed["certificatePem"] = transformedCertificatePem
	}

	transformedPrivateKeyPem, err := expandCertificateManagerCertificateSelfManagedPrivateKeyPem(original["private_key_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrivateKeyPem); val.IsValid() && !isEmptyValue(val) {
		transformed["privateKeyPem"] = transformedPrivateKeyPem
	}

	return transformed, nil
}

func expandCertificateManagerCertificateSelfManagedCertificatePem(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManagedPrivateKeyPem(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManaged(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedState, err := expandCertificateManagerCertificateManagedState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedDomains, err := expandCertificateManagerCertificateManagedDomains(original["domains"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDomains); val.IsValid() && !isEmptyValue(val) {
		transformed["domains"] = transformedDomains
	}

	transformedDnsAuthorizations, err := expandCertificateManagerCertificateManagedDnsAuthorizations(original["dns_authorizations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsAuthorizations); val.IsValid() && !isEmptyValue(val) {
		transformed["dnsAuthorizations"] = transformedDnsAuthorizations
	}

	return transformed, nil
}

func expandCertificateManagerCertificateManagedState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedDomains(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedDnsAuthorizations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
