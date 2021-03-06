package tameshigiri

import "fmt"
import "runtime"
import "testing"

// Static reference number of the processed assertions
var NumberOfProcessedAssertion uint = 0

// Assertion class
//
// When the assertion fails, assuming that the given result is unexpected, the
// most recent call stacks (up to the size of 2KB) will be provided along with
// the human-readable description.
type Assertion struct {
    T                *testing.T
    stackDumpEnabled bool
}

// Create a new assertion.
func NewAssertion(t *testing.T) Assertion {
    assertion := Assertion{ T: t }

    assertion.EnableStackDump()

    return assertion
}

// Check if the result is true.
func (self *Assertion) IsTrue(result bool, description string) bool {
    NumberOfProcessedAssertion += 1

    if !result {
        if self.stackDumpEnabled {
            self.T.Logf("#%d FAILED\n", NumberOfProcessedAssertion)
            self.T.Logf("#%d %s\n", NumberOfProcessedAssertion, description)

            self.T.FailNow()
        }

        return false
    }

    return true
}

// Check if the result is false.
func (self *Assertion) IsFalse(result bool, description string) bool {
    return self.IsTrue(!result, description)
}

// Assert if the actual value is equal to the expected value
func (self *Assertion) Equals(expected interface{}, actual interface{}, description string) bool {
    var yes bool

    var stackDumpEnabled = self.stackDumpEnabled

    if stackDumpEnabled {
        self.DisableStackDump()
    }

    yes = self.IsTrue(expected == actual, "")

    if stackDumpEnabled {
        self.EnableStackDump()
    }

    if yes {
        return true
    }

    if stackDumpEnabled {
        self.T.Logf("#%d FAILED\n", NumberOfProcessedAssertion)
        self.T.Logf("#%d %s\n", NumberOfProcessedAssertion, description)

        prefix := fmt.Sprintf("#%d", NumberOfProcessedAssertion)

        self.T.Log(prefix, "Expected:", expected)
        self.T.Log(prefix, "Given:", actual)

        self.dumpStack()

        self.T.FailNow()
    }

    return false
}

// Enable stack dump
func (self *Assertion) EnableStackDump() {
    self.stackDumpEnabled = true
}

// Disable stack dump
func (self *Assertion) DisableStackDump() {
    self.stackDumpEnabled = false
}

// Dump call stacks
func (self *Assertion) dumpStack() {
    var buffer []byte

    if !self.stackDumpEnabled {
        return
    }

    buffer = make([]byte, 2048); // keep 2KB

    runtime.Stack(buffer, true)

    self.T.Logf("#%d Detail: %s\n", NumberOfProcessedAssertion, buffer)
}
