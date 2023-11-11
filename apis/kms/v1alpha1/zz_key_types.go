/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ExternalKeyReferenceDetailsObservation struct {

	// ExternalKeyId refers to the globally unique key Id associated with the key created in external vault in CTM
	ExternalKeyID *string `json:"externalKeyId,omitempty" tf:"external_key_id,omitempty"`

	// Key version ID associated with the external key.
	ExternalKeyVersionID *string `json:"externalKeyVersionId,omitempty" tf:"external_key_version_id,omitempty"`
}

type ExternalKeyReferenceDetailsParameters struct {
}

type ExternalKeyReferenceObservation struct {
}

type ExternalKeyReferenceParameters struct {

	// ExternalKeyId refers to the globally unique key Id associated with the key created in external vault in CTM
	// +kubebuilder:validation:Required
	ExternalKeyID *string `json:"externalKeyId" tf:"external_key_id,omitempty"`
}

type KeyObservation struct {

	// The OCID of the key version used in cryptographic operations. During key rotation, the service might be in a transitional state where this or a newer key version are used intermittently. The currentKeyVersion property is updated when the service is guaranteed to use the new key version for all subsequent encryption operations.
	CurrentKeyVersion *string `json:"currentKeyVersion,omitempty" tf:"current_key_version,omitempty"`

	// Key reference data to be returned to the customer as a response.
	ExternalKeyReferenceDetails []ExternalKeyReferenceDetailsObservation `json:"externalKeyReferenceDetails,omitempty" tf:"external_key_reference_details,omitempty"`

	// The OCID of the key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A Boolean value that indicates whether the Key belongs to primary Vault or replica vault.
	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	// Key replica details
	ReplicaDetails []ReplicaDetailsObservation `json:"replicaDetails,omitempty" tf:"replica_details,omitempty"`

	// The OCID of the key from which this key was restored.
	RestoredFromKeyID *string `json:"restoredFromKeyId,omitempty" tf:"restored_from_key_id,omitempty"`

	// The key's current lifecycle state.  Example: ENABLED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the key was created, expressed in RFC 3339 timestamp format.  Example: 2018-04-03T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The OCID of the vault that contains this key.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type KeyParameters struct {

	// (Updatable) The OCID of the compartment where you want to create the master encryption key.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Desired state of the key. Possible values : ENABLED or DISABLED
	// +kubebuilder:validation:Optional
	DesiredState *string `json:"desiredState,omitempty" tf:"desired_state,omitempty"`

	// (Updatable) A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// A reference to the key on external key manager.
	// +kubebuilder:validation:Optional
	ExternalKeyReference []ExternalKeyReferenceParameters `json:"externalKeyReference,omitempty" tf:"external_key_reference,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The cryptographic properties of a key.
	// +kubebuilder:validation:Required
	KeyShape []KeyShapeParameters `json:"keyShape" tf:"key_shape,omitempty"`

	// The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
	// +kubebuilder:validation:Required
	ManagementEndpoint *string `json:"managementEndpoint" tf:"management_endpoint,omitempty"`

	// The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of HSM means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of SOFTWARE means that the key persists on the server, protected by the vault's RSA wrapping key which persists on the HSM. All cryptographic operations that use a key with a protection mode of SOFTWARE are performed on the server. By default, a key's protection mode is set to HSM. You can't change a key's protection mode after the key is created or imported. A protection mode of EXTERNAL mean that the key persists on the customer's external key manager which is hosted externally outside of oracle. Oracle only hold a reference to that key. All cryptographic operations that use a key with a protection mode of EXTERNAL are performed by external key manager.
	// +kubebuilder:validation:Optional
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode,omitempty"`

	// (Updatable) Details where key was backed up.
	// +kubebuilder:validation:Optional
	RestoreFromFile []RestoreFromFileParameters `json:"restoreFromFile,omitempty" tf:"restore_from_file,omitempty"`

	// (Updatable) Details where key was backed up
	// +kubebuilder:validation:Optional
	RestoreFromObjectStore []RestoreFromObjectStoreParameters `json:"restoreFromObjectStore,omitempty" tf:"restore_from_object_store,omitempty"`

	// (Updatable) An optional property when flipped triggers restore from restore option provided in config file.
	// +kubebuilder:validation:Optional
	RestoreTrigger *bool `json:"restoreTrigger,omitempty" tf:"restore_trigger,omitempty"`

	// (Updatable) An optional property for the deletion time of the key, expressed in RFC 3339 timestamp format. Example: 2019-04-03T21:10:29.600Z
	// +kubebuilder:validation:Optional
	TimeOfDeletion *string `json:"timeOfDeletion,omitempty" tf:"time_of_deletion,omitempty"`
}

type KeyShapeObservation struct {
}

type KeyShapeParameters struct {

	// The algorithm used by a key's key versions to encrypt or decrypt. Only AES algorithm is supported for External keys.
	// +kubebuilder:validation:Required
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// Supported curve IDs for ECDSA keys.
	// +kubebuilder:validation:Optional
	CurveID *string `json:"curveId,omitempty" tf:"curve_id,omitempty"`

	// The length of the key in bytes, expressed as an integer. Supported values include the following:
	// +kubebuilder:validation:Required
	Length *float64 `json:"length" tf:"length,omitempty"`
}

type ReplicaDetailsObservation struct {

	// ReplicationId associated with a key operation
	ReplicationID *string `json:"replicationId,omitempty" tf:"replication_id,omitempty"`
}

type ReplicaDetailsParameters struct {
}

type RestoreFromFileObservation struct {
}

type RestoreFromFileParameters struct {

	// (Updatable) content length of key's backup binary file
	// +kubebuilder:validation:Required
	ContentLength *string `json:"contentLength" tf:"content_length,omitempty"`

	// (Updatable) content md5 hashed value of key's backup file
	// +kubebuilder:validation:Optional
	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5,omitempty"`

	// Key backup file content.
	// +kubebuilder:validation:Required
	RestoreKeyFromFileDetails *string `json:"restoreKeyFromFileDetails" tf:"restore_key_from_file_details,omitempty"`
}

type RestoreFromObjectStoreObservation struct {
}

type RestoreFromObjectStoreParameters struct {

	// (Updatable) Name of the bucket where key was backed up
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// (Updatable) Type of backup to restore from. Values of "BUCKET", "PRE_AUTHENTICATED_REQUEST_URI" are supported
	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// (Updatable) Namespace of the bucket where key was backed up
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Updatable) Object containing the backup
	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// (Updatable) Pre-authenticated-request-uri of the backup
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

// KeySpec defines the desired state of Key
type KeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyParameters `json:"forProvider"`
}

// KeyStatus defines the observed state of Key.
type KeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Key is the Schema for the Keys API. Provides the Key resource in Oracle Cloud Infrastructure Kms service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeySpec   `json:"spec"`
	Status            KeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyList contains a list of Keys
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

// Repository type metadata.
var (
	Key_Kind             = "Key"
	Key_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Key_Kind}.String()
	Key_KindAPIVersion   = Key_Kind + "." + CRDGroupVersion.String()
	Key_GroupVersionKind = CRDGroupVersion.WithKind(Key_Kind)
)

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
