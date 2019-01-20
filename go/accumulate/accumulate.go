package accumulate

// Accumulate takes a collection and a function and applies the function to each element in the collection
func Accumulate(collection []string, function func(string) string) []string {
	for i := range collection {
		collection[i] = function(collection[i])
	}

	return collection
}
