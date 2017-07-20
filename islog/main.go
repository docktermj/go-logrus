package islog

import (
	"github.com/sirupsen/logrus"
)

type IsLog struct {
	lastLevel string
	isPanic   bool
	isFatal   bool
	isError   bool
	isWarn    bool
	isInfo    bool
	isDebug   bool
}

var isLog *IsLog

func init() {
	isLog = New()
}

func New() *IsLog {
	return new(IsLog)
}

// Determine if the metadata reflect the current debugging level.
func (this *IsLog) isCurrent() bool {
	return this.lastLevel == logrus.GetLevel().String()
}

// Determine if the current level is equal or greater than desired.
// Although the algorithm is a little inefficient for setting all
// levels, it will always reflect the order seen in logrus.AllLevels.
func (this *IsLog) isLevel(desiredLevel string) bool {
	current := 0
	desired := 0
	for key, value := range logrus.AllLevels {
		switch level := value.String(); level {
		case this.lastLevel:
			current = key
		case desiredLevel:
			desired = key
		}
	}
	return current >= desired
}

// Set the flags for all debug levels.
func (this *IsLog) setLevel() {
	this.lastLevel = logrus.GetLevel().String()
	this.isPanic = this.isLevel("panic")
	this.isFatal = this.isLevel("fatal")
	this.isError = this.isLevel("error")
	this.isWarn = this.isLevel("warn")
	this.isInfo = this.isLevel("info")
	this.isDebug = this.isLevel("debug")
}

// Determine if data is up-to-date.
func (this *IsLog) checkLevel() {
	if !this.isCurrent() {
		this.setLevel()
	}
}

// ----------------------------------------------------------------------------
// IsXXX() routines
// ----------------------------------------------------------------------------

func Panic() bool { return isLog.Panic() }
func (this *IsLog) Panic() bool {
	this.checkLevel()
	return this.isPanic
}

func Fatal() bool { return isLog.Fatal() }
func (this *IsLog) Fatal() bool {
	this.checkLevel()
	return this.isFatal
}

func Error() bool { return isLog.Error() }
func (this *IsLog) Error() bool {
	this.checkLevel()
	return this.isError
}

func Warning() bool { return isLog.Warning() }
func (this *IsLog) Warning() bool {
	this.checkLevel()
	return this.isWarn
}

func Info() bool { return isLog.Info() }
func (this *IsLog) Info() bool {
	this.checkLevel()
	return this.isInfo
}

func Debug() bool { return isLog.Debug() }
func (this *IsLog) Debug() bool {
	this.checkLevel()
	return this.isDebug
}
