package mwsparser

import (
        "testing"
)

func Test_parseMoney(t *testing.T) {
        m := parseMoney("9.99")

        if m != 999 {
                t.Fail()
        }

        m = parseMoney("0.00")

        if m != 0 {
                t.Fail()
        }
}
