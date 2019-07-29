package lib

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"io"
	"log"

	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/api"
)

type (
	Func struct {
		rpc   *rpc.RPC
		specs Specs
	}

	// Spec - the user inputable.
	Spec struct {
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Handler     api.Handler  `json:"-"`        // the handler func
		Version     string       `json:"version"`  // the version of this func
		Settings    *Settings    `json:"settings"` // Various settings
		RPCBinding  *RPCBinding  `json:"rpc"`      // the rpc binding
		HTTPBinding *HTTPBinding `json:"http"`     // the http binding, if any
	}

	// Settings - some settings
	Settings struct {
		RPC         bool `json:"rpc"`         // is this func rpc or eventy?
		Timeout     int  `json:"timeout"`     // timeout in seconds in case of rpc=true
		Loadbalance bool `json:"loadbalance"` // do we bind to a queue or not
	}

	// FuncSpec - what the servers see/want/need.
	FuncSpec struct {
		Spec      *Spec  `json:"spec"`
		Namespace string `json:"ns,omitempty"` // our namespace if set
	}

	// HTTPBinding - how to bind this to web requests.
	HTTPBinding struct {
		Verb  string   `json:"verb"`            // prefered http verb
		URL   string   `json:"url"`             // /func/:param
		Query []string `json:"query,omitempty"` // query params to forward if present
	}

	// RPCBinding - how to bind this to rpc requests.
	RPCBinding struct {
		Topic string `json:"topic,omitempty"`
	}

	Specs []*FuncSpec
)

// TODO replace ns with a jwt that contains namespace and optionally a topic
// That would require some sort of key to be distributed too though...

var (
	ns    string
	topic string
)

func NewSpec() *Spec {
	return &Spec{}
}

func NewFunc() *Func {
	flag.StringVar(&ns, "ns", "default", "set the namespace of this func")
	flag.Parse()

	return &Func{
		specs: make(Specs, 0),
	}
}

func (f *Func) Register(spec *Spec) {
	if spec.IsValid() {
		funcSpec := &FuncSpec{
			Spec:      spec,
			Namespace: ns,
		}
		f.specs = append(f.specs, funcSpec)
	} else {
		log.Println("Spec is not valid. It must have name, description, version & handler.")
	}
}

func (f *Func) Start() {
	conn, err := nuts.Connect()

	if err != nil {
		log.Panic(err)
	}

	f.specs.ForAll(func(spec *FuncSpec) error {
		spec.Namespace = ns
		return nil
	})

	f.rpc = rpc.NewRpc(conn)

	err = f.register()

	if err != nil {
		log.Println(err.Error())
	}

	for _, spec := range f.specs {
		f.rpc.Handler(spec.Spec.Topic(spec.Namespace), spec.Namespace, spec.Spec.Handler)
		log.Printf("[%s] (%s) started in namespace [%s].\n", spec.Spec.Name, spec.Spec.Version, ns)
	}

	f.rpc.Run()
}

func (f *Func) register() error {
	return f.specs.ForAll(func(spec *FuncSpec) error {
		msg, err := api.NewMessage(spec)

		if err != nil {
			return err
		}

		return f.rpc.Trigger("func.discovery", msg)
	})
}

func (s Specs) ForAll(action func(*FuncSpec) error) error {
	for _, spec := range s {
		err := action(spec)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s Specs) ID() string {
	hasher := sha1.New()

	s.ForAll(func(spec *FuncSpec) error {
		_, err := io.WriteString(hasher, spec.Spec.Name)
		return err
	})

	bs := hasher.Sum(nil)

	return hex.EncodeToString(bs)
}

func (s *Spec) IsValid() bool {
	return s.Name != "" &&
		s.Description != "" &&
		s.Handler != nil &&
		s.Version != ""
}

func (s *Spec) Topic(ns string) string {
	hasher := sha1.New()

	io.WriteString(hasher, ns)
	io.WriteString(hasher, s.Name)
	io.WriteString(hasher, s.Version)

	return hex.EncodeToString(hasher.Sum(nil))
}
