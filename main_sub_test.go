package main

import (
  "testing"
)

func testInt(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

func TestSub(t *testing.T) {
  t.Run("abc", testInt)
}
