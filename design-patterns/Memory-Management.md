# Memory Manegmenet in Golnag

There are two ways to memory initialization
- make 
- new


## Make:
- Used for Slice, Maps, and channels only
- Slice can take two values: size and capacity. The capacity can be extended using the append function
- performs initilaization
- For Maps, it can take only capacity, and it can be extended. Assigning capacity will help to better allocate memory
- Non-Zeored storage 
- return the data type
- most commonly used

## New:
- Can be used for any type (including int)
- It takes only one argument
- returns Zeored storage value (zeroed storage can be zero, nil, false)
- return the pointer to the data type
- Not often used
