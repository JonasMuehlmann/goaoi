# goaoi
![img](https://img.shields.io/badge/semver-2.0.0-green) [![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

Conventient algorithms for processing iterables, inspired by the algorithm header from the C++ standard template library (STL for short).

Please do not expect a stable API at this point.

Sister project: https://github.com/JonasMuehlmann/pyaoi

## Installation

```go get github.com/JonasMuehlmann/goaoi```

## How to use

All functions live in the ```goaoi``` namespace, you can import it with ```import "goaoi"``` and then call the functions
like this: ```goaoi.AllOf()```

Documentation available at https://pkg.go.dev/github.com/JonasMuehlmann/goaoi.

## Implemented functions

The following list shows planned functions and whether they are implemented yet. Feel free to make a PR for a listed
function's implementation. This list is subject to change at any time.
<details> <summary>Click to expand!</summary>
<p>

### Non-modifying sequence operations

- [x] all_of
- [x] any_of
- [x] none_of


- [x] for_each

- [x] count
- [x] count_if

- [x] mismatch

- [x] find
- [x] find_if
- [x] find_end
- [x] find_first_of
- [x] adjacent_find


- [x] copy_replace
- [x] copy_replace_if
- [x] copy_replace_if_not

- [x] copy_except
- [x] copy_except_if
- [x] copy_except_if_not

### Modifying sequence operations

- [x] fill
- [ ] fill_n


- [x] transform


- [ ] rotate


- [ ] shift_left
- [ ] shift_right


- [ ] random_shuffle
- [ ] shuffle


- [ ] sample


- [ ] unique
- [ ] unique_copy

### Partitioning operations

- [ ] is_partitioned


- [ ] partition
- [ ] partition_copy


- [ ] stable_partition


- [ ] partition_point

### Sorting operations

- [ ] is_sorted
- [ ] is_sorted_until


- [ ] partial_sort
- [ ] partial_sort_copy
- [ ] stable_sort
- [ ] nth_element

### Binary search operations (on sorted ranges)

- [ ] lower_bound
- [ ] upper_bound


- [ ] binary_search


- [ ] equal_range

### Other operations on sorted ranges

- [ ] merge
- [ ] implace_merge

### Set operations (on sorted ranges)

- [ ] includes


- [ ] set_difference
- [ ] set_intersection
- [ ] set_symmetric_difference
- [ ] set_union

### Heap operations

- [ ] is_heap
- [ ] is_heap_until


- [ ] make_heap


- [ ] push_heap


- [ ] pop_heap


- [ ] sort_heap

### Minimum/maximum operations

- [ ] max_index
- [ ] min_index
- [ ] minmax
- [ ] minmax_index


- [ ] clamp

### Comparison operations

- [ ] lexicographical_compare
- [ ] lexicographical_compare_threeway

### Permutation operations

- [ ] is_permutation


- [ ] next_permutation
- [ ] prev_permutation
</p>
</details>

## License
Copyright (C) 2021-2022 [Jonas Muehlmann](https://github.com/JonasMuehlmann)
 
The project is licensed under the terms of the MIT license, you can view it [here](LICENSE.md).
