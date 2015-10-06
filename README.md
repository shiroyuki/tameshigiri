# Tameshigiri

This is an additional test utility for writing tests in Go language. Here is an
example on what you can do.

``` go
package myapp

import "testing"
import "github.com/shiroyuki/tameshigiri"

func TestAssertBoolean(t *testing.T) {
    assertion := tameshigiri.NewAssertion(t)

    assertion.IsTrue(true, "It must be true.")
    assertion.IsFalse(false, "It must be false.")
}

func TestAssertEqual(t *testing.T) {
    assertion := tameshigiri.NewAssertion(t)

    expected = 1
    actual   = 1

    assertion.Equals(expected, actual, "This should be equal.")
}
```

## Installation

Just run ``go get github.com/shiroyuki/tameshigiri``.
