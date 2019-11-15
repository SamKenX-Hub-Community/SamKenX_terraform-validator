// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"fmt"
	"log"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBinaryAuthorizationAttestor() *schema.Resource {
	return &schema.Resource{
		Create: resourceBinaryAuthorizationAttestorCreate,
		Read:   resourceBinaryAuthorizationAttestorRead,
		Update: resourceBinaryAuthorizationAttestorUpdate,
		Delete: resourceBinaryAuthorizationAttestorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBinaryAuthorizationAttestorImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"attestation_authority_note": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"note_reference": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"public_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ascii_armored_pgp_public_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"comment": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
									},
									"pkix_public_key": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"public_key_pem": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"signature_algorithm": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"delegation_service_account_email": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBinaryAuthorizationAttestorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandBinaryAuthorizationAttestorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandBinaryAuthorizationAttestorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	userOwnedDrydockNoteProp, err := expandBinaryAuthorizationAttestorAttestationAuthorityNote(d.Get("attestation_authority_note"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority_note"); !isEmptyValue(reflect.ValueOf(userOwnedDrydockNoteProp)) && (ok || !reflect.DeepEqual(v, userOwnedDrydockNoteProp)) {
		obj["userOwnedDrydockNote"] = userOwnedDrydockNoteProp
	}

	obj, err = resourceBinaryAuthorizationAttestorEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors?attestorId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Attestor: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Attestor: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Attestor %q: %#v", d.Id(), res)

	return resourceBinaryAuthorizationAttestorRead(d, meta)
}

func resourceBinaryAuthorizationAttestorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BinaryAuthorizationAttestor %q", d.Id()))
	}

	res, err = resourceBinaryAuthorizationAttestorDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing BinaryAuthorizationAttestor because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}

	if err := d.Set("name", flattenBinaryAuthorizationAttestorName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}
	if err := d.Set("description", flattenBinaryAuthorizationAttestorDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}
	if err := d.Set("attestation_authority_note", flattenBinaryAuthorizationAttestorAttestationAuthorityNote(res["userOwnedDrydockNote"], d)); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}

	return nil
}

func resourceBinaryAuthorizationAttestorUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandBinaryAuthorizationAttestorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandBinaryAuthorizationAttestorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	userOwnedDrydockNoteProp, err := expandBinaryAuthorizationAttestorAttestationAuthorityNote(d.Get("attestation_authority_note"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority_note"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userOwnedDrydockNoteProp)) {
		obj["userOwnedDrydockNote"] = userOwnedDrydockNoteProp
	}

	obj, err = resourceBinaryAuthorizationAttestorEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Attestor %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Attestor %q: %s", d.Id(), err)
	}

	return resourceBinaryAuthorizationAttestorRead(d, meta)
}

func resourceBinaryAuthorizationAttestorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Attestor %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Attestor")
	}

	log.Printf("[DEBUG] Finished deleting Attestor %q: %#v", d.Id(), res)
	return nil
}

func resourceBinaryAuthorizationAttestorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/attestors/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBinaryAuthorizationAttestorName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenBinaryAuthorizationAttestorDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNote(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["note_reference"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(original["noteReference"], d)
	transformed["public_keys"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(original["publicKeys"], d)
	transformed["delegation_service_account_email"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(original["delegationServiceAccountEmail"], d)
	return []interface{}{transformed}
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"comment":                      flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(original["comment"], d),
			"id":                           flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(original["id"], d),
			"ascii_armored_pgp_public_key": flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(original["asciiArmoredPgpPublicKey"], d),
			"pkix_public_key":              flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(original["pkixPublicKey"], d),
		})
	}
	return transformed
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["public_key_pem"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(original["publicKeyPem"], d)
	transformed["signature_algorithm"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(original["signatureAlgorithm"], d)
	return []interface{}{transformed}
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandBinaryAuthorizationAttestorName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNote(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNoteReference, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(original["note_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNoteReference); val.IsValid() && !isEmptyValue(val) {
		transformed["noteReference"] = transformedNoteReference
	}

	transformedPublicKeys, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(original["public_keys"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKeys); val.IsValid() && !isEmptyValue(val) {
		transformed["publicKeys"] = transformedPublicKeys
	}

	transformedDelegationServiceAccountEmail, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(original["delegation_service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDelegationServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["delegationServiceAccountEmail"] = transformedDelegationServiceAccountEmail
	}

	return transformed, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/notes/(.+)")
	if r.MatchString(v.(string)) {
		return v.(string), nil
	}

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/notes/%s", project, v.(string)), nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedComment, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(original["comment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedComment); val.IsValid() && !isEmptyValue(val) {
			transformed["comment"] = transformedComment
		}

		transformedId, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedAsciiArmoredPgpPublicKey, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(original["ascii_armored_pgp_public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAsciiArmoredPgpPublicKey); val.IsValid() && !isEmptyValue(val) {
			transformed["asciiArmoredPgpPublicKey"] = transformedAsciiArmoredPgpPublicKey
		}

		transformedPkixPublicKey, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(original["pkix_public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPkixPublicKey); val.IsValid() && !isEmptyValue(val) {
			transformed["pkixPublicKey"] = transformedPkixPublicKey
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicKeyPem, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(original["public_key_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKeyPem); val.IsValid() && !isEmptyValue(val) {
		transformed["publicKeyPem"] = transformedPublicKeyPem
	}

	transformedSignatureAlgorithm, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(original["signature_algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignatureAlgorithm); val.IsValid() && !isEmptyValue(val) {
		transformed["signatureAlgorithm"] = transformedSignatureAlgorithm
	}

	return transformed, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceBinaryAuthorizationAttestorEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Field was renamed in GA API
	obj["userOwnedGrafeasNote"] = obj["userOwnedDrydockNote"]
	delete(obj, "userOwnedDrydockNote")

	return obj, nil
}

func resourceBinaryAuthorizationAttestorDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	res["userOwnedDrydockNote"] = res["userOwnedGrafeasNote"]
	delete(res, "userOwnedGrafeasNote")

	return res, nil
}