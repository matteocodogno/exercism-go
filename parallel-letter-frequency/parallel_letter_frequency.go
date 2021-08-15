package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type syncMap struct{
	sync.RWMutex
	m FreqMap
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}

	return m
}

func concurrencyFrequency(wg *sync.WaitGroup, counter *syncMap, sentence string) {
	defer wg.Done()

	fmap := Frequency(sentence)
	for k, v := range fmap {
		counter.Lock()
		counter.m[k] += v
		counter.Unlock()
	}
}

func ConcurrentFrequency(sentences []string) FreqMap {
	var wg sync.WaitGroup
	var counter = syncMap{m: make(FreqMap)}

	for _, sentence := range sentences {
		wg.Add(1)

		go concurrencyFrequency(&wg, &counter, sentence)
	}

	wg.Wait()
	return counter.m
}