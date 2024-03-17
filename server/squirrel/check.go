package squirrel

import (
	"go.uber.org/zap"
	"os"
	"runtime"
	"server/logutil"
)

func checkSupportedArch() {
	lg, err := logutil.CreateDefaultZapLogger(zap.InfoLevel)
	if err != nil {
		panic(err)
	}
	if runtime.GOARCH == "amd64" ||
		runtime.GOARCH == "arm64" ||
		runtime.GOARCH == "ppc64le" ||
		runtime.GOARCH == "s390x" {
		//lg.Info("running arch [" + runtime.GOARCH + "] is supported")
		return
	}
	// unsupported arch only configured via environment variable
	// so unset here to not parse through flag
	defer os.Unsetenv("ETCD_UNSUPPORTED_ARCH")
	if env, ok := os.LookupEnv("ETCD_UNSUPPORTED_ARCH"); ok && env == runtime.GOARCH {
		lg.Info("running etcd on unsupported architecture since ETCD_UNSUPPORTED_ARCH is set", zap.String("arch", env))
		return
	}

	lg.Error("running etcd on unsupported architecture since ETCD_UNSUPPORTED_ARCH is set", zap.String("arch", runtime.GOARCH))
	os.Exit(1)
}
