package gwx

import (
	"errors"
	"fmt"
	"github.com/fullstorydev/grpcurl"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/jhump/protoreflect/desc"
	"github.com/zeromicro/go-zero/core/search"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/rest/router"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"net/http"
	"path"
	"strings"
)

func CreateDescriptorSource(up gateway.Upstream) (grpcurl.DescriptorSource, error) {
	var source grpcurl.DescriptorSource
	var err error

	if len(up.ProtoSets) > 0 {
		source, err = grpcurl.DescriptorSourceFromProtoSets(up.ProtoSets...)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no implement")
	}

	return source, nil
}

func Parse(upstreams []gateway.Upstream) ([]Method, error) {
	var res []Method
	for _, up := range upstreams {
		source, err := CreateDescriptorSource(up)
		if err != nil {
			return nil, err
		}

		methods, err := GetMethods(source)
		if err != nil {
			return nil, err
		} else {
			res = append(res, methods...)
		}
	}
	return res, nil
}

type Method struct {
	HttpMethod   string
	HttpPath     string
	RpcPath      string
	AuthRequired bool
}

func GetMethods(source grpcurl.DescriptorSource) ([]Method, error) {
	svcs, err := source.ListServices()
	if err != nil {
		return nil, err
	}

	var methods []Method
	for _, svc := range svcs {
		d, err := source.FindSymbol(svc)
		if err != nil {
			return nil, err
		}

		switch val := d.(type) {
		case *desc.ServiceDescriptor:
			svcMethods := val.GetMethods()
			for _, method := range svcMethods {
				rpcPath := fmt.Sprintf("%s/%s", svc, method.GetName())
				ext := proto.GetExtension(method.GetMethodOptions(), annotations.E_Http)
				authRequire := true

				if openApiV2Ext := proto.GetExtension(method.GetMethodOptions(), options.E_Openapiv2Operation); openApiV2Ext != nil {
					if op, ok := openApiV2Ext.(*options.Operation); ok {
						for _, security := range op.GetSecurity() {
							for k, v := range security.SecurityRequirement {
								if k == "Anonymous" {
									authRequire = false
								}
								_ = v
							}
						}
					}
				}

				switch rule := ext.(type) {
				case *annotations.HttpRule:
					if rule == nil {
						methods = append(methods, Method{
							RpcPath:      rpcPath,
							AuthRequired: authRequire,
						})
						continue
					}

					switch httpRule := rule.GetPattern().(type) {
					case *annotations.HttpRule_Get:
						methods = append(methods, Method{
							HttpMethod:   http.MethodGet,
							HttpPath:     adjustHttpPath(httpRule.Get),
							RpcPath:      rpcPath,
							AuthRequired: authRequire,
						})
					case *annotations.HttpRule_Post:
						methods = append(methods, Method{
							HttpMethod:   http.MethodPost,
							HttpPath:     adjustHttpPath(httpRule.Post),
							RpcPath:      rpcPath,
							AuthRequired: authRequire,
						})
					case *annotations.HttpRule_Put:
						methods = append(methods, Method{
							HttpMethod:   http.MethodPut,
							HttpPath:     adjustHttpPath(httpRule.Put),
							RpcPath:      rpcPath,
							AuthRequired: authRequire,
						})
					case *annotations.HttpRule_Delete:
						methods = append(methods, Method{
							HttpMethod:   http.MethodDelete,
							HttpPath:     adjustHttpPath(httpRule.Delete),
							RpcPath:      rpcPath,
							AuthRequired: authRequire,
						})
					case *annotations.HttpRule_Patch:
						methods = append(methods, Method{
							HttpMethod:   http.MethodPatch,
							HttpPath:     adjustHttpPath(httpRule.Patch),
							RpcPath:      rpcPath,
							AuthRequired: authRequire,
						})
					default:
						methods = append(methods, Method{
							RpcPath:      rpcPath,
							AuthRequired: authRequire,
						})
					}
				default:
					methods = append(methods, Method{
						RpcPath:      rpcPath,
						AuthRequired: authRequire,
					})
				}
			}
		}
	}

	return methods, nil
}

func adjustHttpPath(path string) string {
	path = strings.ReplaceAll(path, "{", ":")
	path = strings.ReplaceAll(path, "}", "")
	return path
}

type Router struct {
	trees  map[string]*search.Tree
	mapper map[string]bool
}

func MustNewRouter(upstreams []gateway.Upstream) *Router {
	r := &Router{
		trees:  make(map[string]*search.Tree),
		mapper: make(map[string]bool),
	}
	methods, err := Parse(upstreams)
	if err != nil {
		panic(err)
	}
	for _, method := range methods {
		if method.HttpPath == "" {
			continue
		}
		if err := r.Add(method.HttpMethod, method.HttpPath, method.RpcPath); err != nil {
			panic(err)
		}
		r.mapper[method.RpcPath] = method.AuthRequired
	}
	return r
}

func (r *Router) Add(method, httpPath, rpcPath string) error {
	if len(httpPath) == 0 || httpPath[0] != '/' {
		return router.ErrInvalidPath
	}

	httpPath = path.Clean(httpPath)
	tree, ok := r.trees[method]
	if ok {
		return tree.Add(httpPath, rpcPath)
	}

	tree = search.NewTree()
	r.trees[method] = tree
	return tree.Add(httpPath, rpcPath)
}

// Search
// method: get, post, put, delete, etc...
// httpPath: /xxx/yyy?zzz=111
func (r *Router) Search(method, httpPath string) (search.Result, bool) {
	httpPath = path.Clean(httpPath)
	tree, ok := r.trees[method]
	if ok {
		return tree.Search(httpPath)
	} else {
		return search.Result{}, false
	}
}

// IsRequireAuth return true if not found
func (r *Router) IsRequireAuth(method, httpPath string) bool {
	res, ok := r.Search(method, httpPath)
	if ok {
		rpcPath, ok := res.Item.(string)
		if !ok {
			return true
		}

		required, ok := r.mapper[rpcPath]
		if ok {
			return required
		} else {
			return true
		}
	} else {
		return true
	}
}
