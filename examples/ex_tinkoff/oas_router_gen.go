// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
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
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func skipSlash(p string) string {
	if len(p) > 0 && p[0] == '/' {
		return p[1:]
	}
	return p
}

// nextElem return next path element from p and forwarded p.
func nextElem(p string) (elem, next string) {
	p = skipSlash(p)
	idx := strings.IndexByte(p, '/')
	if idx < 0 {
		idx = len(p)
	}
	return p[:idx], p[idx:]
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if len(p) == 0 {
		s.notFound(w, r)
		return
	}

	var (
		elem string            // current element, without slashes
		args map[string]string // lazily initialized
	)

	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		// Root edge.
		elem, p = nextElem(p)
		switch elem {
		case "market": // -> 1
			// Edge: 1, path: "market".
			elem, p = nextElem(p)
			switch elem {
			case "bonds": // -> 2
				// GET /market/bonds
				s.handleMarketBondsGetRequest(args, w, r)
				return
			case "candles": // -> 3
				// GET /market/candles
				s.handleMarketCandlesGetRequest(args, w, r)
				return
			case "currencies": // -> 4
				// GET /market/currencies
				s.handleMarketCurrenciesGetRequest(args, w, r)
				return
			case "etfs": // -> 5
				// GET /market/etfs
				s.handleMarketEtfsGetRequest(args, w, r)
				return
			case "orderbook": // -> 6
				// GET /market/orderbook
				s.handleMarketOrderbookGetRequest(args, w, r)
				return
			case "search": // -> 7
				// Edge: 7, path: "search".
				elem, p = nextElem(p)
				switch elem {
				case "by-figi": // -> 8
					// GET /market/search/by-figi
					s.handleMarketSearchByFigiGetRequest(args, w, r)
					return
				case "by-ticker": // -> 9
					// GET /market/search/by-ticker
					s.handleMarketSearchByTickerGetRequest(args, w, r)
					return
				default:
					s.notFound(w, r)
					return
				}
			case "stocks": // -> 10
				// GET /market/stocks
				s.handleMarketStocksGetRequest(args, w, r)
				return
			default:
				s.notFound(w, r)
				return
			}
		case "operations": // -> 11
			// GET /operations
			s.handleOperationsGetRequest(args, w, r)
			return
		case "orders": // -> 12
			// GET /orders
			s.handleOrdersGetRequest(args, w, r)
			return
		case "portfolio": // -> 13
			// Edge: 13, path: "portfolio".
			elem, p = nextElem(p)
			if len(elem) == 0 {
				// GET /portfolio.
				s.handlePortfolioGetRequest(args, w, r)
				return
			}
			switch elem {
			case "currencies": // -> 14
				// GET /portfolio/currencies
				s.handlePortfolioCurrenciesGetRequest(args, w, r)
				return
			default:
				// GET /portfolio.
				s.handlePortfolioGetRequest(args, w, r)
				return
			}
		case "user": // -> 15
			// Edge: 15, path: "user".
			elem, p = nextElem(p)
			switch elem {
			case "accounts": // -> 16
				// GET /user/accounts
				s.handleUserAccountsGetRequest(args, w, r)
				return
			default:
				s.notFound(w, r)
				return
			}
		default:
			s.notFound(w, r)
			return
		}
	case "POST":
		// Root edge.
		elem, p = nextElem(p)
		switch elem {
		case "orders": // -> 1
			// Edge: 1, path: "orders".
			elem, p = nextElem(p)
			switch elem {
			case "cancel": // -> 2
				// POST /orders/cancel
				s.handleOrdersCancelPostRequest(args, w, r)
				return
			case "limit-order": // -> 3
				// POST /orders/limit-order
				s.handleOrdersLimitOrderPostRequest(args, w, r)
				return
			case "market-order": // -> 4
				// POST /orders/market-order
				s.handleOrdersMarketOrderPostRequest(args, w, r)
				return
			default:
				s.notFound(w, r)
				return
			}
		case "sandbox": // -> 5
			// Edge: 5, path: "sandbox".
			elem, p = nextElem(p)
			switch elem {
			case "clear": // -> 6
				// POST /sandbox/clear
				s.handleSandboxClearPostRequest(args, w, r)
				return
			case "currencies": // -> 7
				// Edge: 7, path: "currencies".
				elem, p = nextElem(p)
				switch elem {
				case "balance": // -> 8
					// POST /sandbox/currencies/balance
					s.handleSandboxCurrenciesBalancePostRequest(args, w, r)
					return
				default:
					s.notFound(w, r)
					return
				}
			case "positions": // -> 9
				// Edge: 9, path: "positions".
				elem, p = nextElem(p)
				switch elem {
				case "balance": // -> 10
					// POST /sandbox/positions/balance
					s.handleSandboxPositionsBalancePostRequest(args, w, r)
					return
				default:
					s.notFound(w, r)
					return
				}
			case "register": // -> 11
				// POST /sandbox/register
				s.handleSandboxRegisterPostRequest(args, w, r)
				return
			case "remove": // -> 12
				// POST /sandbox/remove
				s.handleSandboxRemovePostRequest(args, w, r)
				return
			default:
				s.notFound(w, r)
				return
			}
		default:
			s.notFound(w, r)
			return
		}
	default:
		s.notFound(w, r)
		return
	}
}
