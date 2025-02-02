package pruner

// This file contains some experimental rebuild policies that are not
// currently used anywhere.  Here just for examples and possible
// future use.

// liveRatioTrigger's Rebuild function returns true when there are at
// least MinLive live patterns and the ratio of removed to live
// patterns is greater than 1.
//
// This type is not used anywhere; just here as an example and maybe
// for future consideration.
type liveRatioTrigger struct {
	Ratio   float64
	MinLive int
}

func newLiveRatioTrigger(ratio float64, min int) *liveRatioTrigger {
	return &liveRatioTrigger{
		Ratio:   ratio,
		MinLive: min,
	}
}

func (t *liveRatioTrigger) Rebuild(added bool, s *Stats) bool {
	if added {
		return false
	}
	live := s.Live - s.Deleted
	if live == 0 {
		return false
	}
	if live < t.MinLive {
		return false
	}
	return t.Ratio <= float64(s.Deleted)/float64(live)
}

// neverTrigger is a rebuildTrigger that will never trigger a rebuild.
//
// Setting Matcher.rebuildTrigger to nil will have the same effect.
//
// This type is not used anywhere; just here as an example and maybe
// for future consideration.
type neverTrigger struct {
}

func newNeverTrigger() *neverTrigger {
	return &neverTrigger{}
}

func (t *neverTrigger) Rebuild(added bool, s *Stats) bool {
	return false
}
