package yalog

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger0 := NewLogger(
		"test",
		WithVerboseLevel(VerboseLevelDebug),
		WithPrintCaller(20),
		WithPrintTime("2006-01-02 15:04:05"),
		WithPrintLevel(),
		WithColorEnabled(),
		WithAnotherColor(VerboseLevelDebug, ColorBrightBlack),
		WithAnotherColor(VerboseLevelInfo, ColorCyan),
		WithPrintTreeName(1, true),
		WithDifferentOutput(os.Stderr),
		WithSecondOutput(os.Stdout, VerboseLevelInfo),
	)

	logger01 := logger0.NewLogger(
		"inner_0_1",
	)

	logger02 := logger0.NewLogger(
		"inner_0_2",
	)

	logger011 := logger01.NewLogger(
		"inner_0_1_1",
	)

	logger0.Debug("test;\n")
	logger0.Info("test;\n")
	logger0.Warning("test;\n")
	logger0.Error("test;\n")
	//logger_0.Fatal("test;")

	logger01.SetVerboseLevel(VerboseLevelWarning)
	logger01.Debugf("test: %s, %d\n", "debugf", 1)
	logger01.Infof("test: %s, %d\n", "infof", 1)
	logger01.Warningf("test: %s, %d\n", "warningf", 1)
	logger01.Errorf("test: %s, %d\n", "errorf", 1)
	//logger_0_1.Fatalf("test: %s, %d\n", "fatalf", 1)

	logger02.Debugln("test: debugln, 1")
	logger02.Infoln("test: infoln, 1")
	logger02.Warningln("test: warningln, 1")
	logger02.Errorln("test: errorln, 1")
	//logger_0_2.Fatalln("test: fatalln, 1")

	logger011.Debug("test;\n")
	logger011.Infof("test: %s, %d\n", "infof", 1)
	logger011.Warning("test;\n")
	logger011.Error("test;\n")
	//logger_0_1_1.Fatal("test;")

}
