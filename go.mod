go 1.18

module github.com/cosmos/cosmos-sdk

require (
	cosmossdk.io/api v0.4.2
	cosmossdk.io/collections v0.2.0
	cosmossdk.io/core v0.8.0
	cosmossdk.io/depinject v1.0.0-alpha.3
	cosmossdk.io/errors v1.0.0-beta.7
	cosmossdk.io/log v1.1.0
	cosmossdk.io/math v1.0.1
	cosmossdk.io/store v0.1.0-alpha.1
	cosmossdk.io/x/tx v0.8.0
	github.com/99designs/keyring v1.1.6
	github.com/armon/go-metrics v0.4.1
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd/btcec/v2 v2.3.2
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/coinbase/rosetta-sdk-go v0.7.0
	github.com/cometbft/cometbft v0.37.1
	github.com/confio/ics23/go v0.9.0
	github.com/cosmos/btcutil v1.0.4
	github.com/cosmos/cosmos-db v1.0.0
	github.com/cosmos/cosmos-proto v1.0.0-beta.3
	github.com/cosmos/cosmos-sdk/db v1.0.0-beta.1
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/gogoproto v1.4.10
	github.com/cosmos/iavl v0.21.0-beta.1
	github.com/cosmos/ledger-cosmos-go v0.11.1
	github.com/deckarep/golang-set v1.8.0
	github.com/gogo/gateway v1.1.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/golang-lru/v2 v2.0.1
	github.com/hdevalence/ed25519consensus v0.0.0-20210204194344-59a8610d2b87
	github.com/improbable-eng/grpc-web v0.14.1
	github.com/jhump/protoreflect v1.15.1
	github.com/magiconair/properties v1.8.6
	github.com/manifoldco/promptui v0.9.0
	github.com/mattn/go-isatty v0.0.18
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.15.1
	github.com/prometheus/common v0.44.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/zerolog v1.29.1
	github.com/savaki/jq v0.0.0-20161209013833-0e6baecebbf8
	github.com/sei-protocol/sei-tm-db v0.0.5
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.13.0
	github.com/stretchr/testify v1.8.4
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.37.0-dev
	github.com/tendermint/tm-db v0.6.8-0.20220519162814-e24b96538a12
	github.com/yourbasic/graph v0.0.0-20210606180040-8ecfec1c2869
	go.opentelemetry.io/otel v1.11.0
	go.opentelemetry.io/otel/exporters/jaeger v1.9.0
	go.opentelemetry.io/otel/sdk v1.9.0
	go.opentelemetry.io/otel/trace v1.11.0
	golang.org/x/crypto v0.9.0
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible
	gotest.tools/v3 v3.4.0
	pgregory.net/rapid v0.6.2
	sigs.k8s.io/yaml v1.3.0
)

require (
	filippo.io/edwards25519 v1.0.0-beta.2 // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/bufbuild/protocompile v0.4.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cockroachdb/errors v1.9.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v0.0.0-20230525220056-bb4fc9527b3b // indirect
	github.com/cockroachdb/redact v1.1.4 // indirect
	github.com/cometbft/cometbft-db v0.7.0 // indirect
	github.com/cosmos/gorocksdb v1.2.0 // indirect
	github.com/cosmos/ledger-go v0.9.2 // indirect
	github.com/creachadair/taskgroup v0.3.2 // indirect
	github.com/danieljoos/wincred v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/badger/v3 v3.2103.2 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/dvsekhvalnov/jose2go v0.0.0-20200901110807-248326c1351b // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/getsentry/sentry-go v0.21.0 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/flatbuffers v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/orderedcode v0.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/keybase/go-keychain v0.0.0-20190712205309-48d3d31d256d // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/linxGnu/grocksdb v1.7.16 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20210601165009-122bf33a46e0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20210609091139-0a56a4bca00b // indirect
	github.com/onsi/gomega v1.27.6 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/petermattis/goid v0.0.0-20221215004737-a150e88a970d // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/procfs v0.10.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c // indirect
	github.com/tidwall/btree v1.6.0 // indirect
	github.com/zondax/hid v0.9.0 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/term v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.9.3 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/btcsuite/btcd/btcec => github.com/btcsuite/btcd v0.22.1
	github.com/confio/ics23/go => github.com/cosmos/cosmos-sdk/ics23/go v0.8.0
	github.com/cosmos/iavl => github.com/sei-protocol/sei-iavl v0.1.7

	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/cosmos/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.7.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

	github.com/tendermint/tendermint => github.com/Timwood0x10/sei-tendermint v0.0.0-20230703080427-7a134e7b9339

	// latest grpc doesn't work with with our modified proto compiler, so we need to enforce
	// the following version across all dependencies.
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
