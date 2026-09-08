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

func TestNormalizeNameserversTrimsLowercasesAndDedupes(t *testing.T) {
	got, err := normalizeNameservers([]string{"dns1.serv00.com\t", " DNS2.Serv00.com. ", "dns1.serv00.com", ""})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"dns1.serv00.com", "dns2.serv00.com"}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("normalizeNameservers = %v, want %v", got, want)
	}
}

func TestNormalizeNameserversRejectsInvalidHostnames(t *testing.T) {
	for _, bad := range [][]string{{"ns1"}, {"ns 1.example.com"}, {"-ns.example.com"}, {"ns1.example.com/"}, {"   "}} {
		if _, err := normalizeNameservers(bad); err == nil {
			t.Fatalf("expected error for %v", bad)
		}
	}
}

func TestEnsureCanonicalNSTrimsWhitespace(t *testing.T) {
	got := ensureCanonicalNS([]string{"dns1.serv00.com\t", " NS2.Example.com.", ""})
	if len(got) != 2 || got[0] != "dns1.serv00.com." || got[1] != "ns2.example.com." {
		t.Fatalf("ensureCanonicalNS = %v", got)
	}
}
