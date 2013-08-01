Interview-Question--Smallest-Missing
====================================

A good interview question answered with Go.

In a recent interview, where I shared my screen with the interviewer over Skype, I was asked to develop an algorithm while talking through my methodology and stating the efficiency at each step of refinement.

The question was as followed: In any environment you're comfortable with, develop an algorithm to detect and print the smallest positive integer which is missing from a sequence of integers passed as an array. The array will not always be sorted and there can be repeats of the numbers.

I chose [Go](http://golang.org] as the language of choice here. Its well organized standard library covers a lot of possible requirements and the docs are easy to read with included examples. The syntax is also good for quick implementation or at least better than C/C++ which I normally work in.

I floundered a bit on the implementation but got an algorithm together that used the standard library's sort function to sort the list and then loop through the list, keeping track of the smallest element seen which is not more than one more than the previous smallest. This worked but the efficiency was limited by the most complex operation which in this case is the sort, quicksort, giving an average efficiency for the algorithm of O(nlogn).

This wasn't bad but the interviewer eventually lead me to another answer where I could use a hash to obtain an efficiency of O(n). In this case, I insert every element of the list into the hash with the number being its key and keeping track of the largest integer and then, as the hash will allow O(1) lookups, I start at 1 and ask in turn if each subsequent integer is present in the hash until I reach a missing number or the biggest integer in the original array.

Unfortunately, I didn't get the job but I thought that this question was interesting enough to post my answer with the two algorithm implementations. There's a small demo with a few test cases in the main function and I've also written up a quick benchmark using Go's awesome built in test and benchmarking features to get some more practice in.

The benchmarks show that at small values, the nlogn function wins over the n function but ultimately looses as predicted when n rises to large values as in this [graph](https://www.desmos.com/calculator/w6o9sxgq1a).

Basic stuff but certainly a good reminder to me about the efficiency of hashes after spending a lot of time on tiny embedded projects where you don't have the luxury of trading memory for performance or the clean hash syntax of modern languages.

Build with

    go build list1

Test with

    go test -bench .
