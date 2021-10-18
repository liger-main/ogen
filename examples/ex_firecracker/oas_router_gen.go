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
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
)

func Register(r chi.Router, s Server) {
	r.MethodFunc("GET", "/", NewDescribeInstanceHandler(s))
	r.MethodFunc("PUT", "/actions", NewCreateSyncActionHandler(s))
	r.MethodFunc("GET", "/balloon", NewDescribeBalloonConfigHandler(s))
	r.MethodFunc("PUT", "/balloon", NewPutBalloonHandler(s))
	r.MethodFunc("PATCH", "/balloon", NewPatchBalloonHandler(s))
	r.MethodFunc("GET", "/balloon/statistics", NewDescribeBalloonStatsHandler(s))
	r.MethodFunc("PATCH", "/balloon/statistics", NewPatchBalloonStatsIntervalHandler(s))
	r.MethodFunc("PUT", "/boot-source", NewPutGuestBootSourceHandler(s))
	r.MethodFunc("PUT", "/drives/{drive_id}", NewPutGuestDriveByIDHandler(s))
	r.MethodFunc("PATCH", "/drives/{drive_id}", NewPatchGuestDriveByIDHandler(s))
	r.MethodFunc("PUT", "/logger", NewPutLoggerHandler(s))
	r.MethodFunc("GET", "/machine-config", NewGetMachineConfigurationHandler(s))
	r.MethodFunc("PUT", "/machine-config", NewPutMachineConfigurationHandler(s))
	r.MethodFunc("PATCH", "/machine-config", NewPatchMachineConfigurationHandler(s))
	r.MethodFunc("PUT", "/metrics", NewPutMetricsHandler(s))
	r.MethodFunc("GET", "/mmds", NewMmdsGetHandler(s))
	r.MethodFunc("PUT", "/mmds", NewMmdsPutHandler(s))
	r.MethodFunc("PATCH", "/mmds", NewMmdsPatchHandler(s))
	r.MethodFunc("PUT", "/mmds/config", NewMmdsConfigPutHandler(s))
	r.MethodFunc("PUT", "/network-interfaces/{iface_id}", NewPutGuestNetworkInterfaceByIDHandler(s))
	r.MethodFunc("PATCH", "/network-interfaces/{iface_id}", NewPatchGuestNetworkInterfaceByIDHandler(s))
	r.MethodFunc("PUT", "/snapshot/create", NewCreateSnapshotHandler(s))
	r.MethodFunc("PUT", "/snapshot/load", NewLoadSnapshotHandler(s))
	r.MethodFunc("PATCH", "/vm", NewPatchVmHandler(s))
	r.MethodFunc("GET", "/vm/config", NewGetExportVmConfigHandler(s))
	r.MethodFunc("PUT", "/vsock", NewPutGuestVsockHandler(s))
}
