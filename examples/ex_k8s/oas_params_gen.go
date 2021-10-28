// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
)

type CreateAdmissionregistrationV1MutatingWebhookConfigurationParams struct {
	DryRun       string
	FieldManager string
}

type CreateAdmissionregistrationV1ValidatingWebhookConfigurationParams struct {
	DryRun       string
	FieldManager string
}

type CreateApiextensionsV1CustomResourceDefinitionParams struct {
	DryRun       string
	FieldManager string
}

type CreateApiregistrationV1APIServiceParams struct {
	DryRun       string
	FieldManager string
}

type CreateCertificatesV1CertificateSigningRequestParams struct {
	DryRun       string
	FieldManager string
}

type CreateCoreV1NamespaceParams struct {
	DryRun       string
	FieldManager string
}

type CreateCoreV1NodeParams struct {
	DryRun       string
	FieldManager string
}

type CreateCoreV1PersistentVolumeParams struct {
	DryRun       string
	FieldManager string
}

type CreateFlowcontrolApiserverV1beta1FlowSchemaParams struct {
	DryRun       string
	FieldManager string
}

type CreateFlowcontrolApiserverV1beta1PriorityLevelConfigurationParams struct {
	DryRun       string
	FieldManager string
}

type CreateFlowcontrolApiserverV1beta2FlowSchemaParams struct {
	DryRun       string
	FieldManager string
}

type CreateFlowcontrolApiserverV1beta2PriorityLevelConfigurationParams struct {
	DryRun       string
	FieldManager string
}

type CreateInternalApiserverV1alpha1StorageVersionParams struct {
	DryRun       string
	FieldManager string
}

type CreateNetworkingV1IngressClassParams struct {
	DryRun       string
	FieldManager string
}

type CreateNodeV1RuntimeClassParams struct {
	DryRun       string
	FieldManager string
}

type CreateNodeV1alpha1RuntimeClassParams struct {
	DryRun       string
	FieldManager string
}

type CreateNodeV1beta1RuntimeClassParams struct {
	DryRun       string
	FieldManager string
}

type CreatePolicyV1beta1PodSecurityPolicyParams struct {
	DryRun       string
	FieldManager string
}

type CreateRbacAuthorizationV1ClusterRoleParams struct {
	DryRun       string
	FieldManager string
}

type CreateRbacAuthorizationV1ClusterRoleBindingParams struct {
	DryRun       string
	FieldManager string
}

type CreateSchedulingV1PriorityClassParams struct {
	DryRun       string
	FieldManager string
}

type CreateStorageV1CSIDriverParams struct {
	DryRun       string
	FieldManager string
}

type CreateStorageV1CSINodeParams struct {
	DryRun       string
	FieldManager string
}

type CreateStorageV1StorageClassParams struct {
	DryRun       string
	FieldManager string
}

type CreateStorageV1VolumeAttachmentParams struct {
	DryRun       string
	FieldManager string
}

type DeleteAdmissionregistrationV1CollectionMutatingWebhookConfigurationParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteApiextensionsV1CollectionCustomResourceDefinitionParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteApiregistrationV1CollectionAPIServiceParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteCertificatesV1CollectionCertificateSigningRequestParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteCoreV1CollectionNodeParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteCoreV1CollectionPersistentVolumeParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteFlowcontrolApiserverV1beta1CollectionFlowSchemaParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteFlowcontrolApiserverV1beta1CollectionPriorityLevelConfigurationParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteFlowcontrolApiserverV1beta2CollectionFlowSchemaParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteFlowcontrolApiserverV1beta2CollectionPriorityLevelConfigurationParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteInternalApiserverV1alpha1CollectionStorageVersionParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteNetworkingV1CollectionIngressClassParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteNodeV1CollectionRuntimeClassParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteNodeV1alpha1CollectionRuntimeClassParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteNodeV1beta1CollectionRuntimeClassParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeletePolicyV1beta1CollectionPodSecurityPolicyParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteRbacAuthorizationV1CollectionClusterRoleParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteRbacAuthorizationV1CollectionClusterRoleBindingParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteSchedulingV1CollectionPriorityClassParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteStorageV1CollectionCSIDriverParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteStorageV1CollectionCSINodeParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteStorageV1CollectionStorageClassParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type DeleteStorageV1CollectionVolumeAttachmentParams struct {
	Continue             string
	DryRun               string
	FieldSelector        string
	GracePeriodSeconds   int
	LabelSelector        string
	Limit                int
	OrphanDependents     bool
	PropagationPolicy    string
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
}

type ListAdmissionregistrationV1MutatingWebhookConfigurationParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListAdmissionregistrationV1ValidatingWebhookConfigurationParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListApiextensionsV1CustomResourceDefinitionParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListApiregistrationV1APIServiceParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListCertificatesV1CertificateSigningRequestParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListCoreV1NamespaceParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListCoreV1NodeParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListCoreV1PersistentVolumeParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListFlowcontrolApiserverV1beta1FlowSchemaParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListFlowcontrolApiserverV1beta1PriorityLevelConfigurationParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListFlowcontrolApiserverV1beta2FlowSchemaParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListFlowcontrolApiserverV1beta2PriorityLevelConfigurationParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListInternalApiserverV1alpha1StorageVersionParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListNetworkingV1IngressClassParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListNodeV1RuntimeClassParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListNodeV1alpha1RuntimeClassParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListNodeV1beta1RuntimeClassParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListPolicyV1beta1PodSecurityPolicyParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListRbacAuthorizationV1ClusterRoleParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListRbacAuthorizationV1ClusterRoleBindingParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListSchedulingV1PriorityClassParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListStorageV1CSIDriverParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListStorageV1CSINodeParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListStorageV1StorageClassParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}

type ListStorageV1VolumeAttachmentParams struct {
	AllowWatchBookmarks  bool
	Continue             string
	FieldSelector        string
	LabelSelector        string
	Limit                int
	ResourceVersion      string
	ResourceVersionMatch string
	TimeoutSeconds       int
	Watch                bool
}
