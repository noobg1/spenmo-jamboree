## Output if sequence exists in the given array


### Usage/ Run tests
```
cd numberSequence && go test . -v
```

### Time and memory complexity

1. Time: O(N) 
   (though there are two loops => we move the needle of outer loop(numbers array) if there in match in inner loop(sequence arrays) and break if the inner loop if first element doesn't match)
2. Memory complexity: In-place

N: number of elements in array to searched
