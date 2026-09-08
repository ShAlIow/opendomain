package handler

import "testing"

func TestSameNameserversIgnoresOrderCaseAndTrailingDot(t *testing.T) {
	left := []string{"NS2.NODELOOK.COM.", "ns1.nodelook.com"}
	right := []string{"ns1.nodelook.com.", "ns2.nodelook.com."}
	if !sameNameservers(left, right) {
		t.Fatalf("sameNameservers(%v, %v) = false, want true", left, right)
	}
}

func TestSameNameserversRejectsDifferentSet(t *testing.T) {
	left := []string{"ns1.nodelook.com", "ns3.nodelook.com"}
	right := []string{"ns1.nodelook.com", "ns2.nodelook.com"}
	if sameNameservers(left, right) {
		t.Fatalf("sameNameservers(%v, %v) = true, want false", left, right)
	}
}
