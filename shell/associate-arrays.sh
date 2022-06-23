#! /bin/bash

# Learning associative arrays
# supported in bash version 4 & above

declare -A testMap

testMap["foo"]="bar"
testMap[alice]="bob" # quotes in keys/values doesn't matter. "foo" or foo are treated the same
testMap[12345]=charlie # use quotes when key/value has spaces

echo -e "\n# Retrieving specific key"
echo "foo =>" ${testMap["foo"]}

echo -e "\n# Retrieving specific key using a variable"
testKey="alice"
echo "alice =>" ${testMap[$testKey]}

# Iterating the values
echo -e "\n# Iterating by values -> uses \${var[@]} syntax"
for value in ${testMap[@]};do
    echo $value
done

# Iterating the keys
echo -e "\n# Iterating by keys -> uses \${!var[@]} syntax"
for key in ${!testMap[@]};do
    echo "$key => ${testMap[$key]}"
done


echo -e "\nTesting a key's existence"
# map[key] returns empty string ("") for non-existent keys
# test whether value returned is non-zero length
if [[ ! -z ${testMap[$testKey]} ]]; then
    echo "Key ($testKey) is found"
fi

# test whether value returned is of zero length
if [[ -z ${testMap["moose"]} ]]; then
    echo "Key (moose) is not found"
fi

# Get length of the map
echo -e "\nLength of testMap=${#testMap[@]}"

echo -e "\nCreating a read only map. Add the option -r when declaring the array"
# array can be initialized along with declaration
declare -r -A credentials=(
    ["user"]="root"
    ["pass"]="root123"
)

for k in ${!credentials[@]};do
    echo "$k => ${credentials[$k]}"
done

# On read only maps, can't add new keys or modify existing values
# credentials["token"]="1010"
# credentials["user"]="admin"