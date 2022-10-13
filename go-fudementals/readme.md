
# Types

# nil type
    - nill type is not same as null in other programming languages.
    - nill do not have a type
    - so it can be assigned or compared against values of different types.
    - A nil slice contains nothing.











# ----------------------------------------------------------------------
## SLICE
# ----------------------------------------------------------------------



# append() in slice
    var slice_00 []int
    var slice_01 = []int{1, 2, 3, 4}
    var slice_02 []int

    slice_00 = append(slice_00, slice_01...)
    slice_02 = append(slice_02, 2,3,4,5,6)


# make()
    - we can make slice using make() by define,
        - type
        - length
        - capacity ( optional )
    
    x := make([]int, 5) -> x is int type slice, which length=5 & capacity=5

    this will make slice like this
        [0, 0, 0, 0, 0]
    
    if we append 10
    x := append(x, 10)
    
    new slice will like this
     [0, 0, 0, 0, 0, 10]
    









