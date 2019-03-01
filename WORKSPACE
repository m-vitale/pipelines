load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "492c3ac68ed9dcf527a07e6a1b2dcbf199c6bf8b35517951467ac32e421c06c1",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.17.0/rules_go-0.17.0.tar.gz"],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "7949fc6cc17b5b191103e97481cf8889217263acf52e00b560683413af204fcb",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.16.0/bazel-gazelle-0.16.0.tar.gz"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_archive(
    name = "org_tensorflow",
    sha256 = "24570d860d87dcfb936f53fb8dd30302452d0aa6b8b8537e4555c1bf839121a6",
    strip_prefix = "tensorflow-1.13.0-rc0",
    urls = [
        "https://github.com/tensorflow/tensorflow/archive/v1.13.0-rc0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_closure",
    sha256 = "43c9b882fa921923bcba764453f4058d102bece35a37c9f6383c713004aacff1",
    strip_prefix = "rules_closure-9889e2348259a5aad7e805547c1a0cf311cfcd91",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_closure/archive/9889e2348259a5aad7e805547c1a0cf311cfcd91.tar.gz",
        "https://github.com/bazelbuild/rules_closure/archive/9889e2348259a5aad7e805547c1a0cf311cfcd91.tar.gz",  # 2018-12-21
    ],
)

load("@org_tensorflow//tensorflow:workspace.bzl", "tf_workspace")

tf_workspace()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")

go_repository(
    name = "google_ml_metadata",
    commit = "e95231f88a8f7ed5224007ba554e080a73ac30a1",
    importpath = "github.com/google/ml-metadata",
)

new_git_repository(
    name = "libmysqlclient",
    build_file = "@google_ml_metadata//ml_metadata:libmysqlclient.BUILD",
    remote = "https://github.com/MariaDB/mariadb-connector-c.git",
    tag = "v3.0.8-release",
    workspace_file = "@google_ml_metadata//ml_metadata:libmysqlclient.WORKSPACE",
)

go_repository(
    name = "io_k8s_client_go",
    build_file_proto_mode = "disable_global",
    commit = "59698c7d9724",
    importpath = "k8s.io/client-go",
)

go_repository(
    name = "io_k8s_apimachinery",
    build_file_proto_mode = "disable_global",
    commit = "103fd098999d",
    importpath = "k8s.io/apimachinery",
)

go_repository(
    name = "com_github_google_gofuzz",
    commit = "24818f796faf",
    importpath = "github.com/google/gofuzz",
)

go_repository(
    name = "io_k8s_sigs_controller_runtime",
    build_extra_args = ["-exclude=vendor"],
    commit = "5558165425ef",
    importpath = "sigs.k8s.io/controller-runtime",
)

go_repository(
    name = "io_k8s_api",
    build_file_proto_mode = "disable_global",
    commit = "2d6f90ab1293",
    importpath = "k8s.io/api",
)

go_repository(
    name = "io_k8s_kubernetes",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "k8s.io/kubernetes",
    tag = "v1.11.1",
)

go_repository(
    name = "com_github_googleapis_gnostic",
    build_file_proto_mode = "disable",
    importpath = "github.com/googleapis/gnostic",
    tag = "v0.2.0",
)

# for @io_k8s_kubernetes
http_archive(
    name = "io_kubernetes_build",
    sha256 = "1188feb932cefad328b0a3dd75b3ebd1d79dd26dbdd723f019ceb760e27ba6d8",
    strip_prefix = "repo-infra-84d52408a061e87d45aebf5a0867246bdf66d180",
    urls = ["https://github.com/kubernetes/repo-infra/archive/84d52408a061e87d45aebf5a0867246bdf66d180.tar.gz"],
)

BAZEL_BUILDTOOLS_VERSION = "49a6c199e3fbf5d94534b2771868677d3f9c6de9"

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "edf39af5fc257521e4af4c40829fffe8fba6d0ebff9f4dd69a6f8f1223ae047b",
    strip_prefix = "buildtools-%s" % BAZEL_BUILDTOOLS_VERSION,
    url = "https://github.com/bazelbuild/buildtools/archive/%s.zip" % BAZEL_BUILDTOOLS_VERSION,
)

http_archive(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    strip_prefix = "grpc-gateway-1.6.3",
    url = "https://github.com/grpc-ecosystem/grpc-gateway/archive/v1.6.3.tar.gz",
)

go_repository(
    name = "com_github_go_swagger",
    importpath = "github.com/go-swagger/go-swagger",
    tag = "v0.18.0",
)

http_archive(
    name = "com_github_mbrukman_autogen",
    strip_prefix = "autogen-0.3",
    url = "https://github.com/mbrukman/autogen/archive/v0.3.tar.gz",
)

# The following were generated by Gazelle. If go.mod is updated, delete the
# following lines and run:
# bazel run //:gazelle -- update-repos --from_file=go.mod

go_repository(
    name = "co_honnef_go_tools",
    commit = "88497007e858",
    importpath = "honnef.co/go/tools",
)

go_repository(
    name = "com_github_argoproj_argo",
    importpath = "github.com/argoproj/argo",
    tag = "v2.2.0",
)

go_repository(
    name = "com_github_armon_consul_api",
    commit = "eb2c6b5be1b6",
    importpath = "github.com/armon/consul-api",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    commit = "f9ffefc3facf",
    importpath = "github.com/asaskevich/govalidator",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    tag = "v0.3.1",
)

go_repository(
    name = "com_github_cenkalti_backoff",
    importpath = "github.com/cenkalti/backoff",
    tag = "v2.0.0",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    tag = "v0.3.4",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    tag = "v3.3.10",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    tag = "v2.0.0",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_denisenkom_go_mssqldb",
    commit = "4e0d7dc8888f",
    importpath = "github.com/denisenkom/go-mssqldb",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    tag = "v0.3.3",
)

go_repository(
    name = "com_github_docker_spdystream",
    commit = "bc6354cbbc29",
    importpath = "github.com/docker/spdystream",
)

go_repository(
    name = "com_github_elazarl_goproxy",
    commit = "2ce16c963a8a",
    importpath = "github.com/elazarl/goproxy",
)

go_repository(
    name = "com_github_emicklei_go_restful",
    importpath = "github.com/emicklei/go-restful",
    tag = "v2.8.0",
)

go_repository(
    name = "com_github_erikstmartin_go_testdb",
    commit = "8d10e4a1bae5",
    importpath = "github.com/erikstmartin/go-testdb",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.7",
)

go_repository(
    name = "com_github_ghodss_yaml",
    commit = "c7ce16629ff4",
    importpath = "github.com/ghodss/yaml",
)

go_repository(
    name = "com_github_globalsign_mgo",
    commit = "eeefdecb41b8",
    importpath = "github.com/globalsign/mgo",
)

go_repository(
    name = "com_github_go_ini_ini",
    importpath = "github.com/go-ini/ini",
    tag = "v1.38.1",
)

go_repository(
    name = "com_github_go_openapi_analysis",
    importpath = "github.com/go-openapi/analysis",
    tag = "v0.17.2",
)

go_repository(
    name = "com_github_go_openapi_errors",
    importpath = "github.com/go-openapi/errors",
    tag = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    importpath = "github.com/go-openapi/jsonpointer",
    tag = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    importpath = "github.com/go-openapi/jsonreference",
    tag = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_loads",
    importpath = "github.com/go-openapi/loads",
    tag = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_runtime",
    commit = "aadb2cc7b886",
    importpath = "github.com/go-openapi/runtime",
)

go_repository(
    name = "com_github_go_openapi_spec",
    importpath = "github.com/go-openapi/spec",
    tag = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_strfmt",
    importpath = "github.com/go-openapi/strfmt",
    tag = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_swag",
    importpath = "github.com/go-openapi/swag",
    tag = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_validate",
    importpath = "github.com/go-openapi/validate",
    tag = "v0.17.2",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    tag = "v1.4.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b",
    importpath = "github.com/golang/glog",
)

go_repository(
    name = "com_github_golang_groupcache",
    commit = "24b0969c4cb7",
    importpath = "github.com/golang/groupcache",
)

go_repository(
    name = "com_github_golang_lint",
    commit = "06c8688daad7",
    importpath = "github.com/golang/lint",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_google_btree",
    commit = "e89373fe6b4a",
    importpath = "github.com/google/btree",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    commit = "d547d1d9531e",
    importpath = "github.com/gopherjs/gopherjs",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    commit = "9cad4c3443a7",
    importpath = "github.com/gregjones/httpcache",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    commit = "0fb14efe8c47",
    importpath = "github.com/hashicorp/golang-lru",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    tag = "v0.3.5",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_iris_contrib_go_uuid",
    importpath = "github.com/iris-contrib/go.uuid",
    tag = "v2.0.0",
)

go_repository(
    name = "com_github_jinzhu_gorm",
    importpath = "github.com/jinzhu/gorm",
    tag = "v1.9.1",
)

go_repository(
    name = "com_github_jinzhu_inflection",
    commit = "04140366298a",
    importpath = "github.com/jinzhu/inflection",
)

go_repository(
    name = "com_github_jinzhu_now",
    commit = "8ec929ed50c3",
    importpath = "github.com/jinzhu/now",
)

go_repository(
    name = "com_github_json_iterator_go",
    commit = "ab8a2e0c74be",
    importpath = "github.com/json-iterator/go",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    tag = "v4.2.1",
)

go_repository(
    name = "com_github_kataras_iris",
    importpath = "github.com/kataras/iris",
    tag = "v10.6.7",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_lann_builder",
    commit = "47ae307949d0",
    importpath = "github.com/lann/builder",
)

go_repository(
    name = "com_github_lann_ps",
    commit = "62de8c46ede0",
    importpath = "github.com/lann/ps",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    tag = "v1.8.0",
)

go_repository(
    name = "com_github_mailru_easyjson",
    commit = "60711f1a8329",
    importpath = "github.com/mailru/easyjson",
)

go_repository(
    name = "com_github_masterminds_squirrel",
    commit = "fa735ea14f09",
    importpath = "github.com/Masterminds/squirrel",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    tag = "v1.9.0",
)

go_repository(
    name = "com_github_minio_minio_go",
    importpath = "github.com/minio/minio-go",
    tag = "v6.0.5",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    commit = "3864e76763d9",
    importpath = "github.com/mitchellh/go-homedir",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    tag = "v1.1.2",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    commit = "bacd9c7ef1dd",
    importpath = "github.com/modern-go/concurrent",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    commit = "4b7aa43c6742",
    importpath = "github.com/modern-go/reflect2",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    tag = "v1.7.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    tag = "v1.4.3",
)

go_repository(
    name = "com_github_pborman_uuid",
    importpath = "github.com/pborman/uuid",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    importpath = "github.com/peterbourgon/diskv",
    tag = "v2.0.1",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    tag = "v0.8.0",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    importpath = "github.com/PuerkitoBio/purell",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    commit = "de5bf2ad4578",
    importpath = "github.com/PuerkitoBio/urlesc",
)

go_repository(
    name = "com_github_robfig_cron",
    commit = "b41be1df6967",
    importpath = "github.com/robfig/cron",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.0.6",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    commit = "b2de0cb4f26d",
    importpath = "github.com/smartystreets/assertions",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    commit = "044398e4856c",
    importpath = "github.com/smartystreets/goconvey",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    tag = "v1.1.2",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    tag = "v0.0.3",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    tag = "v1.0.3",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.2.2",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    commit = "d75b2dcb6bc8",
    importpath = "github.com/ugorji/go/codec",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    importpath = "github.com/valyala/bytebufferpool",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    commit = "dcecefd839c4",
    importpath = "github.com/valyala/fasttemplate",
)

go_repository(
    name = "com_github_vividcortex_mysqlerr",
    commit = "6c6b55f8796f",
    importpath = "github.com/VividCortex/mysqlerr",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    importpath = "github.com/xordataexchange/crypt",
    tag = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    tag = "v0.26.0",
)

go_repository(
    name = "in_gopkg_airbrake_gobrake_v2",
    importpath = "gopkg.in/airbrake/gobrake.v2",
    tag = "v2.0.9",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "20d25e280405",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    tag = "v1.4.7",
)

go_repository(
    name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
    importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
    tag = "v2.1.2",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    tag = "v0.9.1",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    tag = "v1.39.3",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    commit = "dd632973f1e7",
    importpath = "gopkg.in/tomb.v1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.2",
)

go_repository(
    name = "io_k8s_kube_openapi",
    commit = "d8ea2fe547a4",
    importpath = "k8s.io/kube-openapi",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    tag = "v1.1.0",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "bd91e49a0898",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable_global",
    importpath = "google.golang.org/grpc",
    tag = "v1.16.0",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "505ab145d0a9",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_lint",
    commit = "06c8688daad7",
    importpath = "golang.org/x/lint",
)

go_repository(
    name = "org_golang_x_net",
    build_file_proto_mode = "disable_global",
    commit = "351d144fa1fc",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "d2e6202438be",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "42b317875d0f",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "a5c9d58dba9a",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.0",
)

go_repository(
    name = "org_golang_x_time",
    commit = "fbb02b2291d2",
    importpath = "golang.org/x/time",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "6cd1fcedba52",
    importpath = "golang.org/x/tools",
)

go_repository(
    name = "com_github_beorn7_perks",
    commit = "3a771d992973",
    importpath = "github.com/beorn7/perks",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_go_logr_zapr",
    importpath = "github.com/go-logr/zapr",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_mattbaird_jsonpatch",
    commit = "81af80346b1a",
    importpath = "github.com/mattbaird/jsonpatch",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    tag = "v0.9.2",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "5c3871d89910",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_common",
    commit = "4724e9255275",
    importpath = "github.com/prometheus/common",
)

go_repository(
    name = "com_github_prometheus_procfs",
    commit = "1dc9a6cbc91a",
    importpath = "github.com/prometheus/procfs",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    tag = "v1.3.2",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    tag = "v1.1.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    tag = "v1.9.1",
)

go_repository(
    name = "com_github_google_pprof",
    commit = "3ea8567a2e57",
    importpath = "github.com/google/pprof",
)

go_repository(
    name = "org_golang_x_arch",
    commit = "5a4828bb7045",
    importpath = "golang.org/x/arch",
)
