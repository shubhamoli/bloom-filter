// package bloom implements a very simple bloom filter
//
// Simple here means:
// - fixed length
// - elements can't be removed
// - not production ready in any sense
package bloom


// Word about external packages used
//
// 1. BitSet is not availble in standard go library so I've to import
//	  third-party library or as an alternative can use go's "big" package as bitset
//
// 2. murmur3 algorithm is used because it has best trade-off between
//    speed and uniformity
//	  inspired by this answer: https://stackoverflow.com/a/40343867

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)


type BloomFilter struct {
	m		uint				// size of bit vector
	k		uint				// no. of hash functions to use
	b		*bitset.BitSet		// bit vector
}


// Add an element to set
func (bf *BloomFilter) Add(key string) {
	data := []byte(key)

	// generate k index to set
	// using k hash functions
	for i := uint(0); i < bf.k; i++ {
		// find hash with seed value "i"
		// so that each time we get diff hash
		v1, _ := murmur3.Sum128WithSeed(data, uint32(i))
		// normalise to 0...m-1
		index := v1 % uint64(bf.m)
		// set bit at idex to 1
		bf.b.Set(uint(index))
	}
}

// query an element against set
func (bf *BloomFilter) Contains(key string) (bool) {
	data := []byte(key)

	for i := uint(0); i < bf.k; i++ {
		v1, _ := murmur3.Sum128WithSeed(data, uint32(i))
		// normalize to 0...m-1
		index := v1 % uint64(bf.m)
		// if any bit is found 0
		// return false immediately
		if ! bf.b.Test(uint(index)) {
			return false
		}
	}

	// Otherwise return true
	// Remeber: element may be false positive
	return true
}


// for formula: see wikipedia page of bloom filter
//
// and you can verify here
// https://www.di-mgt.com.au/bloom-calculator.html
// https://hur.st/bloomfilter/
func calcM(n uint, e float64) (uint) {
	tmp := - ((float64(n) * (math.Log(e))) / float64(math.Pow(math.Log(2), 2)))
	return uint(math.Ceil(tmp))
}


// for formula: see wikipedia page of bloom filter
//
// and you can verify here
// https://www.di-mgt.com.au/bloom-calculator.html
// https://hur.st/bloomfilter/
func calcK(e float64) (uint) {
	return uint(math.Ceil(-1 * math.Log2(e)))
}


// Initialise and return bloom filter
func Init(n uint, e float64) *BloomFilter {
	fmt.Printf("Info :: Creating BloomFilter for %d elements, with %.2f error rate\n", n, e)

	m := calcM(n, e)
	k := calcK(e)

	fmt.Printf("Info :: Required m and k are %d, %d \n", uint(m), k)

	// create bitset of size m and return BloomFilter
	return &BloomFilter{m, k, bitset.New(m)}

}


