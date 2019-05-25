# Adaptive Table

adaptive-table implements the underlying data structure described in the paper [Data Streams as Random Permutations: the Distinct
Element Problem](https://hal.inria.fr/hal-01197221/document).

In the paper, A. Helmi, J. Lumbroso, C. Martínez and A. Viola describe an algorithm for estimating the cardinality of a data stream. The algorithm counts the number of records in the underlying permutation of elements without taking into account repetitions. The authors explain in section 4.1 how the data structure, which is used for coutning the records (same structure used in MinHash), can "grow" following different strategies. The expected memory usage of the data structure is ![equation](https://latex.codecogs.com/gif.latex?O(\log{n}))  or ![equation](https://latex.codecogs.com/gif.latex?O(n^\alpha)) where ![equation](https://latex.codecogs.com/gif.latex?\alpha&space;\in&space;(0,&space;1/2)) and *n* is the number of distinct elements in the data stream.

This plot shows the final size of the table starting with sizes 2, 4, 8 , 16, 32 and 64 after inserting up to 5 billion elements

![mean_sizes](https://user-images.githubusercontent.com/3369356/58375186-accf8000-7f4d-11e9-9de2-190af1967456.png)

This implementation can be used in other algorithms, which can take advantage of the adaptative size of the table.


This implementation is used in some other go packages like:
- KMV for cardinality estiamtion [go-kmv](https://github.com/positiveblue/go-kmv)
- Minhash for set similarity [adaptive-minhash](https://github.com/positiveblue/adaptive-minhash)



